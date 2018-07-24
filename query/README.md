# Query

The query service is accessable to the consumer api or third party clients.
This service preforms queries on the projections created by the projectors.

## Getting started

To run the service manually for example during development.
Make sure that you have the environment variables set and have `postgres` running.

```bash
$ # ... import the environment variables
$ docker-compose up -d postgres
$ go run main.go
```

## Environment variables

```bash
export POSTGRES_HOST=postgres
export POSTGRES_PORT=5432
export POSTGRES_USER=commander
export POSTGRES_PASSWORD=TFgvT3Pb3bWEhXKAfgMk63bo
export POSTGRES_DB=commander
```
