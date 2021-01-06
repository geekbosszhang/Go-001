// referenced Hong and russeellLuo sliding window
package slidingWindow

import (
	"time"
	"sync"
)

type Window interface {
	Start() time.Time
	Count() int64
	AddCount(n int64)
	Reset(s time.Time, c int64)
	Sync(now time.Time)
}

type StopFunc func()

type NewWindow func() (Window, StopFunc)

type Limiter struct {
	size time.Duration
	limit int64
	mu sync.Mutex
	curr Window
	prev Window
}

func NewLimiter(size time.Duration, limit int64, newWindow NewWindow) (*Limiter, StopFunc) {
	currWin, currStop := newWindow()

	prevWin, _ := newWindow()
	limiter := &Limiter{
		size,
		limit,
		sync.Mutex{},
		currWin,
		prevWin,
	}
	return limiter, currStop
}

func (limiter *Limiter) AllowN(now time.Time, n int64) bool {
	limiter.mu.Lock()
	defer limiter.mu.Unlock()

	eclapsed := now.Sub(limiter.curr.Start())
	weight := float64(limiter.size - eclapsed) / float64(limiter.size)
	count := int64(weight * float64(limiter.prev.Count())) + limiter.curr.Count()

	if count + n > limiter.limit {
		return false
	}
	limiter.curr.AddCount(n)
	return true
}