dk=docker-compose
dk_run=$(dk) run --rm
## GOLANG
dkr_go=$(dk_run) go
r_go=$(dkr_go) go

## DOCKER
sh:
	docker-compose run --rm go bash

## TEST
test@run:
	$(r_go) test -race ./...

test@run-cover:
	$(r_go) test -race ./...

# Ex : $ make test@run-specific file="github.com/cifren/youtrack/repository"
test@run-specific:
	$(r_go) test -race $(file)

## DEV
dev@mod.clean:
	$(r_go) mod tidy
