package api

import (
	"context"

	"github.com/brianrafs/microservicesfinal/order/internal/application/core/domain"
	"github.com/brianrafs/microservicesfinal/order/internal/ports"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Application struct {
	db      ports.DBPort
	payment ports.PaymentPort
}

func NewApplication(db ports.DBPort, payment ports.PaymentPort) *Application {
	return &Application{
		db:      db,
		payment: payment,
	}
}

func (a *Application) SaveOrder(ctx context.Context, order *domain.Order) error {
	return a.db.Save(order)
}

func (a *Application) UpdateOrderStatus(ctx context.Context, order *domain.Order) error {
	return a.db.UpdateStatus(order)
}

// Função de cobrança com tratamento de erros
func (a *Application) Charge(ctx context.Context, payment domain.Payment) (domain.Payment, error) {
	if payment.TotalPrice > 1000 {
		return domain.Payment{}, status.Errorf(codes.InvalidArgument, "Payment over 1000 is not allowed")
	}

	if err := a.db.SavePayment(&payment); err != nil {
		return domain.Payment{}, status.Errorf(codes.Internal, "Failed to save payment: %v", err)
	}

	payment.Status = "Paid"
	return payment, nil
}
