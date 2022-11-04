ARG GO_VERSION=1.16.6

FROM golang:${GO_VERSION}-alpine AS builder

RUN go env -w GOPROXY=direct
RUN apk add --no-cache git
RUN apk --no-cache add ca-certificates && update-ca-certificates

WORKDIR /src

COPY ./go.mod ./go.sum ./
RUN go mod download

COPY events events
COPY repository repository
COPY database database
COPY search search
COPY models models
COPY user-service user-service
COPY user-client-service user-client-service
COPY user-external-worker-service user-external-worker-service
COPY user-of-client-service user-of-client-service
COPY user-southpandas-service user-southpandas-service
COPY query-service query-service
COPY pusher-service pusher-service 

RUN go install ./...

FROM alpine:3.11
WORKDIR /usr/bin

COPY --from=builder /go/bin .