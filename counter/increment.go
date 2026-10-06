package counter_

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type Counter struct {
	mu   sync.Mutex
	Size int
}

func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.Size++
}

type CounterAtomic struct {
	Size atomic.Int64
}

func (c *CounterAtomic) Inc() {
	c.Size.Add(1)
}

func (c *CounterAtomic) Value() int64 {
	return c.Size.Load()
}
func main() {
	cnt := Counter{}

	cnt.Inc()
	cnt.Inc()
	cnt.Inc()

	fmt.Println("Counter size", cnt.Size)
}
