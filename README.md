# Clean Architecture - Order Service

Este projeto implementa um serviço de Orders utilizando Clean Architecture, oferecendo múltiplos endpoints para listagem de pedidos através de diferentes protocolos.

## Portas dos Serviços

- REST API: `http://localhost:8080`
- gRPC: `localhost:50051`
- GraphQL: `http://localhost:8080/query`

## Requisitos

- Docker
- Docker Compose
- Go 1.21 ou superior

## Como Executar

1. Clone o repositório:
```bash
git clone https://github.com/mersonff/desafio-03-go-expert.git
cd desafio-03-go-expert
```

2. Inicie os serviços com Docker Compose:
```bash
docker compose up -d
```

3. Execute a aplicação:
```bash
go run cmd/server/main.go
```

## Endpoints Disponíveis

### REST
- Listar Orders: `GET /order`

### GraphQL
- Query para listar orders:
```graphql
query {
  orders {
    id
    price
    tax
    final_price
  }
}
```

### gRPC
O serviço gRPC está disponível na porta 50051 e pode ser testado usando ferramentas como grpcurl:
```bash
grpcurl -plaintext localhost:50051 OrderService.ListOrders
```

## Estrutura do Projeto

```
.
├── api/            # Definições de API (GraphQL, gRPC)
├── cmd/            # Pontos de entrada da aplicação
├── configs/        # Configurações
├── internal/       # Código interno da aplicação
│   ├── domain/    # Entidades e regras de negócio
│   ├── infra/     # Implementações de infraestrutura
│   └── usecase/   # Casos de uso da aplicação
└── pkg/           # Pacotes públicos reutilizáveis
```

## Banco de Dados

O projeto utiliza MySQL como banco de dados, que é automaticamente configurado através do Docker Compose. As credenciais padrão são:

- Host: localhost
- Porta: 3306
- Usuário: root
- Senha: root
- Database: orders

## Testes

Para executar os testes:
```bash
go test ./...
``` 