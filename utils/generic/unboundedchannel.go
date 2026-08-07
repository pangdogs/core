package generic

import "sync/atomic"

// NewUnboundedChannel 创建由后台协程转发的无界队列频道。
func NewUnboundedChannel[T any]() *UnboundedChannel[T] {
	c := &UnboundedChannel[T]{
		in:  make(chan T),
		out: make(chan T),
	}
	go c.dispatch()
	return c
}

// UnboundedChannel 解耦输入与输出速度，并按接收顺序保存所有待输出值。
//
// 必须通过 NewUnboundedChannel 创建，创建后不可复制。In、Out 与 Len 可并发使用；
// Close 只能调用一次，且不能与输入发送并发，否则会发生向已关闭频道发送的 panic。
type UnboundedChannel[T any] struct {
	_       noCopy
	in, out chan T
	queue   FreeList[T]
	count   atomic.Int64
}

// In 返回生产者输入端。发送只等待后台转发协程接收，不等待最终消费者。
func (c *UnboundedChannel[T]) In() chan<- T {
	return c.in
}

// Out 返回消费者输出端；Close 后会先排空队列再关闭该频道。
func (c *UnboundedChannel[T]) Out() <-chan T {
	return c.out
}

// Close 关闭输入端；已排队值仍会全部输出。重复调用会 panic。
func (c *UnboundedChannel[T]) Close() {
	close(c.in)
}

// Len 返回当前等待输出的值数量快照。
func (c *UnboundedChannel[T]) Len() int {
	return int(c.count.Load())
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
				c.count.Add(1)
			} else {
				in = nil
			}
		case out <- v:
			c.queue.PopFront()
			c.count.Add(-1)
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
