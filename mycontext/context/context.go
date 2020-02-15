package contextimpl

import (
	"errors"
	"sync"
	"time"
)

type Context interface {
	Deadline() (deadline time.Time, ok bool)
	Done() <-chan struct{}
	Err() error
	Value(key interface{}) interface{}
}

type emptyCtx int

func (e emptyCtx) Deadline() (deadline time.Time, ok bool) {
	return
}
func (e emptyCtx) Done() <-chan struct{} {
	return nil
}
func (e emptyCtx) Err() error {
	return nil
}
func (e emptyCtx) Value(key interface{}) interface{} {
	return nil
}

var (
	background = new(emptyCtx)
	todo       = new(emptyCtx)
)

func Background() Context {
	return background
	// returning empty context will be the same for both background and todo
	// return emptyCtx{}
}

func TODO() Context {
	return todo
	// returning empty context will be the same for both background and todo
	// return emptyCtx{}
}

type cancelCtx struct {
	Context
	done chan struct{}
	err  error
	mu   sync.Mutex
}

func (ctx *cancelCtx) Done() <-chan struct{} {
	return ctx.done
}
func (ctx *cancelCtx) Err() error {
	ctx.mu.Lock()
	defer ctx.mu.Unlock()
	return ctx.err
}

var Canceled = errors.New("context canceled")

type CancelFunc func()

func WithCancel(parent Context) (Context, CancelFunc) {
	ctx := &cancelCtx{
		Context: parent,
		done:    make(chan struct{}),
	}

	cancel := func() {
		ctx.cancel(Canceled)
	}

	go func() {
		select {
		case <-parent.Done():
			ctx.cancel(parent.Err())
		case <-ctx.Done():
		}
	}()

	return ctx, cancel
}

func (ctx *cancelCtx) cancel(err error) {
	ctx.mu.Lock()
	defer ctx.mu.Unlock()
	// stop here if err != nil to prevent closing
	// a closed channel. this can occur if the cancel
	// function is called twice or the parent cancel is
	// called and the current cancel is also called
	if ctx.err != nil {
		return
	}
	ctx.err = err
	close(ctx.done)
}

type deadlineCtx struct {
	*cancelCtx
	deadline time.Time
}

func (ctx *deadlineCtx) Deadline() (deadline time.Time, ok bool) {
	return ctx.deadline, true
}

var DeadlineExceeded = errors.New("deadline exceeded")

func WithDeadline(parent Context, deadline time.Time) (Context, CancelFunc) {
	cctx, cancel := WithCancel(parent)

	ctx := &deadlineCtx{
		cancelCtx: cctx.(*cancelCtx),
		deadline:  deadline,
	}

	timer := time.AfterFunc(time.Until(deadline), func() {
		ctx.cancel(DeadlineExceeded)
	})

	stop := func() {
		timer.Stop()
		cancel()
	}

	return ctx, stop
}

func WithTimeout(parent Context, timeout time.Duration) (Context, CancelFunc) {
	return WithDeadline(parent, time.Now().Add(timeout))
}

func WithValue(parent Context, key interface{}, value interface{}) {
	return &valueCtx{
		Context: parent,
	}
}
