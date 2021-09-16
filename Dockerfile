ARG GO_VERSION=1.13

FROM golang:${GO_VERSION} AS builder

# RUN apk update && apk add alpine-sdk git && rm -rf /var/cache/apk/*
# RUN rm -rf /var/cache/apk/* && \
#     rm -rf /tmp/* && \
#     apk update
RUN mkdir -p /api
WORKDIR /api

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN export CGO_ENABLED=0 && go build -o ./app 

FROM alpine:3.11.3

RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*

RUN mkdir -p /api
WORKDIR /api
COPY --from=builder /api/.env .
COPY --from=builder /api/app .
COPY --from=builder /api/test.db .

EXPOSE 8080

ENTRYPOINT ["./app"]
