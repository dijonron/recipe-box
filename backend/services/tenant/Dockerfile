FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY services/tenant/go.mod services/tenant/go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o tenant-service ./services/tenant/cmd/main.go

FROM alpine:latest  

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/tenant-service .

ENTRYPOINT ["./tenant-service"]