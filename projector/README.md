# Projector

A projector creates a state from the consumed events. Each projector is responsible for a single model.
Multiple projectors could exists in a project each generating their own projections.

A projector could be connected to a third party/external system (Analytics, Fraud detection, Monitoring etc.). This makes sure that when needed to integrate with other systems the main projections/business logics is not touched.

## Getting started

To run the service manually for example during development make sure that you have all `environment` variables set and have `kafka` and `postgres` running.

```bash
$ # cd docker/dependencies
$ docker-compose up -d kafka postgres
$ # cd projector/
$ go run main.go
```

## Environment variables

These are the required environment variables.

```bash
export KAFKA_GROUP=projector
export KAFKA_BROKERS=kafka:9092
export POSTGRES_HOST=postgres
export POSTGRES_PORT=5432
export POSTGRES_USER=commander
export POSTGRES_PASSWORD=TFgvT3Pb3bWEhXKAfgMk63bo
export POSTGRES_DB=commander
export COMMANDER_EVENT_TOPIC=events
```
