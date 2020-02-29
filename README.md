# Lionel

[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=lionel&metric=alert_status)](https://sonarcloud.io/dashboard?id=lionel)
[![Go Report Card](https://goreportcard.com/badge/github.com/benmatselby/lionel?style=flat-square)](https://goreportcard.com/report/github.com/benmatselby/lionel)

_Trello, is it more you're looking for..._

CLI application for getting information out of [Trello](http://trello.com)

```text
CLI application for retrieving data from Trello

Usage:
  lionel [command]

Available Commands:
  boards      List all the boards
  burndown    Provide a burndown table using the 'scrum for trello' plugin data
  cards       List all the cards for a given board
  help        Help about any command
  people      List all the cards for a given board sorted by the people the card is assigned to

Flags:
      --config string   config file (default is $HOME/.benmatselby/lionel.yaml)
  -h, --help            help for lionel

Use "lionel [command] --help" for more information about a command.
```

## Examples

The burndown table

```text
$ lionel burndown "Sprint 2018.8"
List           Cards Story Points
----           ----- ------------
Sprint Backlog 2     0
Doing          5     0
Stuck          8     0
Code Review    10    0
Ready to Test  15    37
Done           42    85
-----          ----- ------------
Total          42    85
-----          ----- ------------
```

The boards list

```text
$ lionel boards
How to earn a million dollars
Family Days Out
Garden Project
Meal Ideas
```

The cards list for a board

```text
$ lionel cards "Golang the musical"
To do (1)
=========

* (1) Perform the musical

Progress (1)
============

* (250) Write a musical
```

## Requirements

This application uses Go modules, so you will require:

- [Go version 1.11+](https://golang.org/dl/)
- Some of the features will require that you use the "[Scrum for Trello](http://scrumfortrello.com)" plugin in Trello

## Configuration

### Environment variables

You will need the following environment variables defining:

```bash
export TRELLO_CLI_KEY=""
export TRELLO_CLI_SECRET=""
```

Creating a Trello API Token is documented [here](https://developers.trello.com/page/authorization)

## Installation via Docker

Other than requiring [docker](http://docker.com) to be installed, there are no other requirements to run the application this way. This is the preferred method of running the `lionel`. The image is [here](https://hub.docker.com/r/benmatselby/lionel/).

```shell
$ docker run \
  --rm \
  -t \
  -eTRELLO_CLI_KEY \
  -eTRELLO_CLI_SECRET \
  -v "${HOME}/.benmatselby":/root/.benmatselby \
  benmatselby/lionel:latest "$@"
```

The `latest` tag mentioned above can be changed to a released version. For all releases, see [here](https://hub.docker.com/repository/docker/benmatselby/lionel/tags). An example would then be:

```shell
benmatselby/lionel:version-1.1.0
```

This would use the `verson-1.1.0` release in the docker command.

## Installation via Git

```bash
git clone git@github.com:benmatselby/lionel.git
cd lionel
make all
./lionel
```

You can also install into your `$GOPATH/bin` by `go install`

## Easy like Sunday morning

```shell
▒▓▓▓▒░▒▓▒▒▒▒▒▒▒▓▒▒▒▒▒▒▒▒▒▒░░░▒▓████████████▓▒▓▓▒▒▒▒▒▒▒▒▒░░░░░░▒▒▒░▒▒▒▒▒▒▒▓▒▒▓░▒
▓▒▒▒▒▒▒▓▒▒▒▒▒▒▒▒▒░▒▒▒▒▒▒▒▒▓█████████████████████▓▒▒▒▒▒▒▒▒▒▓▓▓▓▒▓▓▒▒▒▒▒▒▒▒▒▒▒▒░▒
▒▒▒▒▒▓▓▓▓▓▓▒▒▒▒▒▒▒▓▒▓▓▓██████████████████████████████▒▒▒▒░▒▒▒▒▒▒▒▒▓▒▓▒▒▓▒▓▓▒▒▓▒
▒▒▒▒▒▒▓▓▒▓█▓▒▒▒▓▒▒▓▓▓██████████████████████████████████▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▓▓▓▓▒
▓▓▓▓▓█▓▓▓▓▒▒▓▓▓▓▓▒██████████████████████████████████████▓▒░░░░▒▒▒▒░▒░░▒▒▒░▒▓▒█▒
▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▓██████████████████████████████████████████░▒░░░░▒▒░░░░░░▒▒▒▒▓▒
▒▒▒▒▒▒▒▒▒▒▒▒▒▒░░▓███████████████████████████████████████████▒░░░░░▒░░░░▒▒░▒▒▒▒▓
▒▒▒▒░░░░░▒▒░▒░░██████████████████████████████████████████████░░░░░▒░░░░▒▒░▒░▒▒▒
▒░▒░░░▒▒▒░▒░▒░▒███████████████████████████████████▓██████████▒░░░▒▒░░░░▒▒▒▒▒▒▓▒
▒▒▒░░░▒▒▒░░░▒░██████████████▓▓▓▓▓▓▓▓▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▓▓▓███████▓░▒▒▒▒░░░░▒▒▒░▒▒▓▓
▒▓▒▒▒▒▒▒░░░░░▒████████▓▓▓▓▓▒▒▒▒▒▒▒▒▒▒░░░░░▒▒▒▒▒▒▒▒▓▓▓▓▓███████░░▒▒▒▒▒░░▓▒▒▒░▒▒▒
▒▒▒▒░░░▒▒▒░▒▒▒███████▓▓▓▓▒▒▒▒▒▒▒▒▒▒▒▒▒▒░░▒▒▒▒▒▒▒▒▒▓▓▓▓▓███████▓▒▒▒▓▓▒░▒▒▓▒▒▒▒▓▒
▒▒▒▒▒▒▒▒▓▓▓▓▓▓████████▓▓▓▓▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▓▓▓▓████████▓▓▒▒▒▓▒▒░▒▓░▓▒▒▒▒
▒▒▒▒▒▒▒▒▒▒▒▒▒█████████▓▓▓▓▒▒▒▒▒▒▒▒▒▒▒░▒░▒░▒▒▒▒▒▒▒▒▓▓▓▓████████▓░░░░▒░░▒░▒░░░░▒▒
▓▓▒▒▒▒▒▒▒▒▒░▒▒█████████▓▓▓▒▒▒▒▒▒▒▒▒▒▒░░▒▒▒▒▒▒▒▒▒▒▒▒▓▓▓▓███████▓░░░░▒▒░░░░░░░░▒▒
▒▒▒▒▒▒▒░▒▒▒▒▒▒████████▓▓▓▓█▓▓▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▓▓██▓▓▓▓▓██████░░░░░░░░░░░░░░░▒░
▒▒░░░▒▒▒▒▒▒▒▒▒▓██████▓████▓▓████▓▓▒▒▒▒▒▒▒▒▓▓▓███▓▓▓███▓▓█████▓░▒░▒░░▒▒░░░░░░░░▒
▒▒▒░▒▒▒▒▒▒▒▓▒▒▒▓█████▓███▓█████▓▓▓▓▓▒▒▒▒▓█▓▓▓██████▓▓█▓▓▓████▒░░▒░░▒░▒░▒░░░░░░░
▒▒▒░░▓▓▓▒░▒▒▒▒▒▓█████▓▓▓▓█▓▒██▓▒▓▓▓▓▒▒░▒▒▓▓▓▒▒▓▓▓▒▓██▒▓▓▓███▓▒░░░░░░░░░▒░░░░░░░
▒▒░▒░░▒▒▓▒▒▒▒▒▓▓▓███▓▓▓▒▒▓▓█▓▓▓▓▒▒▒▒▒▒░▒▒▒▒▒▒▒▓▓▓▓▒▒▒▒▒▓▓█▓█▓▓▒▒▒▒░░▒▒▒▒▒▒▓▓███
▒▒▒▒▒▒▒░▒▒▒▓▓▓▓▓▓████▓▓▒▒▒▒▒▒▒▒▒▒▒▒▓▒░░▒▒▓▒▒▒▒▒▒▒▒▒▒▒▒▒▓▓███▓▓▒▓▓▓▒▒▒▒▒▒▓▒▒░▒▒▓
▒▒▒▒▒▒▒▒▒▒▒▓▓▓████▓███▓▓▓▒▒▒▒▒▒░▒▒▓▓▒▒▒▒▒▒▓▒▒░░░░▒▒▒▒▒▓▓██▓▓▓▓▓▓▒▒▒▒▒░░▒▒░▒▒░▒░
▒▒▒▒▒▒▒▒▒▒▓▓██████████▓▓▓▒▒▒░░░▒▒▓▒▓▒▒░▒▒▒▓▓▒▒░░░▒▒▒▒▓▓▓████▓▓▓▓▓▓▓▓▒▒▒▓▓▒▓▒▓▒░
▒▒▒▒░▒▒▒▒▒▓████████████▓▓▓▒▒▒▒▒▒▒▒▒▒▒░░░▒▓▒▒▒▒▒▒▒▒▒▒▒▓▓▓▓███▓▓▒▒▒▒▒▒▒▒░▒░░▒▓▓▒▒
▒▒▒▒▒░▒▒▒▒▓▒▒▒▓▓▓██████▓▓▓▓▒▒▒▒▒▓▒▒▓▒▒▒▒▒▓▒▓▓▒▒▒▒▒▒▓▓▓▓▓▓██▓▓▓▓▒▒▒▒▒▒▒░░░░▒▒▒▒▒
▒▒░▒▒▒▒░░▒░▒▒▓▓▓▓███████▓▓▓▓▒▒▒▒███████████▓▓▒▒▒▒▒▓▓▓▓▓▓▓█▓▓▓▓▓▒▒▒▒▒▒░░░░▒▒▓▒▒░
░░▒▒▒▒▒▒▒▒░▒▒▓▓▓▓████████▓▓▓▒▒▒▒▒▓█████████▓▓▒▒▒▒▓▓▓▓▓▓▓███▓▓▓▓▒▒▒▒▒▒▒▒░░▒▒▓▒▒▓
░▒▒▒▒▒▒▒▒▒▒▒▒▒▓▓▓▓███████▓▓▓▓▓█████▓▓▓█▓▒███████▓▓▓▓▓▓▓███▓▓▒▒▓▓▒▒▒▒▒░▒░▒▒▒░▒▓█
░▒▒▒▒▒▒▒▒░░░░▒▒▓▓▓████████▓▓██████▓▓▓▓▓▓▓▓▓▓█████▓▓▓▓▓████▓▓▓▓▓▒▒▒▒░░░▒░░▒░▒▒▒▒
░░▒▒▒▒▒░░░░░░▒▒▒▓▓▓█▓██████████▓▓▒▒▒▒▒▓▓▒▒▒▒▓▓▓▓▓▓▓█▓████▓▓▒▓▒▒▓▓▓▒▒▓▒░▒▒▒▒▓▒▒▒
░░▒▒▒▒▒░░░░░▒▒▒▒▓▓███████████▓▓▓▓█▓▓▒▒▒▒▒▒▓██▓▓▓▓▓██████▓▓▓▓██▓▓▓▓▓▓▓▓▓▓▒▒▒▒▒▒▒
▒▒▒▓▒▒▒▒▒▒▒░▒▓▒▓▓▓██▓████████▓▓▓███████████▓▓▓▓▓███████▓█▓▓▓▓▓▓▓▒▒▒▒▒▒▒▓▒▒▒░░░▒
▒▒▒▓▓▒▓▒░▒▒▒▓▓▓▒▒▒▓███████████▓▓▓▓▓▓████▓▓▓▓▓▓▓███████▓▓▓▓▓▓▓▓▓▓▓▓▓▒▒▒▒▒░░░▒▒▒▒
▒▒▓██▒▒▒▒▒▓▒▓▒▓▓▓▓▓██████████████▓▓▒▒▒▒▒▒▒▒▓█▓███████▓▓▓▓▓▓▓▓▓▒▓▓▓▓▓▓▓▓▓▓▒░▒░▒░
▓▒▓▓▓▒▒▒▒▓▒▒▒▓▒▓▓▓▓▓▓▓██████████████▓▓▓▓▓▓▓█████████▓▓▓▓▓▓▓▓▓▒▒▒▒▒▒▒▒▓▓▓▓▓▓▓▓▓▓
▓▒▒▒▓▒▒▒▒▒▒▒▒▒▒▓▓▓▓▓▓▓███▓█████████████████████████▒▒▓▓▓▓▓▒▒▓▓▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▓
█▒▒▒▒▒▒▒▒▒▒▒▒▓▓▓▓▓▒▒▓████▓▒███████████████████████▒▒▒▒▒▓▓▒▒▓▓▓▒▒▒▒▒▒▒▒▒▒▒▓▓▒▒▒▒
▓▒▒█▓▒▒▒▒▓▓▓▓▓▒▒▒▓▒▒▒▓▓█▓▓▒▓███████████████████▓▓▒▒▒▒▒▒▒▓▓▓▓▒▓▒▒▒▒▒▒▒▒▒▒▒▒▒▓▒▒▒
▓▒▓██▒▒▓▓▓▒▒▓▒▒▓▓▒▒▒▒▓▓▓▓▓▒▓▓██████████████████▓▒▒▒▒▒▒▒▒▓▓▓▓▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▓▓▒▒
```
