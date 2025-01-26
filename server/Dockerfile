
FROM golang:1.23-alpine AS builder

WORKDIR /server
COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o api ./cmd/api

FROM alpine:3.14

WORKDIR /server

COPY --from=builder /server/api .

EXPOSE 50051
ENTRYPOINT [ "./api" ]