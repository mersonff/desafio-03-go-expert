package main

import (
	"database/sql"
	"log"
	"net"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	_ "github.com/go-sql-driver/mysql"
	"github.com/mersonff/desafio-03-go-expert/api/proto"
	"github.com/mersonff/desafio-03-go-expert/internal/infra/graphql"
	"github.com/mersonff/desafio-03-go-expert/internal/infra/grpc"
	"github.com/mersonff/desafio-03-go-expert/internal/infra/http/handlers"
	"github.com/mersonff/desafio-03-go-expert/internal/infra/repository"
	"github.com/mersonff/desafio-03-go-expert/internal/usecase"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	// Conectar ao banco de dados
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/orders?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Inicializar repositório
	orderRepository := repository.NewMySQLOrderRepository(db)

	// Inicializar caso de uso
	orderUseCase := usecase.NewOrderUseCase(orderRepository)

	// Inicializar handlers HTTP
	orderHandler := handlers.NewOrderHandler(orderUseCase)

	// Configurar rotas HTTP
	http.HandleFunc("/order", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			orderHandler.List(w, r)
		case http.MethodPost:
			orderHandler.Create(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// Configurar GraphQL
	srv := handler.NewDefaultServer(graphql.NewExecutableSchema(graphql.Config{Resolvers: graphql.NewResolver(orderUseCase)}))
	http.Handle("/query", srv)
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))

	// Iniciar servidor HTTP em uma goroutine
	go func() {
		log.Printf("HTTP server starting on :8080")
		if err := http.ListenAndServe(":8080", nil); err != nil {
			log.Fatal(err)
		}
	}()

	// Iniciar servidor gRPC
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	proto.RegisterOrderServiceServer(grpcServer, grpc.NewOrderServer(orderUseCase))
	reflection.Register(grpcServer)

	log.Printf("gRPC server starting on :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
