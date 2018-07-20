# Command

The command service is responsible for the creation of commands. Commands can be executed sync or async. If a sync command is not resolved within 5 seconds is a timeout thrown.

## Environment variables

```
export KAFKA_GROUP=command
export KAFKA_SERVERS=kafka:9092
```
