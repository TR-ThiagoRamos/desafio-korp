# Definindo a versao da build
FROM golang:1.22-alpine AS builder

WORKDIR /app

# Copia arquivos de dependência e baixa
COPY go.mod go.sum ./
RUN go mod download

# Copia o código-fonte e compila
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Estágio final de execução
FROM alpine:latest
WORKDIR /root/

# Copia o binário compilado
COPY --from=builder /app/main .

# Expõe a porta 8080 e define o como iniciar
EXPOSE 8080
CMD ["./main"]