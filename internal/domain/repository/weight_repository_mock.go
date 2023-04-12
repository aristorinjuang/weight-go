package repository

import (
	"time"

	"github.com/aristorinjuang/weight-go/internal/domain/aggregate"
	"github.com/aristorinjuang/weight-go/internal/domain/entity"
	"github.com/stretchr/testify/mock"
)

type WeightRepositoryMock struct {
	mock.Mock
}

func (m *WeightRepositoryMock) List() (*aggregate.Weights, error) {
	args := m.Called()

	return args.Get(0).(*aggregate.Weights), args.Error(1)
}

func (m *WeightRepositoryMock) Create(weight *entity.Weight) error {
	args := m.Called(weight)

	return args.Error(0)
}

func (m *WeightRepositoryMock) Read(date time.Time) (*entity.Weight, error) {
	args := m.Called(date)

	return args.Get(0).(*entity.Weight), args.Error(1)
}

func (m *WeightRepositoryMock) Update(weight *entity.Weight) error {
	args := m.Called(weight)

	return args.Error(0)
}

func (m *WeightRepositoryMock) Delete(date time.Time) error {
	args := m.Called(date)

	return args.Error(0)
}
