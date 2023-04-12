package http

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/aristorinjuang/weight-go/internal/domain/repository"
	pb "github.com/aristorinjuang/weight-go/tools/grpc"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Server struct {
	grpcPort         int
	port             int
	weightRepository repository.WeightRepository
}

func (s *Server) Run() error {
	conn, err := grpc.DialContext(
		context.Background(),
		fmt.Sprintf(":%d", s.grpcPort),
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return err
	}

	gwmux := runtime.NewServeMux()
	if err = pb.RegisterWeightServiceHandler(context.Background(), gwmux, conn); err != nil {
		return err
	}

	gwServer := &http.Server{
		Addr:    fmt.Sprintf(":%d", s.port),
		Handler: gwmux,
	}

	log.Printf("The HTTP server is running on port %d.", s.port)

	return gwServer.ListenAndServe()
}

func New(grpcPort, port int, weightRepository repository.WeightRepository) *Server {
	return &Server{
		grpcPort:         grpcPort,
		port:             port,
		weightRepository: weightRepository,
	}
}
