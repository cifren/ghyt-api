dk=docker-compose
dk_run=$(dk) run --rm
go_dk=$(dk_run) go
go=$(go_dk) go
#project_path=github.com/cifren/ghyt
project_path=ghyt

go.gopath:
	$(go_dk) env | grep GOPATH

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

go.get:
	$(go) get -d -v $(project_path)

go.install: go.get
	$(go) install -v $(project_path)

go.build: go.get
	$(go) build -v $(project_path)
