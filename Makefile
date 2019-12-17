dk=docker-compose
dk_run=$(dk) run --rm
## GOLANG
dkr_go=$(dk_run) go
r_go=$(dkr_go) go
## NGROK
ngrok=$(dk_run) ngrok ngrok

## PROJECT
project_name=github.com/cifren/ghyt
project_path=/go/src/github.com/cifren/ghyt

## DOCKER
sh:
	docker-compose run --rm go bash

build:
	$(dk) build go

stop:
	$(dk) down --remove-orphan

start:
	$(dk) up -d go

restart: stop start

logs:
	$(dk) logs -f

install: stop go.deps-get go.get start

## GO
go.deps-get:
	$(r_go) get -v github.com/codegangsta/gin

go.get:
	$(r_go) get -d -v

go.install:
	$(r_go) install -v $(project_name)

go.build:
	$(r_go) build -v $(project_name)

## SERVICES
ngrok.up:
	$(ngrok) http go:8080

test-install:
	$(r_go) get -v -t github.com/cifren/...
	$(r_go) test -i -v github.com/cifren/...

test: test-install
	$(r_go) test -cover -race github.com/cifren/youtrack/repository

test-cover: test-install
	$(r_go) test -cover github.com/cifren/...

## New world ######

test@run:
	$(r_go) test -race ./...
