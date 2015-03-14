package reservoir

import (
	"sort"
	"sync/atomic"
)
import "math/rand"

type Reservoir interface {
	Add(value uint64)
	Count() int32
	View() uint64
}

type reservoir struct {
	size   int32 // current reservior size
	limit  int32 // reservior size limit
	values []uint64
}

func NewReservoir(limit int32) {
	return &reservior{
		size:   0,
		limit:  limit,
		values: make([]uint64, limit),
	}
}

func (r *reservoir) Add(value uint64) {
	currSize := atomic.AddInt32(&r.size, 1) - 1
	if currSize < r.limit {
		r.values = append(r.values, value) // TODO: make thread-safe
	} else {
		i := rand.Int31n(currSize)
		if i < r.limit {
			r.values[i] = value
		}
	}
}

func (r *reservoir) Count() int32 {
	s := atomic.LoadInt32(&r.size)
	if s < r.limit {
		return s
	}
	return r.limit
}

func (r *reservoir) View() []uint64 {
	c := make([]uint64, len(r.values))
	copy(c, r.values)
	sort.Sort(Uint64Slice(c))
	return c
}

type Uint64Slice []uint64

func (p Uint64Slice) Len() int           { return len(p) }
func (p Uint64Slice) Less(i, j int) bool { return p[i] < p[j] }
func (p Uint64Slice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
