# pgwalstreams

Event sourcing from Postgresql database to Nats Jetstreams in kubernetes with [wal-listener](https://github.com/ihippik/wal-listener):

## Overview

A service that helps implement the **Event-driven architecture**.

To maintain the consistency of data in the system, we will use **transactional messaging** - 
publishing events in a single transaction with a domain model change.

The service allows you to subscribe to changes in the PostgreSQL database using its logical decoding capability 
and publish them to the NATS Jetstreams server.

### Logic of work
To receive events about data changes in our PostgreSQL DB
  we use the standard logic decoding module (**pgoutput**) This module converts
 changes read from the WAL into a logical replication protocol.
  And we already consume all this information on our side.
Then we filter out only the events we need and publish them in the queue

### Event publishing

NATS Jetstreams is used as a message broker.
Service publishes the following structure.
The name of the topic for subscription to receive messages is formed from:
* the topic prefix
* the name of the database
* the name of the table
* the action on the record

Message structure: `prefix.schema_table.table.action`

Example: `db.public.users.insert`

```
{
	ID        uuid.UUID   # unique ID           
	Schema    string                 
	Table     string                 
	Action    string                 
	Data      map[string]interface{} 
	EventTime time.Time   # commit time          
}
```

### Filter configuration example

```yaml
database:
  filter:
    tables:
      users:
        - insert
        - update

```
This filter means that we only process events occurring with the `users` table, 
and in particular `insert` and `update` data.

Filter tables from postgresql events is updated when `database.filter` config is updated via k8s configmap / docker-compose volume.

### DB setting
You must make the following settings in the db configuration (postgresql.conf)
* wal_level >= “logical”
* max_replication_slots >= 1

The publication & slot created automatically when the service starts (for all tables and all actions). 
You can delete the default publication and create your own (name: _wal-listener_) with the necessary filtering conditions, and then the filtering will occur at the database level and not at the application level.

https://www.postgresql.org/docs/current/sql-createpublication.html

If you change the publication, do not forget to change the slot name or delete the current one.

## Tutorial

### Requirements

* [kubectl](https://kubernetes.io/docs/reference/kubectl/kubectl/)
* [k3d](https://k3d.io/v5.4.1/) / [kind](https://kind.sigs.k8s.io/)
* [helm](https://helm.sh/)

### Deploy

* Start local k8s:

```bash
k3d cluster create -c deploy/k8s/k3d.yaml
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
kubectl create configmap postgresql-conf --from-file=deploy/postgresql/postgresql.conf
helm upgrade --install postgres -f deploy/k8s/postgres-values.yaml bitnami/postgresql
```

* Create nats resouces:

```bash
kubectl apply -f https://raw.githubusercontent.com/nats-io/nack/v0.6.0/deploy/crds.yml
helm upgrade --install nats -f deploy/k8s/nats-values.yaml nats/nats
helm upgrade --install nack -f deploy/k8s/nack-values.yaml nats/nack
```

* Create pgwalstreams resouces:

```bash
kubectl apply -f deploy/k8s/pgwalstreams.yaml
```

### Develop

* Start dev container with hot-reload using [okteto](https://www.okteto.com/):

```bash
okteto up
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
kubectl exec -ti pod/postgresql-0 -- bash -c 'PGPASSWORD=$POSTGRES_POSTGRES_PASSWORD psql -U postgres -p 5432 -d app' < deploy/postgresql/init.sql
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
echo "127.0.0.1 nats.k8s.local" | sudo tee -a /etc/hosts

firefox http://nats.k8s.local:8080/
```

## License

This project is licensed under the terms of the APACHE license.

Original work comes from [ihippik/wal-listener](https://github.com/ihippik/wal-listener)
