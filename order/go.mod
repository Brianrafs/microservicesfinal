module github.com/brianrafs/microservicesfinal/order //order

go 1.24.3

require (
	github.com/brianrafs/microservices-protofinal/golang/order v0.1.1
	github.com/brianrafs/microservices-protofinal/golang/shipping v0.1.1
	google.golang.org/grpc v1.75.0
)

require (
	filippo.io/edwards25519 v1.1.0 // indirect
	github.com/go-sql-driver/mysql v1.8.1 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
)

require (
	golang.org/x/net v0.41.0 // indirect
	golang.org/x/sys v0.33.0 // indirect
	golang.org/x/text v0.26.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250707201910-8d1bb00bc6a7 // indirect
	google.golang.org/protobuf v1.36.8 // indirect
	gorm.io/driver/mysql v1.6.0
	gorm.io/gorm v1.30.1
)

require (
	github.com/brianrafs/microservices-protofinal/golang/payment v0.1.1
	github.com/grpc-ecosystem/go-grpc-middleware v1.4.0
)

replace github.com/brianrafs/microservices-protofinal/golang/payment => ../../microservices-protofinal/golang/payment

replace github.com/brianrafs/microservices-protofinal/golang/shipping => ../../microservices-protofinal/golang/shipping

replace github.com/brianrafs/microservices-protofinal/golang/order => ../../microservices-protofinal/golang/order
