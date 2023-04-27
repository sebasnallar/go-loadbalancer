package loadbalancer

import (
	"sync/atomic"
)

type RoundRobin struct {
	index int64
	size  int
}

func NewRoundRobin(size int) *RoundRobin {
	return &RoundRobin{
		index: -1,
		size:  size,
	}
}

func (r *RoundRobin) Next() int {
	newIndex := atomic.AddInt64(&r.index, 1)
	return int(newIndex % int64(r.size))
}
