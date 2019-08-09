dk=docker-compose
dk_run=$(dk) run --rm
go_dk=$(dk_run) go
go=$(go_dk) go

go.gopath:
	$(go_dk) env | grep GOPATH

console:
	docker-compose run --rm go bash

install:
	echo "intall"

build:
	$(dk) build go

up: go.install
	$(dk) up go

go.get:
	$(go) get -d -v main

go.install: go.get
	$(go) install -v main

go.build: go.get
	$(go) build -v main