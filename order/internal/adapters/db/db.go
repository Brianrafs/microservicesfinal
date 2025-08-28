package db

import (
	"fmt"
	"time"

	"github.com/brianrafs/microservicesfinal/order/internal/application/core/domain"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Estruturas de banco
type Order struct {
	gorm.Model
	CustomerID int64
	Status     string
	OrderItems []OrderItem `gorm:"foreignKey:OrderID"`
}

type OrderItem struct {
	gorm.Model
	ProductCode string
	UnitPrice   float32
	Quantity    int32
	OrderID     uint
}

type Payment struct {
	gorm.Model
	CustomerID int64
	OrderID    int64
	TotalPrice float32
	Status     string
}

// Adapter do DB
type Adapter struct {
	db *gorm.DB
}

// Inicializa o Adapter e faz migração
func NewAdapter(dataSourceUrl string) (*Adapter, error) {
	db, openErr := gorm.Open(mysql.Open(dataSourceUrl), &gorm.Config{})
	if openErr != nil {
		return nil, fmt.Errorf("db connection error: %v", openErr)
	}

	// Migra as tabelas
	err := db.AutoMigrate(&Order{}, &OrderItem{}, &Payment{})
	if err != nil {
		return nil, fmt.Errorf("db migration error: %v", err)
	}

	return &Adapter{db: db}, nil
}

// Get retorna um pedido
func (a *Adapter) Get(id string) (domain.Order, error) {
	var orderEntity Order
	res := a.db.Preload("OrderItems").First(&orderEntity, id)

	var orderItems []domain.OrderItem
	for _, item := range orderEntity.OrderItems {
		orderItems = append(orderItems, domain.OrderItem{
			ProductCode: item.ProductCode,
			UnitPrice:   item.UnitPrice,
			Quantity:    item.Quantity,
		})
	}

	order := domain.Order{
		ID:         int64(orderEntity.ID),
		CustomerID: orderEntity.CustomerID,
		Status:     orderEntity.Status,
		OrderItems: orderItems,
		CreatedAt:  orderEntity.CreatedAt.Unix(),
	}
	return order, res.Error
}

// Save salva um pedido
func (a *Adapter) Save(order *domain.Order) error {
	var orderItems []OrderItem
	for _, item := range order.OrderItems {
		orderItems = append(orderItems, OrderItem{
			ProductCode: item.ProductCode,
			UnitPrice:   item.UnitPrice,
			Quantity:    item.Quantity,
		})
	}

	orderModel := Order{
		CustomerID: order.CustomerID,
		Status:     order.Status,
		OrderItems: orderItems,
	}

	res := a.db.Create(&orderModel)
	if res.Error == nil {
		order.ID = int64(orderModel.ID)
	}
	return res.Error
}

// UpdateStatus atualiza o status do pedido
func (a *Adapter) UpdateStatus(order *domain.Order) error {
	res := a.db.Model(&Order{}).Where("id = ?", order.ID).Update("status", order.Status)
	return res.Error
}

// SavePayment salva um pagamento
func (a *Adapter) SavePayment(payment *domain.Payment) error {
	paymentModel := Payment{
		CustomerID: payment.CustomerID,
		OrderID:    payment.OrderId,
		TotalPrice: payment.TotalPrice,
		Status:     payment.Status,
		Model:      gorm.Model{CreatedAt: time.Now()},
	}
	res := a.db.Create(&paymentModel)
	if res.Error == nil {
		payment.ID = int64(paymentModel.ID)
	}
	return res.Error
}
