# ghyt-api

## Description

This package allow you to quickly put on a system that operate actions between Youtrack and Github, using
a webserver receiving hooks from github and sending actions to Youtrack.

It needs for that a configuration that defines jobs, conditions and actions that will be played.

## Pre-install

You need :
- make
- docker
- docker-compose

## Run your test

```bash
$ make test@run
```

## Run only some tests

```bash
$ make test@run-specific file="github.com/cifren/ghyt-api/youtrack/repository"
```
