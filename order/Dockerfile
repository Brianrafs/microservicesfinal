FROM golang:1.24.3 AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /out/order ./cmd

FROM gcr.io/distroless/base-debian12
COPY --from=builder /out/order /order
ENV APPLICATION_PORT=3000
EXPOSE 3000
USER nonroot:nonroot
ENTRYPOINT ["/order"]
