This is a boilerplate project for [commander](https://github.com/sysco-middleware/commander).

![The pattern](https://github.com/sysco-middleware/commander/wiki/commander-pattern.jpg)

## Getting started

Check out the README in the docker directory. Or you could run every service manually. When wanting to run a service manually make sure to have the required `environment` variables set.

## State

Every part can hold it's own state/view of the source (events). The state can be used to validate uniqueness or fetch the current state of a row.
