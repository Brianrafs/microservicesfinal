package ports

import (
	"context"

	"github.com/brianrafs/microservicesfinal/order/internal/application/core/domain"
)

// ApplicationPort define os métodos que o Adapter gRPC do Order precisa
type ApplicationPort interface {
	SaveOrder(ctx context.Context, order *domain.Order) error
	UpdateOrderStatus(ctx context.Context, order *domain.Order) error
}
