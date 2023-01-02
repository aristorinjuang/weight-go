package entity

import (
	"errors"
	"time"
)

type Weight struct {
	date time.Time
	max  float64
	min  float64
}

func (w *Weight) Date() time.Time {
	return w.date
}

func (w *Weight) Max() float64 {
	return w.max
}

func (w *Weight) Min() float64 {
	return w.min
}

func (w *Weight) Difference() float64 {
	return w.max - w.min
}

func NewWeight(date time.Time, max, min float64) (*Weight, error) {
	if date == (time.Time{}) {
		return nil, errors.New("date cannot be empty")
	}
	if max <= 0 {
		return nil, errors.New("max cannot be empty")
	}
	if min <= 0 {
		return nil, errors.New("min cannot be empty")
	}
	if max <= min {
		return nil, errors.New("max cannot be equal or lower than min")
	}
	return &Weight{
		date: date,
		max:  max,
		min:  min,
	}, nil
}

func RebuildWeight(date time.Time, max, min float64) *Weight {
	return &Weight{
		date: date,
		max:  max,
		min:  min,
	}
}

type (
	Weights   []*Weight
	WeightMap map[time.Time]*Weight
)

func (w Weights) Map() (m WeightMap, maxs, mins float64) {
	m = make(WeightMap)

	for _, weight := range w {
		m[weight.date] = weight
		maxs += weight.Max()
		mins += weight.Min()
	}

	return
}
