package aggregate

import (
	"errors"

	"github.com/aristorinjuang/weight-go/internal/domain/entity"
)

type Weights struct {
	weights entity.WeightMap
	maxs    float64
	mins    float64
}

func (w *Weights) Weights() entity.WeightMap {
	return w.weights
}

func (w *Weights) AverageMax() float64 {
	return w.maxs / float64(len(w.weights))
}

func (w *Weights) AverageMin() float64 {
	return w.mins / float64(len(w.weights))
}

func (w *Weights) AverageDifference() float64 {
	return w.AverageMax() - w.AverageMin()
}

func (w *Weights) AddWeight(weight *entity.Weight) error {
	if weight == nil {
		return errors.New("weight cannot be empty")
	}

	if _, exists := w.weights[weight.Date()]; !exists {
		w.weights[weight.Date()] = weight
		w.maxs += weight.Max()
		w.mins += weight.Min()
	}

	return nil
}

func NewWeights(weights entity.Weights) *Weights {
	if weights == nil {
		return &Weights{
			weights: make(entity.WeightMap),
		}
	}

	weightsMap, maxs, mins := weights.Map()

	return &Weights{
		weights: weightsMap,
		maxs:    maxs,
		mins:    mins,
	}
}
