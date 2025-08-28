# Etapa 1: Build
FROM golang:1.24.3 AS builder
WORKDIR /app

# Copiar a pasta order
COPY microservicesfinal/order ./order

# Copiar todos os módulos do protofinal necessários
COPY microservices-protofinal/golang/order ./microservices-protofinal/golang/order
COPY microservices-protofinal/golang/payment ./microservices-protofinal/golang/payment
COPY microservices-protofinal/golang/shipping ./microservices-protofinal/golang/shipping

# Configurar proxy
ENV GOPROXY=https://proxy.golang.org,direct

# Build do order
WORKDIR /app/order
RUN go mod edit -replace github.com/brianrafs/microservices-protofinal/golang/order=../microservices-protofinal/golang/order
RUN go mod edit -replace github.com/brianrafs/microservices-protofinal/golang/payment=../microservices-protofinal/golang/payment
RUN go mod edit -replace github.com/brianrafs/microservices-protofinal/golang/shipping=../microservices-protofinal/golang/shipping


RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /out/order ./cmd

# Etapa 2: Runtime
FROM gcr.io/distroless/base-debian12
COPY --from=builder /out/order /order
ENV APPLICATION_PORT=3000
EXPOSE 3000
USER nonroot:nonroot
ENTRYPOINT ["/order"]
