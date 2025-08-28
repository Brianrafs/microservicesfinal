package main

import (
    "context"
    "log"
    "net"
    "os"

    "google.golang.org/grpc"
    "google.golang.org/grpc/reflection"

    shippingpb "github.com/brianrafs/microservices-protofinal/golang/shipping"
)

type ShippingServer struct {
    shippingpb.UnimplementedShippingServer
}

func (s *ShippingServer) GetDeliveryEstimate(ctx context.Context, req *shippingpb.ShippingRequest) (*shippingpb.ShippingResponse, error) {
    var total int64 = 0
    for _, it := range req.Items {
        total += it.Quantity
    }
    days := int64(1)
    if total > 0 {
        extra := (total - 1)
        days += extra / 5
    }
    return &shippingpb.ShippingResponse{EstimatedDays: days}, nil
}

func main() {
    port := os.Getenv("SHIPPING_PORT")
    if port == "" { port = "50052" }
    lis, err := net.Listen("tcp", ":"+port)
    if err != nil { log.Fatalf("listen: %v", err) }
    s := grpc.NewServer()
    shippingpb.RegisterShippingServer(s, &ShippingServer{})
    if os.Getenv("ENV") == "development" {
        reflection.Register(s)
    }
    log.Printf("shipping service listening on :%s", port)
    if err := s.Serve(lis); err != nil {
        log.Fatalf("serve: %v", err)
    }
}
