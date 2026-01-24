package generic

func NewUnboundedChannel[T any]() *UnboundedChannel[T] {
	c := &UnboundedChannel[T]{
		in:  make(chan T),
		out: make(chan T),
	}
	go c.dispatch()
	return c
}

type UnboundedChannel[T any] struct {
	in, out chan T
	queue   FreeList[T]
}

func (c *UnboundedChannel[T]) In() chan<- T {
	return c.in
}

func (c *UnboundedChannel[T]) Out() <-chan T {
	return c.out
}

func (c *UnboundedChannel[T]) Close() {
	close(c.in)
}

func (c *UnboundedChannel[T]) dispatch() {
	in := c.in
	var out chan T
	var v T

	for in != nil || out != nil {
		select {
		case v, ok := <-in:
			if ok {
				c.queue.PushBack(v)
			} else {
				in = nil
			}
		case out <- v:
			c.queue.PopFront()
		}

		if c.queue.Len() > 0 {
			v = c.queue.Front().V
			out = c.out
		} else {
			out = nil
		}
	}

	close(c.out)
}
