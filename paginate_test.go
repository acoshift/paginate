package paginate_test

import (
	"testing"

	"github.com/acoshift/paginate"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	p := paginate.New(0, -1, 100)
	assert.Equal(t, int64(1), p.Page())
	assert.Equal(t, int64(1), p.PerPage())
	assert.Equal(t, int64(0), p.Offset())
	assert.Equal(t, int64(1), p.Limit())
	assert.Equal(t, int64(100), p.Items())
	assert.Equal(t, int64(100), p.Count())
	assert.Equal(t, int64(100), p.MaxPage())
	assert.False(t, p.CanPrev())
	assert.True(t, p.CanNext())
	{
		limit, offset := p.LimitOffset()
		assert.Equal(t, p.Limit(), limit)
		assert.Equal(t, p.Offset(), offset)
	}

	p = paginate.New(1, 10, 100)
	assert.Equal(t, int64(1), p.Page())
	assert.Equal(t, int64(10), p.PerPage())
	assert.Equal(t, int64(0), p.Offset())
	assert.Equal(t, int64(10), p.Limit())
	assert.Equal(t, int64(100), p.Items())
	assert.Equal(t, int64(100), p.Count())
	assert.Equal(t, int64(10), p.MaxPage())
	assert.False(t, p.CanPrev())
	assert.True(t, p.CanNext())

	p = paginate.New(5, 10, 100)
	assert.Equal(t, int64(5), p.Page())
	assert.Equal(t, int64(10), p.PerPage())
	assert.Equal(t, int64(40), p.Offset())
	assert.Equal(t, int64(10), p.Limit())
	assert.Equal(t, int64(100), p.Items())
	assert.Equal(t, int64(100), p.Count())
	assert.Equal(t, int64(10), p.MaxPage())
	assert.True(t, p.CanPrev())
	assert.True(t, p.CanNext())

	p = paginate.New(10, 10, 100)
	assert.Equal(t, int64(10), p.Page())
	assert.Equal(t, int64(10), p.PerPage())
	assert.Equal(t, int64(90), p.Offset())
	assert.Equal(t, int64(10), p.Limit())
	assert.Equal(t, int64(100), p.Items())
	assert.Equal(t, int64(100), p.Count())
	assert.Equal(t, int64(10), p.MaxPage())
	assert.True(t, p.CanPrev())
	assert.False(t, p.CanNext())

	p = paginate.New(1, 10, -1)
	assert.Equal(t, int64(1), p.Page())
	assert.Equal(t, int64(10), p.PerPage())
	assert.Equal(t, int64(0), p.Items())
	assert.Equal(t, int64(1), p.MaxPage())

	assert.Equal(t, int64(4), paginate.New(1, 5, 16).MaxPage())
	assert.Equal(t, int64(4), paginate.New(100, 5, 16).Page())
	assert.Equal(t, int64(1), paginate.New(1, 5, 16).Prev())
	assert.Equal(t, int64(1), paginate.New(2, 5, 16).Prev())
	assert.Equal(t, int64(2), paginate.New(1, 5, 16).Next())
	assert.Equal(t, int64(4), paginate.New(4, 5, 16).Next())

	p = paginate.FromLimitOffset(10, 0, 100)
	assert.Equal(t, int64(1), p.Page())
	assert.Equal(t, int64(10), p.PerPage())
	assert.Equal(t, int64(0), p.Offset())
	assert.Equal(t, int64(10), p.Limit())
	assert.Equal(t, int64(100), p.Items())
	assert.Equal(t, int64(100), p.Count())
	assert.Equal(t, int64(10), p.MaxPage())
	assert.False(t, p.CanPrev())
	assert.True(t, p.CanNext())

	p = paginate.FromLimitOffset(10, 40, 100)
	assert.Equal(t, int64(5), p.Page())
	assert.Equal(t, int64(10), p.PerPage())
	assert.Equal(t, int64(40), p.Offset())
	assert.Equal(t, int64(10), p.Limit())
	assert.Equal(t, int64(100), p.Items())
	assert.Equal(t, int64(100), p.Count())
	assert.Equal(t, int64(10), p.MaxPage())
	assert.True(t, p.CanPrev())
	assert.True(t, p.CanNext())

	p = paginate.FromLimitOffset(1, 10, -1)
	assert.Equal(t, int64(1), p.Page())
	assert.Equal(t, int64(1), p.PerPage())
	assert.Equal(t, int64(0), p.Items())
	assert.Equal(t, int64(1), p.MaxPage())

	p = paginate.FromLimitOffset(-10, -10, -10)
	assert.Equal(t, int64(1), p.Page())
	assert.Equal(t, int64(1), p.PerPage())
	assert.Equal(t, int64(0), p.Items())
	assert.Equal(t, int64(1), p.MaxPage())
}

