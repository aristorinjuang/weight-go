package aggregate

import (
	"testing"

	"github.com/aristorinjuang/weight-go/test/data"
	"github.com/stretchr/testify/assert"
)

func TestWeights(t *testing.T) {
	t.Run("no initial weights", func(t *testing.T) {
		weights := NewWeights(nil)

		assert.NotNil(t, weights)
		assert.Len(t, weights.Weights(), 0)
	})

	t.Run("with initial weights", func(t *testing.T) {
		weights := NewWeights(data.Weights)

		assert.NotNil(t, weights)
		assert.Len(t, weights.Weights(), len(data.Weights))
	})

	t.Run("success add weight", func(t *testing.T) {
		weights := NewWeights(nil)

		assert.NoError(t, weights.AddWeight(data.Weights[0]))
	})

	t.Run("failed add weight", func(t *testing.T) {
		weights := NewWeights(nil)

		assert.Error(t, weights.AddWeight(nil))
	})

	t.Run("averages", func(t *testing.T) {
		weights := NewWeights(nil)

		for _, weight := range data.Weights {
			assert.NoError(t, weights.AddWeight(weight))
		}

		assert.Len(t, weights.Weights(), len(data.Weights))
		assert.Equal(t, data.AverageMax, weights.AverageMax())
		assert.Equal(t, data.AverageMin, weights.AverageMin())
		assert.Equal(t, data.AverageDifference, weights.AverageDifference())
	})
}
