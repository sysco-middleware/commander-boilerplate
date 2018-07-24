# Projector

A projector creates a current state from the consumed events. Each projector is responsible for a single model.
Multiple projectors could exists in a project each generating their own projections.

## Getting started

To run the service manually for example during development.
Make sure that you have the environment variables set and have `kafka` and `postgres` running.

```bash
$ # ... import the environment variables
$ docker-compose up -d kafka postgres
$ go run main.go
```

## Environment variables

```bash
export KAFKA_GROUP=projector
export KAFKA_SERVERS=kafka:9092
export POSTGRES_HOST=postgres
export POSTGRES_PORT=5432
export POSTGRES_USER=commander
export POSTGRES_PASSWORD=TFgvT3Pb3bWEhXKAfgMk63bo
export POSTGRES_DB=commander
```
