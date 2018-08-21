# Logic

The logic service consumes commands which should always result in a resulting event.

## Getting started

To run the service manually for example during development.
Make sure that you have the `environment` variables set and have `kafka` and `postgres` running.

```bash
$ # cd docker/dependencies
$ docker-compose up -d kafka postgres
$ # cd logic/
$ go run main.go
```

## Environment variables

These are the required environment variables.

```bash
export KAFKA_GROUP=users
export KAFKA_BROKERS=kafka:9092
export POSTGRES_HOST=postgres
export POSTGRES_PORT=5432
export POSTGRES_USER=commander
export POSTGRES_PASSWORD=TFgvT3Pb3bWEhXKAfgMk63bo
export POSTGRES_DB=commander
export COMMANDER_EVENT_TOPIC=events
export COMMANDER_COMMAND_TOPIC=commands
```
