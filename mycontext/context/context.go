package contextimpl

import (
	"errors"
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
	parent Context
	done   chan struct{}
	err    error
}

func (ctx *cancelCtx) Deadline() (deadline time.Time, ok bool) {
	return ctx.parent.Deadline()
}
func (ctx *cancelCtx) Done() <-chan struct{} {
	return ctx.done
}
func (ctx *cancelCtx) Err() error {
	return ctx.err
}
func (ctx *cancelCtx) Value(key interface{}) interface{} {
	return ctx.parent.Value(key)
}

var Canceled = errors.New("context canceled")

type CancelFunc func()

func WithCancel(parent Context) (Context, CancelFunc) {
	ctx := &cancelCtx{
		parent: parent,
		done:   make(chan struct{}),
	}

	cancel := func() {
		ctx.err = Canceled
		close(ctx.done)
	}

	return ctx, cancel
}
