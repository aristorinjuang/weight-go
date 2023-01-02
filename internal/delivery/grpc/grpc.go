package grpc

import (
	"fmt"
	"log"
	"net"

	handler "github.com/aristorinjuang/weight-go/internal/delivery/grpc/weight"
	"github.com/aristorinjuang/weight-go/internal/repository"
	pb "github.com/aristorinjuang/weight-go/tools/grpc"
	"google.golang.org/grpc"
)

type Server struct {
	port             int
	weightRepository repository.WeightRepository
}

func (s *Server) Run() error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", s.port))
	if err != nil {
		return err
	}

	server := grpc.NewServer()
	pb.RegisterWeightServiceServer(server, handler.New(s.weightRepository))

	log.Printf("The gRPC server is running on port %d.\n", s.port)

	return server.Serve(lis)
}

func New(port int, weightRepository repository.WeightRepository) *Server {
	return &Server{
		port:             port,
		weightRepository: weightRepository,
	}
}
