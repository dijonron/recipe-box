FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY services/user/go.mod services/user/go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o user-service ./services/user/cmd/main.go

FROM alpine:latest  

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/user-service .

# Copy .env file (optional if you need it inside the container)
# COPY services/user/.env .  # if needed at runtime

ENTRYPOINT ["./user-service"]