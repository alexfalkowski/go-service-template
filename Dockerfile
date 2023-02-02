FROM golang:1.20.0-bullseye AS build

ARG version=latest

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./
RUN go build -ldflags="-X 'github.com/alexfalkowski/go-service-template/cmd.Version=${version}'" -a -o go-service-template main.go

FROM gcr.io/distroless/base-debian11

WORKDIR /

COPY --from=build /app/go-service-template /go-service-template
ENTRYPOINT ["/go-service-template"]
