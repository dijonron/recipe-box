FROM golang:1.23-alpine AS builder

WORKDIR /service
COPY ../../../go.mod ../../../go.sum ./
RUN go mod download

COPY . .

RUN go build -o app ./services/ingredient_service/v1

FROM alpine:3.14

WORKDIR /service

COPY --from=builder /service/app .

EXPOSE 50051
ENTRYPOINT [ "./app" ]