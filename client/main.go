package main

import (
	"context"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func reqNoContext() {
	res, err := http.Get("http://localhost:8080")
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		log.Fatal(res.StatusCode)
	}
	io.Copy(os.Stdout, res.Body)
}

func reqWithContext() {
	// create a context with timeout of 1 second
	// and defer the cancel method
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	// create a new request and add the context to the new request
	req, err := http.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	if err != nil {
		log.Fatal(err)
	}
	req = req.WithContext(ctx)

	// send the request
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		log.Fatal(res.StatusCode)
	}
	io.Copy(os.Stdout, res.Body)
}

func main() {

	// reqNoContext()
	reqWithContext()
}
