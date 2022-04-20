# pgwalstreams

Event sourcing events from Postgresql database to Nats Jetstreams in kubernetes with [wal-listener](https://github.com/ihippik/wal-listener)

## Tutorial

### Requirements

* kubectl
* k3d
* helm
* okteto [optional]

### Deploy

* Start local k8s:

```bash
k3d cluster create -c k3d.yaml
export KUBECONFIG=$(k3d kubeconfig write poc)
kubectl cluster-info
```

* Helm-charts repositories:

```bash
helm repo add bitnami https://charts.bitnami.com/bitnami
helm repo add nats https://nats-io.github.io/k8s/helm/charts/
helm repo update
```

* Create postgresql resouces:

```bash
kubectl create configmap postgresql-conf --from-file=postgresql/postgresql.conf
helm upgrade --install postgres -f k8s/postgres-values.yaml bitnami/postgresql
```

* Create nats resouces:

```bash
kubectl apply -f https://raw.githubusercontent.com/nats-io/nack/v0.6.0/deploy/crds.yml
helm upgrade --install nats -f k8s/nats-values.yaml nats/nats
helm upgrade --install nack -f k8s/nack-values.yaml nats/nack
```

* Create pgwalstreams resouces:

```bash
kubectl apply -f k8s/pgwalstreams.yaml
```

### Develop

* Start dev container with hot-reload:

```bash
okteto up -f okteto.yml
```

### Test

```bash
cat <<EOF | kubectl apply -f -
apiVersion: jetstream.nats.io/v1beta2
kind: Stream
metadata:
  name: db
spec:
  name: db
  subjects: ["db.>"]
  storage: memory
  maxAge: 1h
  replicas: 1
EOF
```

```
cat <<EOF | kubectl apply -f -
apiVersion: jetstream.nats.io/v1beta2
kind: Consumer
metadata:
  name: public-users
spec:
  streamName: db
  durableName: public-users
  deliverPolicy: all
  filterSubject: db.public.users.>
  maxDeliver: 20
  ackPolicy: explicit
EOF
```

* Adds new records to the database tables

```bash
kubectl exec -ti pod/postgresql-0 -- bash -c 'PGPASSWORD=$POSTGRES_POSTGRES_PASSWORD psql -U postgres -p 5432 -d app' < postgresql/init.sql
```

* Check pgwalstreams logs

```bash
kubectl logs deployment.apps/pgwalstreams
```

* Check nats streams queue

```bash
kubectl exec -it deploy/nats-box -- /bin/sh -l
nats-box:~# nats consumer next db public-users
```

```bash
kubectl delete stream db
kubectl delete consumer public-users
```

* Monitoring:

  https://docs.nats.io/reference/faq#how-can-i-monitor-my-nats-cluster


```bash
kubectl exec -it deploy/nats-box -- /bin/sh -l
nats-box:~# nats-top
```

```bash
cat <<EOF | kubectl apply -f -
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: nats
spec:
  rules:
    - host: nats.k8s.local
      http:
        paths:
          - path: /
            pathType: ImplementationSpecific
            backend:
              service:
                name: nats
                port:
                  number: 8222
EOF
```

```bash
echo "172.18.0.2 nats.k8s.local" | sudo tee -a /etc/hosts

firefox nats.k8s.local
```

## License

This project is licensed under the terms of the APACHE license.

Original work comes from [ihippik/wal-listener](https://github.com/ihippik/wal-listener)
