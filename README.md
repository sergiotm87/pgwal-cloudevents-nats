# pgwal-cludevents-nats

## how-to

docker-compose up -d

docker-compose exec -u postgres db bash -c 'psql -U $POSTGRES_USER -d $POSTGRES_DB' < postgresql/init.sql
