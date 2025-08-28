FROM golang:1.24.3 AS builder
WORKDIR /app

# Copia os protos
COPY microservices-protofinal/golang/shipping ./microservices-protofinal/golang/shipping

# Copia o serviço inteiro
COPY microservicesfinal/shipping ./shipping

WORKDIR /app/shipping

# Ajusta o replace do módulo
RUN go mod edit -replace github.com/brianrafs/microservices-protofinal/golang/shipping=../microservices-protofinal/golang/shipping

# Resolve dependências
RUN go mod tidy

# Build
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /out/shipping ./cmd

# Imagem final
FROM gcr.io/distroless/base-debian12
COPY --from=builder /out/shipping /shipping
ENV SHIPPING_PORT=50052
EXPOSE 50052
USER nonroot:nonroot
ENTRYPOINT ["/shipping"]
