# Tog CLI

[![npm version](https://badge.fury.io/js/tog-cli.svg)](https://badge.fury.io/js/tog-cli)
[![CircleCI](https://circleci.com/gh/escaletech/tog-cli/tree/master.svg?style=svg)](https://circleci.com/gh/escaletech/tog-cli/tree/master)

Tog (short for toggle) is a framework for clients and servers to converse about feature flags over Redis.

This is the command-line tool that interacts with the [Server API](https://github.com/escaletech/tog-server) to update flags and experiments.

## Prerequisites

1. [Node.js](https://nodejs.org) >= 10.0


## Versioning

Tog command-line uses [Semantic Versioning 2](https://semver.org/spec/v2.0.0.html).


## Getting Started

1. Install: `npm install -g tog-cli`
2. Log in: `tog login -h <host-url>`, where `<host-url>` is the address to your deployment of [Tog Server](https://github.com/escaletech/tog-server) (e.g. `https://tog.mysite.com`)

## Examples

* Set a `foobar` as default namespace: `tog config namespace foobar`

### Flags

* List flags: `tog flags`
* Get a flag: `tog flag my-flag`
* Enable a flag: `tog flag my-flag on`, or disable it: `tog flag my-flag off`

### Experiments

* List experiments: `tog exp`
* Get an experiment: `tog exp my-exp`
* Create/update an experiment `tog exp my-exp --on flag-one --on flag-two --weight 30`
* Disable an expertiment: `tog exp my-exp --weight 0` (just set its weight to zero)
