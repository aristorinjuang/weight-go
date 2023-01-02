package data

import (
	"time"

	"github.com/aristorinjuang/weight-go/internal/entity"
)

var (
	Weights entity.Weights = []*entity.Weight{
		entity.RebuildWeight(time.Date(2022, 5, 22, 0, 0, 0, 0, time.UTC), 2, 1),
		entity.RebuildWeight(time.Date(2022, 5, 21, 0, 0, 0, 0, time.UTC), 3, 2),
		entity.RebuildWeight(time.Date(2022, 5, 20, 0, 0, 0, 0, time.UTC), 4, 3),
	}
	AverageMax        float64 = 3
	AverageMin        float64 = 2
	AverageDifference float64 = 1
)
