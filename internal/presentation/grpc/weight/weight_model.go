package weight

import (
	"github.com/aristorinjuang/weight-go/internal/domain/entity"
	pb "github.com/aristorinjuang/weight-go/tools/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type weights []*pb.Weight

func (w weights) From(weights entity.WeightMap) weights {
	for _, entry := range weights {
		w = append(w, &pb.Weight{
			Date:       timestamppb.New(entry.Date()),
			Max:        entry.Max(),
			Min:        entry.Min(),
			Difference: entry.Difference(),
		})
	}

	return w
}
