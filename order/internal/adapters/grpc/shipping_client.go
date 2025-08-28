package grpc

import (
    "context"
    "fmt"
    "os"
    "time"

    "google.golang.org/grpc"
    "google.golang.org/grpc/credentials/insecure"

    shippingpb "github.com/brianrafs/microservices-protofinal/golang/shipping"
)

type ShippingClient struct {
    client shippingpb.ShippingClient
    conn   *grpc.ClientConn
}

func NewShippingClient() (*ShippingClient, error) {
    addr := os.Getenv("SHIPPING_ADDR")
    if addr == "" {
        addr = "localhost:50052"
    }
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    conn, err := grpc.DialContext(ctx, addr, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
    if err != nil {
        return nil, fmt.Errorf("dial shipping: %w", err)
    }
    return &ShippingClient{
        client: shippingpb.NewShippingClient(conn),
        conn:   conn,
    }, nil
}

func (c *ShippingClient) Close() error {
    if c.conn != nil { return c.conn.Close() }
    return nil
}

func (c *ShippingClient) Estimate(ctx context.Context, orderID int64, items []struct{ProductCode string; Quantity int64}) (int64, error) {
    req := &shippingpb.ShippingRequest{ OrderId: orderID }
    for _, it := range items {
        req.Items = append(req.Items, &shippingpb.ShippingItem{ ProductCode: it.ProductCode, Quantity: it.Quantity })
    }
    resp, err := c.client.GetDeliveryEstimate(ctx, req)
    if err != nil { return 0, err }
    return resp.EstimatedDays, nil
}
