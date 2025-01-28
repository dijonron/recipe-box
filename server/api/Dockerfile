FROM golang:1.23-alpine AS builder

WORKDIR /service
COPY ../go.mod ../go.sum ./
RUN go mod download

COPY . .

RUN go build -o api ./api

FROM alpine:3.14

WORKDIR /service

COPY --from=builder /service/api .

EXPOSE 3000
ENTRYPOINT [ "./api" ]