func TestPages(t *testing.T) {
	assert.Equal(t, []int64{1, 2, 0, 8, 9, 10, 11, 12, 0, 19, 20}, paginate.New(10, 1, 20).Pages(2, 2))
	assert.Equal(t, []int64{0, 10, 0}, paginate.New(10, 1, 20).Pages(0, 0))
	assert.Equal(t, []int64{1}, paginate.New(1, 1, 1).Pages(2, 2))
	assert.Equal(t, []int64{1, 2, 3, 4, 5, 6, 0, 9, 10}, paginate.New(4, 1, 10).Pages(2, 2))
	assert.Equal(t, []int64{1, 2, 0, 6, 7, 8, 9, 10, 11, 12, 13}, paginate.New(8, 1, 13).Pages(2, 2))
	assert.Equal(t, []int64{1, 2, 3, 4, 5, 6, 7, 8, 0, 12, 13}, paginate.New(6, 1, 13).Pages(2, 2))
	assert.Equal(t, []int64{1, 2, 3, 4, 5, 6}, paginate.New(4, 1, 6).Pages(2, 2))
	assert.Equal(t, []int64{1, 2, 0, 6}, paginate.New(1, 1, 6).Pages(1, 1))
	assert.Equal(t, []int64{1, 2, 0, 5, 6}, paginate.New(1, 1, 6).Pages(1, 2))
}

func TestNewMovable(t *testing.T) {
	p := paginate.NewMovable(0, -1, 100)
	assert.Equal(t, int64(1), p.Page())
	assert.Equal(t, int64(1), p.PerPage())
	assert.Equal(t, int64(0), p.CountOffset())
	assert.Equal(t, int64(101), p.CountLimit())
	assert.Equal(t, int64(0), p.Offset())
	assert.Equal(t, int64(1), p.Limit())
	assert.Equal(t, int64(100), p.Count())
	assert.Equal(t, int64(100), p.MaxPage())
	assert.False(t, p.CanPrev())
	assert.True(t, p.CanNext())
	{
		limit, offset := p.LimitOffset()
		assert.Equal(t, p.Limit(), limit)
		assert.Equal(t, p.Offset(), offset)
	}
	p.Counting(func(limit, offset int64) int64 {
		assert.Equal(t, p.CountLimit(), limit)
		assert.Equal(t, p.CountOffset(), offset)
		return 10
	})
	assert.Equal(t, int64(10), p.Count())

	p = paginate.NewMovable(1, 10, 100)
	assert.Equal(t, int64(1), p.Page())
	assert.Equal(t, int64(10), p.PerPage())
	assert.Equal(t, int64(0), p.CountOffset())
	assert.Equal(t, int64(1001), p.CountLimit())
	assert.Equal(t, int64(0), p.Offset())
	assert.Equal(t, int64(10), p.Limit())
	assert.Equal(t, int64(1000), p.Count())
	assert.Equal(t, int64(100), p.MaxPage())
	assert.False(t, p.CanPrev())
	assert.True(t, p.CanNext())

	p = paginate.NewMovable(1, 10, -1)
	assert.Equal(t, int64(1), p.Page())
	assert.Equal(t, int64(10), p.PerPage())
	assert.Equal(t, int64(10), p.Count())
	assert.Equal(t, int64(1), p.MaxPage())

	assert.Equal(t, int64(10), paginate.NewMovable(1, 5, 10).MaxPage())
	assert.Equal(t, int64(11), paginate.NewMovable(1, 5, 10).SetCount(52).MaxPage())
	assert.Equal(t, int64(100), paginate.NewMovable(100, 5, 16).Page())
	assert.Equal(t, int64(1), paginate.NewMovable(1, 5, 10).Prev())
	assert.Equal(t, int64(1), paginate.NewMovable(2, 5, 10).Prev())
	assert.Equal(t, int64(2), paginate.NewMovable(1, 5, 10).Next())
	assert.Equal(t, int64(4), paginate.NewMovable(4, 5, 0).SetCount(4).Next())
}

func TestMovablePages(t *testing.T) {
	assert.Equal(t, []int64{1, 2, 3, 4, 5, 6, 7, 8}, paginate.NewMovable(1, 1, 8).Pages())
	assert.Equal(t, []int64{6, 7, 8, 9, 10, 11, 12, 13}, paginate.NewMovable(10, 1, 8).Pages())
	assert.Equal(t, []int64{1, 2, 3, 4, 5}, paginate.NewMovable(1, 1, 5).Pages())
	assert.Equal(t, []int64{1, 2, 3, 4, 5}, paginate.NewMovable(2, 1, 5).Pages())
	assert.Equal(t, []int64{1, 2, 3, 4, 5}, paginate.NewMovable(3, 1, 5).Pages())
	assert.Equal(t, []int64{2, 3, 4, 5, 6}, paginate.NewMovable(4, 1, 5).Pages())
	assert.Equal(t, []int64{3, 4, 5, 6, 7}, paginate.NewMovable(5, 1, 5).Pages())
	assert.Equal(t, []int64{4, 5, 6, 7, 8}, paginate.NewMovable(6, 1, 5).Pages())
	assert.Equal(t, []int64{1, 2, 3, 4}, paginate.NewMovable(1, 1, 4).Pages())
	assert.Equal(t, []int64{1, 2, 3, 4}, paginate.NewMovable(2, 1, 4).Pages())
	assert.Equal(t, []int64{1, 2, 3, 4}, paginate.NewMovable(3, 1, 4).Pages())
	assert.Equal(t, []int64{2, 3, 4, 5}, paginate.NewMovable(4, 1, 4).Pages())
	assert.Equal(t, []int64{3, 4, 5, 6}, paginate.NewMovable(5, 1, 4).Pages())
	assert.Equal(t, []int64{4, 5, 6, 7}, paginate.NewMovable(6, 1, 4).Pages())
}

func TestEmpty(t *testing.T) {
	assert.Equal(t, []int64{1}, paginate.New(0, 0, 0).Pages(2, 2))
	assert.Equal(t, []int64{1}, (&paginate.Paginate{}).Pages(2, 2))
}
