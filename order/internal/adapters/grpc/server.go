package grpc

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/brianrafs/microservices-proto/golang/order"
	"github.com/brianrafs/microservices/order/config"
	"github.com/brianrafs/microservices/order/internal/application/core/domain"
	"github.com/brianrafs/microservices/order/internal/ports"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

// Adapter implementa o servidor gRPC do Order
type Adapter struct {
	api  ports.ApplicationPort
	pay  ports.PaymentPort
	port int
	order.UnimplementedOrderServer
}

// NewAdapter cria uma instância do Adapter
func NewAdapter(api ports.ApplicationPort, pay ports.PaymentPort, port int) *Adapter {
	return &Adapter{
		api:  api,
		pay:  pay,
		port: port,
	}
}

// Create implementa o endpoint gRPC para criar um pedido
func (a *Adapter) Create(ctx context.Context, req *order.CreateOrderRequest) (*order.CreateOrderResponse, error) {
	log.Printf("Creating order for customer %d...", req.CustomerId)

	// Monta o domain.Order
	items := make([]domain.OrderItem, len(req.OrderItems))
	for i, it := range req.OrderItems {
		items[i] = domain.OrderItem{
			ProductCode: it.ProductCode,
			UnitPrice:   it.UnitPrice,
			Quantity:    it.Quantity,
		}
	}
	newOrder := domain.NewOrder(int64(req.CustomerId), items)

	// Verifica limite de itens
	if newOrder.TotalItems() > 50 {
		return nil, status.Errorf(codes.InvalidArgument, "Order cannot have more than 50 items")
	}

	// Salva pedido com status Pending
	if err := a.api.SaveOrder(ctx, &newOrder); err != nil {
		newOrder.Status = "Canceled"
		a.api.UpdateOrderStatus(ctx, &newOrder)
		return nil, status.Errorf(codes.Internal, "Failed to save order: %v", err)
	}

	// Tenta realizar cobrança
	payment := domain.NewPayment(newOrder.CustomerID, newOrder.ID, newOrder.TotalPrice())
	_, err := a.pay.Charge(ctx, payment)
	if err != nil {
		newOrder.Status = "Canceled"
		a.api.UpdateOrderStatus(ctx, &newOrder)
		if status.Code(err) == codes.InvalidArgument {
			return nil, err
		}
		return nil, status.Errorf(codes.Internal, "Payment failed: %v", err)
	}

	// Pedido aprovado
	newOrder.Status = "Paid"
	a.api.UpdateOrderStatus(ctx, &newOrder)

	return &order.CreateOrderResponse{
		OrderId: int32(payment.ID),
	}, nil
}

// Run inicia o servidor gRPC
func (a *Adapter) Run() {
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		log.Fatalf("failed to listen on port %d: %v", a.port, err)
	}

	grpcServer := grpc.NewServer()
	order.RegisterOrderServer(grpcServer, a)

	if config.GetEnv() == "development" {
		reflection.Register(grpcServer)
	}

	log.Printf("gRPC server running on port %d", a.port)
	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to serve gRPC: %v", err)
	}
}
