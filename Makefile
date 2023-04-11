.PHONY: vendor

include bin/build/make/service.mak

# Build release binary.
build:
	go build -race -ldflags="-X 'github.com/alexfalkowski/go-service-template/cmd.Version=latest'" -mod vendor -o go-service-template main.go

# Build test binary.
build-test:
	go test -race -ldflags="-X 'github.com/alexfalkowski/go-service-template/cmd.Version=latest'" -mod vendor -c -tags features -covermode=atomic -o go-service-template -coverpkg=./... github.com/alexfalkowski/go-service-template

# Release to docker hub.
docker:
	bin/build/docker/push go-service-template
