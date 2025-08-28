module github.com/brianrafs/microservices/shipping

go 1.23.0

toolchain go1.24.3

require (
	github.com/brianrafs/microservices-proto/golang/shipping v0.1.0
	google.golang.org/grpc v1.75.0
	google.golang.org/protobuf v1.36.6
)

require (
	golang.org/x/net v0.41.0 // indirect
	golang.org/x/sys v0.33.0 // indirect
	golang.org/x/text v0.26.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250707201910-8d1bb00bc6a7 // indirect
)

replace github.com/brianrafs/microservices-proto/golang/shipping => ../../microservices-proto/golang/shipping
