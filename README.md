This is a boilerplate project for [commander](https://github.com/sysco-middleware/commander).

## Getting started

In order to get started do you need to have [`docker`](https://docs.docker.com/install/) and [`docker-compose`](https://docs.docker.com/compose/install/) installed.
Clone the master branch of this repository and pull the latest docker images. Once the images are pulled can you build and start the services.

```bash
$ docker-compose pull
$ docker-compose build
$ docker-compose up -d
```

## State

Every part can hold it's own state/view of the source (events). The state can be used to validate uniqueness or fetch the current state of a row.
