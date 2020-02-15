package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", nil))
}

func DoSomethingForLong(ctx context.Context, done chan struct{}) {
	t, ok := ctx.Value(1).(time.Duration)
	if ok {
		time.Sleep(t)
	}
	close(done)
}

func handler(w http.ResponseWriter, r *http.Request) {
	// we can actually get a context from the request
	ctx := r.Context()
	log.Printf("handler started\n")
	defer log.Printf("handler ended\n")

	type key int
	waiting := make(chan struct{}, 1)
	ctx = context.WithValue(ctx, key(1), 5*time.Second)
	go DoSomethingForLong(ctx, waiting)

	select {
	case <-ctx.Done():
		err := ctx.Err()
		log.Println(err)
		http.Error(w, ctx.Err().Error(), http.StatusInternalServerError)
	case <-waiting:
		fmt.Fprintln(w, "Hello gopher")
	}
}
