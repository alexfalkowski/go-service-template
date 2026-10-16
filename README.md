[![CircleCI](https://circleci.com/gh/alexfalkowski/go-service-template.svg?style=svg)](https://circleci.com/gh/alexfalkowski/go-service-template)
[![codecov](https://codecov.io/gh/alexfalkowski/go-service-template/graph/badge.svg)](https://codecov.io/gh/alexfalkowski/go-service-template)
[![Go Reference](https://pkg.go.dev/badge/github.com/alexfalkowski/go-service-template.svg)](https://pkg.go.dev/github.com/alexfalkowski/go-service-template)
[![Stability: Active](https://masterminds.github.io/stability/active.svg)](https://masterminds.github.io/stability/active.html)

# go-service-template

<!-- repository:description -->
A template for Go services with HTTP and gRPC transports.
<!-- /repository:description -->

<!--
Service owners and documentation tools: expand the introduction with the intended
users, the problem this service solves, and the responsibilities it owns. Use
repository metadata for identity, api/ for RPC contracts, internal/ for behavior,
test/.config/server.yaml for local configuration, and test/features/ for verified
usage. Add domain request samples and design decisions when supported by code or
owner input. Link to InfraOps for deployment details; do not invent infrastructure,
dependencies, or rationale. Remove this comment once the service docs are complete.
-->

<!-- template:setup -->
## Create a service

Create the repository through InfraOps using this template, with its final name
and description. Clone it into a directory matching the repository name, then run
from the clone's root:

```sh
gh auth login
tools/init
```

Initialization requires Bash 4+, GNU sed, Git, the GitHub CLI, and the development
tools below. It reads the name, owner, and description from the GitHub repository
at `origin`, then updates the module path, service names, README, and badges.
An empty description or a checkout directory that differs from the repository
name stops initialization before any files are changed.

The script installs dependencies, regenerates API code, formats the project, and
removes `tools/` and this section. It then creates an initialization branch and
runs `make ready`, which commits, pushes, opens a pull request, and enables
auto-merge.
<!-- /template:setup -->

## Development

Use the Go and Ruby toolchain provided by the
[CI image](.circleci/config.yml). Install Buf for protobuf generation; the
development workflow also uses betteralign, golangci-lint, and RuboCop. See
`make help` for supported commands and their prerequisites.

Initialize the shared tooling before using Make, then install dependencies and
build the server:

```sh
git submodule sync && git submodule update --init
make dep
make build
```

## Server

Start the server with the local development configuration:

```sh
./go-service-template server -config file:test/.config/server.yaml
```

The local configuration listens for HTTP on port `11000` and gRPC on port
`12000`. It exports traces to an OTLP collector at `http://localhost:4318`.
See [the configuration](test/.config/server.yaml) for transport timeouts, rate
limits, logging, and telemetry settings.

The initial [gRPC contract](api/example/v1/service.proto) defines
`example.v1.Service/Hello`, with a message field in the request and response.
Implement the service's domain behavior and document a working request here.
Regenerate Go and Ruby bindings with `make proto-generate` after changing the
contract.

## Health

The server exposes service health, liveness, and readiness over HTTP:

```sh
curl --fail http://localhost:11000/go-service-template/healthz
curl --fail http://localhost:11000/go-service-template/livez
curl --fail http://localhost:11000/go-service-template/readyz
```

The initial checks use a no-op check for liveness and readiness, and an online
check for service health. The gRPC health service checks `example.v1.Service`.
See [the health wiring](internal/health/health.go) and
[health scenarios](test/features/health/) when adding dependency checks.

```yaml
health:
  duration: 1s # Interval between checks.
  timeout: 1s # Timeout for the online check.
```

## Deployment

Container builds use the shared
[Go service Dockerfile](bin/build/docker/go/Dockerfile). Deployment configuration
is managed in [InfraOps](https://github.com/alexfalkowski/infraops); link to the
service's configuration there once it is provisioned.

## References

- [API contract](api/example/v1/service.proto)
- [Go package documentation](https://pkg.go.dev/github.com/alexfalkowski/go-service-template)
- [Releases](https://github.com/alexfalkowski/go-service-template/releases)
- [Shared tooling](bin/README.md)
