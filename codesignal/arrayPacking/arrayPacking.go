package arraypacking

import (
	"sync"
	"sync/atomic"
)

func arrayPacking(a []int) int {
	r := 0
	for i, n := range a {
		r += n << (i * 8)
	}
	return r
}

func arrayPackingV2(a []int) int {
	var wg sync.WaitGroup
	var sum uint32
	for i, n := range a {
		wg.Add(1)
		go func(i int, n int) {
			atomic.AddUint32(&sum, uint32(n<<(i*8)))
			wg.Done()
		}(i, n)
	}
	wg.Wait()
	return int(sum)
}
