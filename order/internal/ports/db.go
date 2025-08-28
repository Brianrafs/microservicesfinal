package ports

import "github.com/brianrafs/microservicesfinal/order/internal/application/core/domain"

type DBPort interface {
	Get(id string) (domain.Order, error)
	Save(*domain.Order) error
	UpdateStatus(order *domain.Order) error
	SavePayment(payment *domain.Payment) error 
}
