package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strconv"

	interceptors "github.com/llarsson/grpc-caching-interceptors/server"
	pb "github.com/llarsson/value-service-estimator/value"
	"google.golang.org/grpc"
)

const (
	upstreamAddrKey = "VALUE_SERVICE_ADDR"
)

func main() {
	port, err := strconv.Atoi(os.Getenv("PROXY_LISTEN_PORT"))
	if err != nil {
		log.Fatalf("PROXY_LISTEN_PORT cannot be parsed as integer")
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	estimator := new(interceptors.ConfigurableValidityEstimator)
	estimator.Initialize()

	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(estimator.UnaryServerInterceptor()))

	upstreamAddr, ok := os.LookupEnv(upstreamAddrKey)
	if !ok {
		log.Fatalf("Must supply upstream address as environment variable (%s)", upstreamAddrKey)
	}

	conn, err := grpc.Dial(upstreamAddr, grpc.WithUnaryInterceptor(estimator.UnaryClientInterceptor()), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Cannot connect to upstream %s : %v", upstreamAddr, err)
	}
	defer conn.Close()

	proxy := pb.ValueServiceProxy{Client: pb.NewValueServiceClient(conn)}
	pb.RegisterValueServiceServer(grpcServer, &proxy)

	log.Printf("Proxying service calls to %s", upstreamAddr)
	grpcServer.Serve(lis)
}
