.PHONY: vendor

include bin/build/make/service.mak

# Get git submodule.
git-submodule-get:
	git submodule sync
	git submodule update --init

# Update git submodule.
git-submodule-update:
	git submodule foreach git pull origin master

# Build release binary.
build:
	go build -race -ldflags="-X 'github.com/alexfalkowski/go-service-template/cmd.Version=latest'" -mod vendor -o go-service-template main.go

# Build test binary.
build-test:
	go test -race -ldflags="-X 'github.com/alexfalkowski/go-service-template/cmd.Version=latest'" -mod vendor -c -tags features -covermode=atomic -o go-service-template -coverpkg=./... github.com/alexfalkowski/go-service-template

sanitize-coverage:
	bin/quality/go/cov

# Get the HTML coverage for go.
html-coverage: sanitize-coverage
	go tool cover -html test/reports/final.cov

# Get the func coverage for go.
func-coverage: sanitize-coverage
	go tool cover -func test/reports/final.cov

# Send coveralls data.
goveralls: sanitize-coverage
	goveralls -coverprofile=test/reports/final.cov -service=circle-ci -repotoken=HrIzEFrP9Gr876IhiV1vYCObeXqpBzcmT

# Run go security checks.
go-sec:
	gosec -quiet -exclude-dir=test -exclude=G104 ./...

# Run security checks.
sec: go-sec

# Release to docker hub.
docker:
	bin/build/docker/push go-service-template

# Start the environment
start:
	bin/build/docker/env start

# Stop the environment
stop:
	bin/build/docker/env stop
