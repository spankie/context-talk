package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	// we can actually get a context from the request
	ctx := r.Context()
	log.Printf("handler started\n")
	defer log.Printf("handler ended\n")

	select {
	case <-ctx.Done():
		err := ctx.Err()
		log.Println(err)
		http.Error(w, ctx.Err().Error(), http.StatusInternalServerError)
	case <-time.After(5 * time.Second):
		fmt.Fprintln(w, "Hello gopher")
	}
}
