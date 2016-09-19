package main

import (
	"context"
	"io"
	"log"
	"net/http"
	"time"
)

func main() {

	http.HandleFunc("/test", handlerSearchCancel)

	log.Fatal(http.ListenAndServe(":8181", nil))
}

// START1 OMIT
func handlerSearchCancel(w http.ResponseWriter, r *http.Request) {
	var ctx context.Context
	var cancel context.CancelFunc

	ctx, cancel = context.WithCancel(r.Context()) // HL
	defer cancel()                                // HL

	// Close context.Done channel in 4 seconds
	go func() {
		time.Sleep(4 * time.Second)
		cancel() // HL
	}()

	select {
	case <-ctx.Done(): // HL
		log.Print(ctx.Err())
		return
	case result := <-longRunningCalculation(): // HL
		io.WriteString(w, result)
	}

	return
}

// STOP1 OMIT

func longRunningCalculation() <-chan string {
	r := make(chan string)

	go func() {
		time.Sleep(10 * time.Second)
		r <- "I am done"
	}()

	return r
}
