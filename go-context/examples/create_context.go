package main

import (
	"context"
	"io"
	"log"
	"net/http"
	// "time"
)

func main() {
	http.HandleFunc("/test", middleware(handleRequest))

	log.Fatal(http.ListenAndServe(":8181", nil))
}

func middleware(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Create context from context package
		// and attach it to request later with r.WithContext()
		ctx := context.Background() // HL
		// OR

		// Use request method context, which returns context, if ctx is nil
		// then creates new background context - context.Background()
		ctx = r.Context() // HL

		// Build context variations on top of background context
		ctx = context.WithValue(ctx, "some_key", "some_value")

		// tctx, cancelFunc := context.WithTimeout(ctx, time.Duration(5 * time.Second))
		// deadline := time.Now().Add(time.Duration(30 * time.Second))
		// dctx, cancelFunc := context.WithDeadline(ctx, deadline)

		// WithContext returns a shallow copy of r with its context changed to ctx.
		r = r.WithContext(ctx) // HL

		h(w, r)
	}
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context() // HL

	if err := doSomething(ctx, "payload"); err != nil { // HL
		log.Print(err)
	}

	io.WriteString(w, "Bazinga!")
}

// STOP OMIT

func doSomething(ctx context.Context, payload string) error {

	_ = payload // OMIT

	return nil
}
