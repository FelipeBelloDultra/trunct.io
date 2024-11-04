# Builder
FROM golang:1.23.2-alpine AS builder
RUN apk update && \
    adduser -D -u 1001 appuser
USER appuser
WORKDIR /src/trunctio
COPY go.mod go.sum ./
RUN go mod download
COPY . .
ENV GOOS=linux CGO_ENABLED=0 GOARCH=amd64 GO111MODULE=on
RUN go build -o ./bin/trunctio ./cmd/api

# Development
FROM builder AS development
USER appuser
RUN go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest && \
    go install github.com/jackc/tern/v2@latest
CMD [ "tail", "-f", "/dev/null" ]

# Production
FROM alpine:3.20 AS production
RUN apk --no-cache add ca-certificates && \
    adduser -D -u 1001 produser
USER produser
WORKDIR /app
COPY --from=builder /src/trunctio/bin/trunctio /app/
EXPOSE 3333
CMD [ "/app/trunctio" ]
