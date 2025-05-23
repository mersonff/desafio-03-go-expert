# Clean Architecture - Order System

Este é um sistema de pedidos implementado usando Clean Architecture, oferecendo múltiplas interfaces para interação: REST API, gRPC e GraphQL.

# Build OK

docker compose up -d --build
![Screenshot 2025-05-23 223430](https://github.com/user-attachments/assets/ef04d805-eabe-46b2-911a-06b58b00d524)

docker compose up --build
![image](https://github.com/user-attachments/assets/a64bfc38-c2cb-4174-acc9-f4fcf93bc01c)

## Tecnologias Utilizadas

- Go 1.23
- MySQL 8.0
- RabbitMQ 3
- Docker
- gRPC
- GraphQL

## Estrutura do Projeto

```
.
├── cmd/
│   ├── ordersystem/    # Aplicação principal
│   └── client/         # Cliente gRPC de teste
├── internal/
│   ├── domain/         # Entidades e regras de negócio
│   ├── usecase/        # Casos de uso
│   └── infra/          # Implementações de infraestrutura
│       ├── graph/      # GraphQL
│       ├── grpc/       # gRPC
│       └── web/        # REST API
└── pkg/                # Pacotes compartilhados
```

## Portas Utilizadas

- REST API: `:8080`
- gRPC: `:50051`
- GraphQL: `:8081`
- MySQL: `:3306`
- RabbitMQ: `:5672` (AMQP) e `:15672` (Management UI)

## Requisitos

- Docker
- Docker Compose

## Como Executar

1. Clone o repositório:
```bash
git clone https://github.com/mersonff/desafio-03-go-expert.git
cd desafio-03-go-expert
```

2. Inicie a aplicação:
```bash
docker compose up -d
```

A aplicação estará disponível em:
- REST API: http://localhost:8080
- GraphQL Playground: http://localhost:8081
- gRPC: localhost:50051

## Testando os Endpoints

### REST API

1. Criar um pedido:
```http
POST http://localhost:8080/order
Content-Type: application/json

{
    "id": "123",
    "price": 100.0,
    "tax": 10.0
}
```

2. Listar pedidos:
```http
GET http://localhost:8080/orders?page=1&limit=10
```

### gRPC

1. Criar um pedido:
```bash
grpcurl -plaintext -d '{"id": "123", "price": 100.0, "tax": 10.0}' localhost:50051 pb.OrderService/CreateOrder
```

2. Listar pedidos:
```bash
grpcurl -plaintext -d '{"page": 1, "limit": 10}' localhost:50051 pb.OrderService/ListOrders
```

### GraphQL

Acesse o playground GraphQL em `http://localhost:8081` e use as seguintes queries:

1. Criar um pedido:
```graphql
mutation {
  createOrder(input: {
    id: "123"
    price: 100.0
    tax: 10.0
  }) {
    id
    price
    tax
    finalPrice
  }
}
```

2. Listar pedidos:
```graphql
query {
  listOrders(page: 1, limit: 10) {
    id
    price
    tax
    finalPrice
  }
}
```

## Estrutura do Banco de Dados

O banco de dados MySQL é criado automaticamente com a seguinte estrutura:

```sql
CREATE TABLE orders (
    id VARCHAR(255) PRIMARY KEY,
    price FLOAT NOT NULL,
    tax FLOAT NOT NULL,
    final_price FLOAT NOT NULL
);
```

## Eventos

O sistema utiliza RabbitMQ para publicar eventos quando um pedido é criado. O evento é publicado no formato:

```json
{
    "id": "123",
    "price": 100.0,
    "tax": 10.0,
    "final_price": 110.0
}
```

## Contribuindo

1. Faça um fork do projeto
2. Crie uma branch para sua feature (`git checkout -b feature/nova-feature`)
3. Commit suas mudanças (`git commit -m 'Adiciona nova feature'`)
4. Push para a branch (`git push origin feature/nova-feature`)
5. Abra um Pull Request 
