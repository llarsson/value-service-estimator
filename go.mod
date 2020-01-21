module github.com/llarsson/value-service-estimator

go 1.13

require (
	github.com/golang/protobuf v1.3.2
	github.com/llarsson/grpc-caching-interceptors v0.0.1
	google.golang.org/grpc v1.26.0
)

replace github.com/llarsson/grpc-caching-interceptors => ../grpc-caching-interceptors
