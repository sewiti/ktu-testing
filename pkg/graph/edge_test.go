package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEdge(t *testing.T) {
	t.Run("should create graph edge with zero weight", func(t *testing.T) {
		v0 := NewVertex(0)
		v1 := NewVertex(1)
		e := NewEdge(v0, v1, 0)

		assert.Equal(t, "0 to 1", e.String())
		assert.Equal(t, v0, e.start)
		assert.Equal(t, v1, e.end)
		assert.Equal(t, float64(0), e.Weight)
	})

	t.Run("should create graph edge with negative weight", func(t *testing.T) {
		v0 := NewVertex(0)
		v1 := NewVertex(1)
		e := NewEdge(v0, v1, -5)

		assert.Equal(t, "0 to 1", e.String())
		assert.Equal(t, v0, e.start)
		assert.Equal(t, v1, e.end)
		assert.Equal(t, float64(-5), e.Weight)
	})

	t.Run("should be possible to do edge reverse", func(t *testing.T) {
		v0 := NewVertex(0)
		v1 := NewVertex(1)
		e := NewEdge(v0, v1, 6)

		assert.Equal(t, "0 to 1", e.String())
		assert.Equal(t, v0, e.start)
		assert.Equal(t, v1, e.end)
		assert.Equal(t, float64(6), e.Weight)

		e.Reverse()

		assert.Equal(t, "1 to 0", e.String())
		assert.Equal(t, v1, e.start)
		assert.Equal(t, v0, e.end)
		assert.Equal(t, float64(6), e.Weight)
	})
}
