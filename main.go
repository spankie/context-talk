package main

import (
	"context"
	"fmt"
	"time"
)

func sleepAndTalk(ctx context.Context, t time.Duration, s string) {
	// time.Sleep(t)
	fmt.Println("I am doing a time consuming task")
	select {
	case <-time.After(t):
		fmt.Printf("%s", s)
	case <-ctx.Done():
		fmt.Printf("%s - %s", ctx.Err().Error(), s)
	}
}

// func sleepAndTalk(t time.Duration, s string) {
// time.Sleep(t)
// fmt.Println(s)
// }

func main() {
	// greeting := "Hello Gophers!"

	// 1
	// sleepAndTalk(5*time.Second, greeting)

	// 2
	// ctx := context.Background()
	// sleepAndTalk(ctx, 4*time.Second, greeting)

	// 3
	// context with cancel
	/*
		ctx := context.Background()
		ctx, cancel := context.WithCancel(ctx)
		go func() {
			fmt.Printf("Type in something to stop the process: ")
			// as soon as we read something, cancel
			s := bufio.NewScanner(os.Stdin)
			s.Scan()
			cancel()
			// or sleep for a second and cancel the function.
			// time.Sleep(time.Second)
			// cancel()
		}()
		sleepAndTalk(ctx, 7*time.Second, greeting)
	*/

	/*
		// context with timeout
		ctx := context.Background()
		ctx, cancel := context.WithTimeout(ctx, time.Second)
		// Its very important to call cancel regardless of the timeout, because
		// resources are allocated for the timers and they are only freed when cancel is called.
		// this is done by calling the remove child method inside the cancel function
		// it is documented in WithDeadline.
		defer cancel()
		sleepAndTalk(ctx, 5*time.Second, greeting)
	*/

	/*
		// propagate contexts
		ctx := context.Background()
		ctxA, cancelA := context.WithDeadline(ctx, time.Now().Add(5*time.Second))
		ctxB, cancelB := context.WithTimeout(ctxA, 5*time.Second)

		defer func() { cancelB() }()
		go func() { time.AfterFunc(2*time.Second, cancelA) }()
		sleepAndTalk(ctxB, 11*time.Second, greeting)
	*/
}
