package main

import (
	"context"
	"fmt"
	"time"
)

func sleepAndTalk(ctx context.Context, t time.Duration, s string) {
	// time.Sleep(t)
	select {
	case <-time.After(t):
		fmt.Printf("%s", s)
	case <-ctx.Done():
		fmt.Printf("%s - %s", ctx.Err().Error(), s)
	}
}

func main() {
	fmt.Println("Hello gophers, lets talk about context")
	/*
		ctx := context.Background()
		sleepAndTalk(ctx, 4*time.Second, "Hello gophers")
	*/

	/*
		// context with cancel
		ctx := context.Background()
		ctx, cancel := context.WithCancel(ctx)
		go func() {
			// as soon as we read something, cancel
			s := bufio.NewScanner(os.Stdin)
			s.Scan()
			cancel()
			// or sleep for a second and cancel the function.
			// time.Sleep(time.Second)
			// cancel()
		}()
		sleepAndTalk(ctx, 7*time.Second, "Hello gophers")
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
		sleepAndTalk(ctx, 5*time.Second, "Hello Eko Gophers")
	*/

	/*
		// propagate contexts
		ctx := context.Background()
		ctx, cancel1 := context.WithDeadline(ctx, time.Now().Add(5*time.Second))
		ctx, cancel0 := context.WithTimeout(ctx, 5*time.Second)

		// Its very important to call cancel regardless of the timeout, because
		// resources are allocated for the timers and they are only freed when cancel is called.
		// this is done by calling the remove child method inside the cancel function
		// it is documented in WithDeadline.
		defer func() { cancel0(); cancel1() }()
		sleepAndTalk(ctx, 11*time.Second, "Hello Eko Gophers")
	*/
	// i stopped at 12:53
}
