# Stage 1: Build
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Copiar go.mod e go.sum (se existir)
COPY go.mod .
RUN go mod download

# Copiar todo o código
COPY . .


RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Stage 2: Runtime
FROM alpine:latest


RUN apk --no-cache add ca-certificates

WORKDIR /root/


COPY --from=builder /app/main .


EXPOSE 8080


CMD ["./main"]
