package entity_test

import (
	"testing"
	"time"

	"github.com/aristorinjuang/weight-go/internal/entity"
	"github.com/aristorinjuang/weight-go/test/data"
	"github.com/stretchr/testify/assert"
)

func TestWeight(t *testing.T) {
	t.Run("it has date", func(t *testing.T) {
		now := time.Now()
		weight, err := entity.NewWeight(now, 2, 1)

		assert.Equal(t, now, weight.Date())
		assert.NoError(t, err)
	})

	t.Run("it has max", func(t *testing.T) {
		now := time.Now()
		weight, err := entity.NewWeight(now, 2, 1)

		assert.Equal(t, float64(2), weight.Max())
		assert.NoError(t, err)
	})

	t.Run("it has min", func(t *testing.T) {
		now := time.Now()
		weight, err := entity.NewWeight(now, 2, 1)

		assert.Equal(t, float64(1), weight.Min())
		assert.NoError(t, err)
	})

	t.Run("it has difference", func(t *testing.T) {
		weight, err := entity.NewWeight(time.Now(), 2, 1)

		assert.Equal(t, float64(1), weight.Difference())
		assert.NoError(t, err)
	})

	t.Run("date cannot be empty", func(t *testing.T) {
		weight, err := entity.NewWeight(time.Time{}, 2, 1)

		assert.Nil(t, weight)
		assert.Error(t, err)
	})

	t.Run("max cannot be empty", func(t *testing.T) {
		weight, err := entity.NewWeight(time.Now(), 0, 1)

		assert.Nil(t, weight)
		assert.Error(t, err)
	})

	t.Run("min cannot be empty", func(t *testing.T) {
		weight, err := entity.NewWeight(time.Now(), 2, 0)

		assert.Nil(t, weight)
		assert.Error(t, err)
	})

	t.Run("max cannot be equal or lower than min", func(t *testing.T) {
		weight, err := entity.NewWeight(time.Now(), 1, 1)

		assert.Nil(t, weight)
		assert.Error(t, err)
	})

	t.Run("rebuild", func(t *testing.T) {
		assert.NotNil(t, entity.RebuildWeight(time.Now(), 2, 1))
	})
}

func TestWeights(t *testing.T) {
	t.Run("map", func(t *testing.T) {
		weights, _, _ := data.Weights.Map()

		assert.Len(t, weights, 3)
	})
}
