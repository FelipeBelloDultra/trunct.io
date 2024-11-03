FROM golang:1.23.2-alpine as base
RUN apk update
WORKDIR /src/trunctio
COPY go.mod go.sum ./
ENV GOOS=linux CGO_ENABLED=0 GOARCH=amd64 GO111MODULE=on
RUN go build -o trunctio ./cmd/api

FROM alpine:3.20 as production
COPY --from=base /src/trunctio/trunctio ./
EXPOSE 3333
CMD [ "./trunctio" ]
