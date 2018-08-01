# Docker

The project is divided over 2 docker compose files and are connected via the external networks `kafka` and `postgres`.
This way you are able to run all services across multiple clusters/machines.

## Getting started

In order to get started do you need to have [`docker`](https://docs.docker.com/install/) and [`docker-compose`](https://docs.docker.com/compose/install/) installed.
Clone the master branch of this repository. Once you have the repo can you build and start the services.

```bash
$ docker-compose build
$ docker-compose up -d
```

NOTE:
The project path should not end with a `/`
