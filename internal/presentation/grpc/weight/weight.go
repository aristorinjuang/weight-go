package weight

import (
	"context"
	"time"

	"github.com/aristorinjuang/weight-go/internal/domain/entity"
	"github.com/aristorinjuang/weight-go/internal/domain/repository"
	pb "github.com/aristorinjuang/weight-go/tools/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type handler struct {
	pb.UnimplementedWeightServiceServer
	repository repository.WeightRepository
}

func (h *handler) ListWeights(context.Context, *emptypb.Empty) (*pb.Weights, error) {
	entries, err := h.repository.List()
	if err != nil {
		return nil, err
	}

	return &pb.Weights{
		Weights:           weights{}.From(entries.Weights()),
		AverageMax:        entries.AverageMax(),
		AverageMin:        entries.AverageMin(),
		AverageDifference: entries.AverageDifference(),
	}, nil
}

func (h *handler) CreateWeight(ctx context.Context, in *pb.WeightParams) (*emptypb.Empty, error) {
	date, err := time.Parse("2006-01-02", in.GetDate())
	if err != nil {
		return &emptypb.Empty{}, err
	}

	weight, err := entity.NewWeight(date, in.GetMax(), in.GetMin())
	if err != nil {
		return &emptypb.Empty{}, err
	}

	if err := h.repository.Create(weight); err != nil {
		return &emptypb.Empty{}, err
	}

	return &emptypb.Empty{}, nil
}

func (h *handler) ReadWeight(ctx context.Context, in *pb.WeightParams) (*pb.Weight, error) {
	date, err := time.Parse("2006-01-02", in.GetDate())
	if err != nil {
		return nil, err
	}

	weight, err := h.repository.Read(date)
	if err != nil {
		return nil, err
	}

	return &pb.Weight{
		Date:       timestamppb.New(weight.Date()),
		Max:        weight.Max(),
		Min:        weight.Min(),
		Difference: weight.Difference(),
	}, nil
}

func (h *handler) UpdateWeight(ctx context.Context, in *pb.WeightParams) (*emptypb.Empty, error) {
	date, err := time.Parse("2006-01-02", in.GetDate())
	if err != nil {
		return &emptypb.Empty{}, err
	}

	weight, err := entity.NewWeight(date, in.GetMax(), in.GetMin())
	if err != nil {
		return &emptypb.Empty{}, err
	}

	if err := h.repository.Update(weight); err != nil {
		return &emptypb.Empty{}, err
	}

	return &emptypb.Empty{}, nil
}

func (h *handler) DeleteWeight(ctx context.Context, in *pb.WeightParams) (*emptypb.Empty, error) {
	date, err := time.Parse("2006-01-02", in.GetDate())
	if err != nil {
		return &emptypb.Empty{}, err
	}
	if err := h.repository.Delete(date); err != nil {
		return &emptypb.Empty{}, err
	}
	return &emptypb.Empty{}, nil
}

func New(repository repository.WeightRepository) *handler {
	return &handler{
		repository: repository,
	}
}
