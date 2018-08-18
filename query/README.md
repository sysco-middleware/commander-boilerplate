# Query

The query service is accessible to the consumer facing api or any third party clients.
This service preforms queries on the projections created by the projectors.

## Getting started

To run the service manually for example during development make sure that you have all `environment` variables set and have `postgres` running.

```bash
$ # cd docker/dependencies
$ docker-compose up -d postgres
$ # cd query/
$ go run main.go
```

## Environment variables

These are the required environment variables.

```bash
export POSTGRES_HOST=postgres
export POSTGRES_PORT=5432
export POSTGRES_USER=commander
export POSTGRES_PASSWORD=TFgvT3Pb3bWEhXKAfgMk63bo
export POSTGRES_DB=commander
export HOST_ADDRESS=:8080
```
