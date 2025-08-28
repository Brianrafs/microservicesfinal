FROM golang:1.24.3 AS builder
WORKDIR /app

# Copiar a pasta order
COPY microservicesfinal/order ./order

# Copiar apenas o que é necessário do microservices-protofinal
COPY microservices-protofinal/golang/order ./microservices-protofinal/golang/order
COPY microservices-protofinal/golang/payment ./microservices-protofinal/golang/payment

ENV GOPROXY=https://proxy.golang.org,direct

WORKDIR /app/order
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /out/order ./cmd

FROM gcr.io/distroless/base-debian12
COPY --from=builder /out/order /order
ENV APPLICATION_PORT=3000
EXPOSE 3000
USER nonroot:nonroot
ENTRYPOINT ["/order"]
