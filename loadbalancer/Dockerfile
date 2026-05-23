# Stage 1: Build
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Copiar go.mod
COPY go.mod .
RUN go mod download

# Copiar código
COPY . .

# Build da aplicação
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./src/main/main.go

# Stage 2: Runtime
FROM alpine:latest

# Instalar CA certificates para HTTPS
RUN apk --no-cache add ca-certificates

WORKDIR /app

# Copiar o binário do stage anterior
COPY --from=builder /app/main .

# Expor porta
EXPOSE 8080

# Comando para rodar
CMD ["./main"]
