# Logic

The logic service consumes commands which results in events.

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
export KAFKA_GROUP=users
export KAFKA_SERVERS=kafka:9092
export POSTGRES_HOST=postgres
export POSTGRES_PORT=5432
export POSTGRES_USER=commander
export POSTGRES_PASSWORD=TFgvT3Pb3bWEhXKAfgMk63bo
export POSTGRES_DB=commander
```