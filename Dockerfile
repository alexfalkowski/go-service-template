FROM golang:1.23rc1-bullseye AS build

ARG version=latest

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./
RUN go build -ldflags="-s -w -X 'github.com/alexfalkowski/go-service-template/cmd.Version=${version}'" -a -o go-service-template main.go

FROM gcr.io/distroless/base-debian12

WORKDIR /

COPY --from=build /app/go-service-template /go-service-template
ENTRYPOINT ["/go-service-template"]
