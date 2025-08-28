package main

import (
	"log"

	"github.com/brianrafs/microservicesfinal/order/config"
	"github.com/brianrafs/microservicesfinal/order/internal/adapters/db"
	"github.com/brianrafs/microservicesfinal/order/internal/adapters/grpc"
	payment_adapter "github.com/brianrafs/microservicesfinal/order/internal/adapters/payment"
	"github.com/brianrafs/microservicesfinal/order/internal/application/core/api"
)

func main() {
	// Conecta ao DB
	dbAdapter, err := db.NewAdapter(config.GetDataSourceURL())
	if err != nil {
		log.Fatalf("Failed to connect to database. Error: %v", err)
	}

	// Conecta ao Payment
	paymentAdapter, err := payment_adapter.NewAdapter(config.GetPaymentServiceUrl())
	if err != nil {
		log.Fatalf("Failed to initialize payment adapter. Error: %v", err)
	}

	// Cria aplicação
	application := api.NewApplication(dbAdapter, paymentAdapter)

	// Cria e inicia servidor gRPC
	grpcAdapter := grpc.NewAdapter(application, paymentAdapter, config.GetApplicationPort())
	grpcAdapter.Run()
}
