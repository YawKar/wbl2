package orchan

import (
	"testing"
	"time"
)

func TestSample(t *testing.T) {
	sig := func(after time.Duration) <-chan struct{} {
		c := make(chan struct{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}
	start := time.Now()
	<-Or(
		sig(2*time.Second),
		sig(1*time.Second),
		sig(4*time.Second),
		sig(300*time.Millisecond),
	)
	delta := time.Since(start) - 300*time.Millisecond
	if delta >= 10*time.Millisecond {
		t.Fatalf("delta is >= 10ms: %d", delta)
	}
}
