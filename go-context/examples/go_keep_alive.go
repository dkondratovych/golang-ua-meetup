package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/test", handlerRequest)

	log.Fatal(http.ListenAndServe(":8181", nil))
}

// START1 OMIT
// handlerRequest on each incoming request is handled in its own goroutine // HL
func handlerRequest(w http.ResponseWriter, r *http.Request) {
	longRunningCalculation := func() {
		for i := 0; ; i++ {
			time.Sleep(1 * time.Second)
			fmt.Printf("Worker %d \n", i)
		}
	}

	// goroutine keep working when we return response to a client // HL
	go longRunningCalculation()

	io.WriteString(w, "bazinga!")
	return
}

// STOP1 OMIT
