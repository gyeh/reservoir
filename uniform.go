package reservoir

import (
	"sync/atomic"
	"unsafe"
)

type UniformR struct {
	res *reservoir
}

func (r *UniformR) Add(value uint64) {
	r.load().Add(value)
}

func (r *UniformR) Count() int32 {
	return r.load().Count()
}

func (r *UniformR) View() []uint64 {
	return r.load().View()
}

func (r *UniformR) load() *reservoir {
	return (*reservoir)(atomic.LoadPointer(unsafe.Pointer(&r.res)))
}

// returns View of reservior and resets
func (r *UniformR) Snapshot() []uint64 {
	n := NewReservoir(r.load().limit)
	v := r.View()
	atomic.StorePointer(unsafe.Pointer(r), &n)
	return v
}
