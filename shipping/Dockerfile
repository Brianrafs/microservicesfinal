FROM golang:1.24.3 AS builder
WORKDIR /app
COPY . .
COPY ../../microservices-proto/golang/shipping /app/microservices-proto/golang/shipping
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /out/shipping ./cmd

FROM gcr.io/distroless/base-debian12
COPY --from=builder /out/shipping /shipping
ENV SHIPPING_PORT=50052
EXPOSE 50052
USER nonroot:nonroot
ENTRYPOINT ["/shipping"]
