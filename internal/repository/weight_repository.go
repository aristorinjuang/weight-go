package repository

import (
	"time"

	"github.com/aristorinjuang/weight-go/internal/aggregate"
	"github.com/aristorinjuang/weight-go/internal/entity"
)

type WeightRepository interface {
	List() (*aggregate.Weights, error)
	Create(*entity.Weight) error
	Read(date time.Time) (*entity.Weight, error)
	Update(*entity.Weight) error
	Delete(date time.Time) error
}
