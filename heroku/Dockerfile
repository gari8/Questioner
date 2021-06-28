FROM golang:alpine as builder

RUN apk update \
  && apk add --no-cache git curl make gcc g++

WORKDIR /app
COPY ../go.mod .
COPY ../go.sum .

RUN go mod download
COPY ../ .

RUN GOOS=linux GOARCH=amd64 go build -o /main

FROM alpine:latest

COPY --from=builder /main .

ENV PORT=${PORT}
ENTRYPOINT ["/main"]
