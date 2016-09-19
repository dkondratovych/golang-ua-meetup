package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/test", handlerRequestWithCancelation)

	log.Fatal(http.ListenAndServe(":8181", nil))
}

// START1 OMIT
// handlerRequest on each incoming request is handled in its own goroutine // HL
func handlerRequestWithCancelation(w http.ResponseWriter, r *http.Request) {
	longRunningCalculation := func(ctx context.Context) {
		for i := 0; ; i++ {
			select {
			case <-ctx.Done(): // HL
				return
			default:
				time.Sleep(1 * time.Second)
				fmt.Printf("Worker %d \n", i)
			}
		}
	}

	// the context is canceled when the ServeHTTP method returns  // HL
	go longRunningCalculation(r.Context()) // HL

	// give some time for longRunningCalculation to do some work
	time.Sleep(5 * time.Second)

	io.WriteString(w, "bazinga!")
	return
}

// STOP1 OMIT
