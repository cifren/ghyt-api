# ghyt-api

## Description

This will install a webserver that will allow you to receive notification from external tools

Define your configuration containing conditions and actions

Depending on the conditions the webserver will execute the actions based on the event you wanted

## Pre-install

You need :
- docker
- docker-compose
- make

## Install the project for the first time

```bash
$ make install
```

## Usage - Run the project

```bash
$ make start
```

You can now access on http://localhost:5000

## See the logs

```bash
$ make logs
```
