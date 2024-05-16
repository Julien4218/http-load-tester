package process

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAddPool(t *testing.T) {
	pool := NewJobPool()
	pool.CreateProcessor()
	require.Equal(t, 1, pool.Size())
}

func TestRemovePool(t *testing.T) {
	pool := NewJobPool()
	pool.CreateProcessor()
	pool.CreateProcessor()
	pool.CreateProcessor()
	pool.RemoveProcessor()
	require.Equal(t, 2, pool.Size())
}

func TestAdjustAdd(t *testing.T) {
	pool := NewJobPool()
	pool.CreateProcessor()

	spec := NewBatchSpec(2)

	pool.AdjustSize(spec, 1, 4)
	require.Equal(t, 2, pool.Size())
}

func TestAdjustRemove(t *testing.T) {
	pool := NewJobPool()
	pool.CreateProcessor()
	pool.CreateProcessor()
	pool.CreateProcessor()
	pool.CreateProcessor()

	spec := NewBatchSpec(2)

	pool.AdjustSize(spec, 1, 4)
	require.Equal(t, 2, pool.Size())
}

func TestAdjustShouldNotAddMore(t *testing.T) {
	pool := NewJobPool()
	pool.CreateProcessor()
	pool.CreateProcessor()
	pool.CreateProcessor()

	spec := NewBatchSpec(6)

	pool.AdjustSize(spec, 1, 4)
	require.Equal(t, 4, pool.Size())
}

func TestAdjustShouldNotRemoveMore(t *testing.T) {
	pool := NewJobPool()
	pool.CreateProcessor()
	pool.CreateProcessor()
	pool.CreateProcessor()
	pool.CreateProcessor()

	spec := NewBatchSpec(1)

	pool.AdjustSize(spec, 2, 4)
	require.Equal(t, 2, pool.Size())
}
