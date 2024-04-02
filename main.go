package main

import (
	"io/ioutil"
	"net"
	"os"

	"github.com/jasonkolodziej/protoc-gen-go-redact/gateway"
	assetv1 "github.com/jasonkolodziej/protoc-gen-go-redact/gen/example/asset/v1"
	"github.com/jasonkolodziej/protoc-gen-go-redact/insecure"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/grpclog"
)

func registerServers(registrar *grpc.Server) {
	// TODO: To fill with your service server func
	assetv1.RegisterAssetServiceServer(registrar, assetv1.UnimplementedAssetServiceServer{})
}

func main() {
	// Adds gRPC internal logs. This is quite verbose, so adjust as desired!
	log := grpclog.NewLoggerV2(os.Stdout, ioutil.Discard, ioutil.Discard)
	grpclog.SetLoggerV2(log)

	addr := "0.0.0.0:10000"
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	s := grpc.NewServer(
		// TODO: Replace with your own certificate!
		grpc.Creds(credentials.NewServerTLSFromCert(&insecure.Cert)),
	)

	registerServers(s)

	// Serve gRPC Server
	log.Info("Serving gRPC on https://", addr)
	go func() {
		log.Fatal(s.Serve(lis))
	}()

	err = gateway.Run("dns:///" + addr)
	log.Fatalln(err)
}
