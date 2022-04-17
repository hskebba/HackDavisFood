FROM golang:1.14-alpine as builder

RUN apk update && apk upgrade && apk add --no-cache bash git openssh

RUN mkdir -p /go/src/HackDavisFood/services/data_api
# Install deps
RUN apk add bash ca-certificates git gcc g++ libc-dev

# Force the go compiler to use modules
ENV GO111MODULE=on

WORKDIR /go/src/HackDavisFood/services/data_api
ADD ./data_api/go.mod .
ADD ./data_api/go.sum .
RUN go mod download

WORKDIR /go/src/HackDavisFood/services/data_api
ADD ./orders /usr/local/go/src/HackDavisFood/services/orders
ADD ./inventory /usr/local/go/src/HackDavisFood/services/inventory

WORKDIR /go/src/HackDavisFood/services/data_api
ADD ./data_api .

WORKDIR /go/src/HackDavisFood/services/data_api/cmd/server
RUN go build -o data_api

FROM alpine:3.12
RUN apk --no-cache add ca-certificates && \
    GRPC_HEALTH_PROBE_VERSION=v0.3.1 && \
    wget -qO/bin/grpc_health_probe https://github.com/grpc-ecosystem/grpc-health-probe/releases/download/${GRPC_HEALTH_PROBE_VERSION}/grpc_health_probe-linux-amd64 && \
    chmod +x /bin/grpc_health_probe
WORKDIR /HackDavisFood
COPY --from=builder /go/src/HackDavisFood/services/data_api/cmd/server/data_api /HackDavisFood/data_api
ENTRYPOINT ./data_api
