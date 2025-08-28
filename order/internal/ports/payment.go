package ports

import (
	"context"

	"github.com/brianrafs/microservices/order/internal/application/core/domain"
)

// PaymentPort define os métodos que o Adapter gRPC precisa para cobrar
type PaymentPort interface {
	Charge(ctx context.Context, payment domain.Payment) (domain.Payment, error)
}

