package weight

import (
	"context"
	"errors"
	"net"
	"testing"

	"github.com/aristorinjuang/weight-go/internal/aggregate"
	"github.com/aristorinjuang/weight-go/internal/entity"
	"github.com/aristorinjuang/weight-go/internal/repository"
	"github.com/aristorinjuang/weight-go/test/data"
	pb "github.com/aristorinjuang/weight-go/tools/grpc"
	"github.com/stretchr/testify/suite"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
	"google.golang.org/protobuf/types/known/emptypb"
)

type weightGRPCTest struct {
	suite.Suite
	weightRepository *repository.WeightRepositoryMock
	client           pb.WeightServiceClient
}

func (t *weightGRPCTest) dialer() func(context.Context, string) (net.Conn, error) {
	listener := bufconn.Listen(1024 * 1024)
	server := grpc.NewServer()

	pb.RegisterWeightServiceServer(server, New(t.weightRepository))

	go func() {
		server.Serve(listener)
	}()

	return func(context.Context, string) (net.Conn, error) {
		return listener.Dial()
	}
}

func (t *weightGRPCTest) SetupSuite() {
	t.weightRepository = new(repository.WeightRepositoryMock)

	conn, _ := grpc.DialContext(
		context.Background(),
		"bufnet",
		grpc.WithContextDialer(t.dialer()),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	t.client = pb.NewWeightServiceClient(conn)
}

func (t *weightGRPCTest) TestListWeights() {
	t.Run("success", func() {
		t.weightRepository.On("List").Return(aggregate.NewWeights(data.Weights), nil).Once()

		weights, err := t.client.ListWeights(context.Background(), &emptypb.Empty{})

		t.NotNil(weights)
		t.NoError(err)
	})

	t.Run("failed", func() {
		t.weightRepository.On("List").Return((*aggregate.Weights)(nil), errors.New("failed list weights")).Once()

		weights, err := t.client.ListWeights(context.Background(), &emptypb.Empty{})

		t.Nil(weights)
		t.Error(err)
	})
}

func (t *weightGRPCTest) TestCreateWeight() {
	t.Run("success", func() {
		t.weightRepository.On("Create", data.Weights[0]).Return(nil).Once()

		_, err := t.client.CreateWeight(context.Background(), &pb.WeightParams{
			Date: data.Weights[0].Date().Format("2006-01-02"),
			Max:  data.Weights[0].Max(),
			Min:  data.Weights[0].Min(),
		})

		t.NoError(err)
	})

	t.Run("failed", func() {
		t.weightRepository.On("Create", data.Weights[0]).Return(errors.New("failed to create a weight")).Once()

		_, err := t.client.CreateWeight(context.Background(), &pb.WeightParams{
			Date: data.Weights[0].Date().Format("2006-01-02"),
			Max:  data.Weights[0].Max(),
			Min:  data.Weights[0].Min(),
		})

		t.Error(err)
	})

	t.Run("failed create a weight", func() {
		_, err := t.client.CreateWeight(context.Background(), &pb.WeightParams{
			Date: "2023-01-01",
		})

		t.Error(err)
	})

	t.Run("failed to parse the date", func() {
		_, err := t.client.CreateWeight(context.Background(), &pb.WeightParams{})

		t.Error(err)
	})
}

func (t *weightGRPCTest) TestReadWeight() {
	t.Run("success", func() {
		t.weightRepository.On("Read", data.Weights[0].Date()).Return(data.Weights[0], nil).Once()

		weight, err := t.client.ReadWeight(context.Background(), &pb.WeightParams{
			Date: data.Weights[0].Date().Format("2006-01-02"),
		})

		t.NotNil(weight)
		t.NoError(err)
	})

	t.Run("failed", func() {
		t.weightRepository.On("Read", data.Weights[0].Date()).Return(
			(*entity.Weight)(nil),
			errors.New("failed to read a weight"),
		).Once()

		weight, err := t.client.ReadWeight(context.Background(), &pb.WeightParams{
			Date: data.Weights[0].Date().Format("2006-01-02"),
		})

		t.Nil(weight)
		t.Error(err)
	})

	t.Run("failed to parse the date", func() {
		weight, err := t.client.ReadWeight(context.Background(), &pb.WeightParams{})

		t.Nil(weight)
		t.Error(err)
	})
}

func (t *weightGRPCTest) TestUpdateWeight() {
	t.Run("success", func() {
		t.weightRepository.On("Update", data.Weights[0]).Return(nil).Once()

		_, err := t.client.UpdateWeight(context.Background(), &pb.WeightParams{
			Date: data.Weights[0].Date().Format("2006-01-02"),
			Max:  data.Weights[0].Max(),
			Min:  data.Weights[0].Min(),
		})

		t.NoError(err)
	})

	t.Run("failed", func() {
		t.weightRepository.On("Update", data.Weights[0]).Return(errors.New("failed update the weight")).Once()

		_, err := t.client.UpdateWeight(context.Background(), &pb.WeightParams{
			Date: data.Weights[0].Date().Format("2006-01-02"),
			Max:  data.Weights[0].Max(),
			Min:  data.Weights[0].Min(),
		})

		t.Error(err)
	})

	t.Run("failed create a weight", func() {
		_, err := t.client.UpdateWeight(context.Background(), &pb.WeightParams{
			Date: "2023-01-01",
		})

		t.Error(err)
	})

	t.Run("failed to parse the date", func() {
		weight, err := t.client.UpdateWeight(context.Background(), &pb.WeightParams{})

		t.Nil(weight)
		t.Error(err)
	})
}

func (t *weightGRPCTest) TestDeleteWeight() {
	t.Run("success", func() {
		t.weightRepository.On("Delete", data.Weights[0].Date()).Return(nil).Once()

		_, err := t.client.DeleteWeight(context.Background(), &pb.WeightParams{
			Date: data.Weights[0].Date().Format("2006-01-02"),
		})

		t.NoError(err)
	})

	t.Run("failed", func() {
		t.weightRepository.On("Delete", data.Weights[0].Date()).Return(errors.New("failed delete the weight")).Once()

		_, err := t.client.DeleteWeight(context.Background(), &pb.WeightParams{
			Date: data.Weights[0].Date().Format("2006-01-02"),
		})

		t.Error(err)
	})

	t.Run("failed to parse the date", func() {
		_, err := t.client.DeleteWeight(context.Background(), &pb.WeightParams{})

		t.Error(err)
	})
}

func TestWeightGRPC(t *testing.T) {
	suite.Run(t, new(weightGRPCTest))
}
