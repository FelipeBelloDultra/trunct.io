# Builder
FROM golang:1.23.2 AS builder
RUN apt-get update
WORKDIR /go/src
ENV PATH="/go/bin:${PATH}"
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN GOOS=linux go build -ldflags="-s -w" -o ./bin/trunctio ./cmd/api

# Development
FROM builder AS development
RUN go install github.com/air-verse/air@latest && \
    go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest && \
    go install github.com/jackc/tern/v2@latest
CMD [ "tail", "-f", "/dev/null" ]

# Production
FROM alpine:3.20 AS production
RUN apk --no-cache add ca-certificates && \
    adduser -D -u 1001 produser
USER produser
WORKDIR /app
COPY --from=builder /go/src/bin/trunctio /app/
EXPOSE 3333
CMD [ "/app/trunctio" ]
