dk=docker-compose
dk_run=$(dk) run --rm
## GOLANG
go_dk=$(dk_run) go
go=$(go_dk) go
## NGROK
ngrok=$(dk_run) ngrok ngrok

## PROJECT
project_path=github.com/cifren/ghyt

## DOCKER
console:
	docker-compose run --rm go bash

install:
	echo "intall"

build:
	$(dk) build go

down:
	$(dk) down --remove-orphan

up: go.install
	$(dk) up go

## GO
go.get:
	$(go) get -d -v $(project_path)

go.install: go.get
	$(go) install -v $(project_path)

go.build: go.get
	$(go) build -v $(project_path)

## SCRIPTS
database-create: go.get
	$(go) run -v src/github.com/cifren/ghyt/internal/scripts/create_database.go

fixture: go.get
	$(go) run -v src/github.com/cifren/ghyt/internal/scripts/fixture.go

## SERVICES
ngrok.up:
	$(ngrok) http go:8080

