FROM golang:1.23.2-alpine AS builder
RUN apk update
WORKDIR /src/trunctio
COPY go.mod go.sum ./
RUN go mod download
COPY . .
ENV GOOS=linux CGO_ENABLED=0 GOARCH=amd64 GO111MODULE=on
RUN go build -o trunctio ./cmd/api

FROM alpine:3.20 AS production
RUN apk --no-cache add ca-certificates && \
    adduser -D -u 1001 produser
USER produser
WORKDIR /app
COPY --from=builder /src/trunctio/trunctio /app/
EXPOSE 3333
CMD [ "/app/trunctio" ]
