# Tog CLI

![GitHub release (latest SemVer)](https://img.shields.io/github/v/release/escaletech/tog-cli)
![GitHub Workflow Status (branch)](https://img.shields.io/github/workflow/status/escaletech/tog-cli/continuous-integration/master)

Tog (short for toggle) is a framework for clients and servers to converse about feature flags over Redis.

This is the command-line tool that interacts with the [Server API](https://github.com/escaletech/tog-management-server) to update flags and experiments.


## Table of contents

- [Tog CLI](#tog-cli)
  - [Table of contents](#table-of-contents)
  - [Usage](#usage)
  - [Installation](#installation)
    - [macOS](#macos)
    - [Windows (via Scoop)](#windows-via-scoop)
    - [Linux](#linux)
      - [Debian/Ubuntu Linux](#debianubuntu-linux)
      - [Fedora Linux](#fedora-linux)
      - [Centos Linux](#centos-linux)
      - [openSUSE/SUSE Linux](#opensusesuse-linux)
    - [Any platform, using Go](#any-platform-using-go)

## Usage

* `tog config`
* `tog login`
* `tog list`
* `tog get <flag-name>`
* `tog set <flag-name> [...options]`
* `tog delete <flag-name>`
* `tog help [<command>]`


## Installation

### macOS

Install: `brew install escaletech/tog/tog`

Upgrade: `brew update && brew upgrade tog`

### Windows (via [Scoop](https://scoop.sh/))

Install:

```
scoop bucket add github-gh https://github.com/cli/scoop-gh.git
scoop install gh
```

Upgrade: `scoop update tog`

### Linux

#### Debian/Ubuntu Linux

Install and upgrade:

1. Download the `.deb` file from the [releases page][]
2. `sudo apt install git && sudo dpkg -i tog_*_linux_amd64.deb`  install the downloaded file

#### Fedora Linux

Install and upgrade:

1. Download the `.rpm` file from the [releases page][]
2. `sudo dnf install tog_*_linux_amd64.rpm` install the downloaded file

#### Centos Linux

Install and upgrade:

1. Download the `.rpm` file from the [releases page][]
2. `sudo yum localinstall tog_*_linux_amd64.rpm` install the downloaded file

#### openSUSE/SUSE Linux

Install and upgrade:

1. Download the `.rpm` file from the [releases page][]
2. `sudo zypper in tog_*_linux_amd64.rpm` install the downloaded file

### Any platform, using Go

1. Verify that you have Go 1.13+ installed
    ```sh 
    $ go version
    go version go1.13.7
    ```
2. Go get
    ```sh
    go get -u github.com/escaletech/tog-cli
    ```

[releases page]: https://github.com/escaletech/tog-cli/releases/latest
