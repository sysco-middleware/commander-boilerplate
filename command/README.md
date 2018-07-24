# Command

The command service is responsible for the creation of commands. Commands can be executed sync or async.
If a sync command is not resolved within 5 seconds is a timeout thrown.

## Getting started

To run the service manually for example during development.
Make sure that you have the environment variables set and have `kafka` running.

```bash
$ # ... import the environment variables
$ docker-compose up -d kafka
$ go run main.go
```

## Environment variables

```bash
export KAFKA_GROUP=command
export KAFKA_SERVERS=kafka:9092
```
