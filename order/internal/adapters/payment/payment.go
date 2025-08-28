package payment_adapter

import (
	"context"
	"log"
	"time"

	pbPayment "github.com/brianrafs/microservices-proto/golang/payment/payment"
	grpc_retry "github.com/grpc-ecosystem/go-grpc-middleware/retry"
	"github.com/brianrafs/microservices/order/internal/application/core/domain"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

type Adapter struct {
	client pbPayment.PaymentClient
}

func NewAdapter(paymentServiceUrl string) (*Adapter, error) {
    opts := []grpc.DialOption{
        grpc.WithTransportCredentials(insecure.NewCredentials()),
        grpc.WithUnaryInterceptor(grpc_retry.UnaryClientInterceptor(
            grpc_retry.WithCodes(codes.Unavailable, codes.ResourceExhausted),
            grpc_retry.WithMax(5),
            grpc_retry.WithBackoff(grpc_retry.BackoffLinear(time.Second)),
        )),
    }

    conn, err := grpc.Dial(paymentServiceUrl, opts...)
    if err != nil {
        return nil, err
    }

    client := pbPayment.NewPaymentClient(conn)
    return &Adapter{client: client}, nil
}

func (a *Adapter) Charge(ctx context.Context, pay domain.Payment) (domain.Payment, error) {
    // Cria contexto com deadline de 2 segundos
    ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
    defer cancel()

    req := &pbPayment.CreatePaymentRequest{
        UserId:     pay.CustomerID,
        OrderId:    pay.OrderId,
        TotalPrice: pay.TotalPrice,
    }

    res, err := a.client.Create(ctx, req)
    if err != nil {
        if status.Code(err) == codes.DeadlineExceeded {
            log.Printf("Payment service timeout exceeded for order %d", pay.OrderId)
        }
        return domain.Payment{}, err
    }

    pay.ID = res.PaymentId
    pay.Status = "Paid"
    return pay, nil
}


