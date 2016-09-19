package main

import (
	"context"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {

	http.HandleFunc("/test", handlerSearchTimeout)

	http.HandleFunc("/timeout", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(4 * time.Second)
	})

	log.Fatal(http.ListenAndServe(":8181", nil))
}

// START1 OMIT
func handlerSearchTimeout(w http.ResponseWriter, r *http.Request) {
	var ctx context.Context
	var cancel context.CancelFunc

	ctx, cancel = context.WithTimeout(r.Context(), time.Duration(2*time.Second)) // HL
	defer cancel()

	request, err := http.NewRequest(http.MethodGet, "http://localhost:8181/timeout", nil)
	if err != nil {
		log.Println(err.Error())
		return
	}
	// You can context separately for each request
	request = request.WithContext(ctx) // HL

	client := &http.Client{}            // HL
	response, err := client.Do(request) // HL
	// You will get an error "net/http: request canceled" when request timeout exceeds limits // HL
	if err != nil {
		log.Println(err.Error())
		return
	}
	// ......
	// STOP2 OMIT

	b, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println(err.Error())
		return
	}

	log.Printf("%s", b)

	return
}

// STOP1 OMIT
