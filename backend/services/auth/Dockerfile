FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY services/auth/go.mod services/auth/go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o auth-service ./services/auth/cmd/main.go

FROM alpine:latest  

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/auth-service .

ENTRYPOINT ["./auth-service"]