package orchan

func Or[T any](channels ...<-chan T) <-chan T {
	c := make(chan T)
	for _, ch := range channels {
		ch := ch
		go func() {
			c <- <-ch
		}()
	}
	out := make(chan T)
	go func() {
		out <- <-c
	}()
	return out
}
