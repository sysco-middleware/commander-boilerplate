# Docker

This project contains 3 docker compose files.

- **Dependencies** - This docker compose file contains the dependencies (kafka, postgres ...) used by commander.
- **Images** - This docker compose file contains the images used by the services during the building process.
- **Services** - This docker compose file contains the commander services (command, logic ...).

The dependencies and services are separated of one another and communicate over an external docker network.
This way you are able to run de dependencies and services across multiple clusters/machines.

## Getting started

In order to get started do you need to have [`docker`](https://docs.docker.com/install/) and [`docker-compose`](https://docs.docker.com/compose/install/) installed.
Clone the master branch of this repository. Once you have the repo can you build and start the services.

```bash
$ # cd docker/dependencies/
$ docker-compose up -d
$ # cd docker/services/
$ docker-compose build
$ docker-compose up -d
```
