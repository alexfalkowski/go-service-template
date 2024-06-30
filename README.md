[![CircleCI](https://circleci.com/gh/alexfalkowski/go-service-template.svg?style=svg)](https://circleci.com/gh/alexfalkowski/go-service-template)
[![codecov](https://codecov.io/gh/alexfalkowski/go-service-template/graph/badge.svg?token=S9SPVVYQAY)](https://codecov.io/gh/alexfalkowski/go-service-template)
[![Go Report Card](https://goreportcard.com/badge/github.com/alexfalkowski/go-service-template)](https://goreportcard.com/report/github.com/alexfalkowski/go-service-template)
[![Go Reference](https://pkg.go.dev/badge/github.com/alexfalkowski/go-service-template.svg)](https://pkg.go.dev/github.com/alexfalkowski/go-service-template)
[![Stability: Active](https://masterminds.github.io/stability/active.svg)](https://masterminds.github.io/stability/active.html)

# Service

Make sure you add the name of the service what what it is.

## Background

Add a background.

### Why a service?

Why is it important to have a service.

## Server

Explain the server side of things.

## Health

The system defines a way to monitor all of it's dependencies.

To configure we just need the have the following configuration:

```yaml
health:
  duration: 1s (how often to check)
  timeout: 1s (when we should timeout the check)
```

## Deployment

Since we are advocating building microservices, you would normally use a [container orchestration system](https://newrelic.com/blog/best-practices/container-orchestration-explained). Here is what we recommend when using this system:
- You could have a global migration service or shard these services per [bounded context](https://martinfowler.com/bliki/BoundedContext.html).
- The client should be used as an [init container](https://kubernetes.io/docs/concepts/workloads/pods/init-containers/).

## Design

Add anything interesting about the design.

## Other Systems

Describe any other similar systems you took inspiration from.

## Development

If you would like to contribute, here is how you can get started.

### Structure

The project follows the structure in [golang-standards/project-layout](https://github.com/golang-standards/project-layout).

### Dependencies

Please make sure that you have the following installed:
- [Ruby](.ruby-version)
- [Golang](go.mod)

### Style

This project favours the [Uber Go Style Guide](https://github.com/uber-go/guide/blob/master/style.md)

### Setup

Check out [CI](.circleci/config.yml).

### Changes

To see what has changed, please have a look at `CHANGELOG.md`
