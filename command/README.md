# Command

The command service is responsible for the creation of commands. A command request can be executed in a sync or async manner.
If a sync command is not resolved within 5 seconds a timeout is thrown.

## Getting started

To run the service manually for example during development.
Make sure that you have the `environment` variables set and have `kafka` running.

```bash
$ # cd docker/dependencies
$ docker-compose up -d kafka
$ # cd command/
$ go run main.go
```

## Environment variables

These are the required environment variables.

```bash
export KAFKA_GROUP=command
export KAFKA_BROKERS=kafka:9092
export COMMANDER_EVENT_TOPIC=events
export COMMANDER_COMMAND_TOPIC=commands
export HOST_ADDRESS=:7070
```
