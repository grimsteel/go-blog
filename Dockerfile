FROM golang:1.24 AS builder
ARG CGO_ENABLED=0
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN go build

FROM scratch
WORKDIR /app

COPY --from=builder /app/go-blog /app/server
COPY templates /app/templates
COPY static /app/static

ENV LISTEN_ADDRESS=0.0.0.0:8080
ENTRYPOINT ["/app/server"]
