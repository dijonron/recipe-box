FROM golang:1.23-alpine AS builder

WORKDIR /service
COPY ../go.mod ../go.sum ./
RUN go mod download

COPY . .

RUN go build -o gateway ./api/gateway.go

FROM alpine:3.14

WORKDIR /service

COPY --from=builder /service/gateway .

EXPOSE 3000
ENTRYPOINT [ "./gateway" ]