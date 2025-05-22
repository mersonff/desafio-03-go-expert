FROM golang:1.23-alpine

WORKDIR /app

# Instalar dependências necessárias
RUN apk add --no-cache gcc musl-dev

# Copiar arquivos de dependências
COPY go.mod go.sum ./
COPY .env ./
RUN go mod download && go mod tidy

# Copiar o código fonte
COPY . .

# Compilar a aplicação
RUN go build -o main ./cmd/ordersystem

# Expor as portas
EXPOSE 8080 50051 8081

# Comando para executar a aplicação
CMD ["./main"] 