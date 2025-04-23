FROM golang:1.21-alpine

WORKDIR /app

# Instalar dependências do sistema
RUN apk add --no-cache git

# Copiar os arquivos do projeto
COPY . .

# Baixar dependências
RUN go mod download

# Compilar a aplicação
RUN go build -o server ./cmd/server/main.go

# Expor as portas necessárias
EXPOSE 8080
EXPOSE 50051

# Executar a aplicação
CMD ["./server"] 