package slidingWindow

import (
	"sync"
	"time"
)

type DefaultMetricsCollector struct {
	mu *sync.RWMutex

	numRequest Number
	success    Number
	failures   Number
}

type Number struct {
	Buckets map[int64]*numberBucket
	mu *sync.RWMutex
}

type numberBucket struct{
	Value float64
}

func (r *Number) GetCurrentBucket() *numberBucket {
	now := time.Now().Unix()
	var bucket *numberBucket
	var ok bool

	if bucket, ok = r.Buckets[now]; !ok {
		bucket = &numberBucket{}
		r.Buckets[now] = bucket
	}
	return bucket
}

func (r *Number) removeOldBuckets() {
	now := time.Now().Unix() - 10

	for timestamp := range r.Buckets {
		if timestamp <= now {
			delete(r.Buckets, timestamp)
		}
	}
}

func (r *Number) Increase(n float64) {
	if n == 0 {
		return
	}

	r.mu.Lock()
	defer r.mu.Unlock()

	b := r.GetCurrentBucket()
	b.Value += n
	r.removeOldBuckets()
}