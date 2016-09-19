package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/test", handler)

	log.Fatal(http.ListenAndServe(":8181", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {

	go func() {
		for i := 0; ; i++ {
			time.Sleep(1 * time.Second)
			fmt.Printf("Worker %d \n", i)
		}
	}()

	io.WriteString(w, "bazinga!")
}
