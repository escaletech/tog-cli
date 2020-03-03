# Tog CLI

[![npm version](https://badge.fury.io/js/tog-cli.svg)](https://badge.fury.io/js/tog-cli)
[![CircleCI](https://circleci.com/gh/escaletech/tog-cli/tree/master.svg?style=svg)](https://circleci.com/gh/escaletech/tog-cli/tree/master)

Tog (short for toggle) is a framework for clients and servers to converse about feature flags over Redis.

This is the command-line tool that interacts with the [Server API](https://github.com/escaletech/tog-management-server) to update flags and experiments.

## Prerequisites

1. [Node.js](https://nodejs.org) >= 10.0


## Versioning

Tog command-line uses [Semantic Versioning 2](https://semver.org/spec/v2.0.0.html).


## Getting Started

1. Install: `npm install -g tog-cli`
2. Log in: `tog login -h <host-url>`, where `<host-url>` is the address to your deployment of [Tog Server](https://github.com/escaletech/tog-server) (e.g. `https://tog.mysite.com`)

## Examples

```sh
# Set my_app as default namespace
> tog config namespace my_app

# List flags
> tog list
namespace: my_app
┌──────────────┬─────────────┬───────────────┐
│ name         │ description │ rollout       │
├──────────────┼─────────────┼───────────────┤
│ blue-button  │ -           │ - value: true │
└──────────────┴─────────────┴───────────────┘

# Get a flag
> tog get blue-button
namespace: my_app
name: blue-button
description: "Make the button blue"
rollout:
  - value: true
    percentage: 30
  - value: false

# Set a flag's description
> tog set blue-button -d "Make the button blue"

# Set a flag's rollout
> tog set blue-button --rollout "[{ percentage: 30, value: true }, value: false]"

# Set a flag's rollout to always true
> tog set blue-button --on

# Set a flag's rollout to always false
> tog set blue-button --off
```
