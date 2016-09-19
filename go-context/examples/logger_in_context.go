package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dkondratovych/golang-ua-meetup/go-context/examples/logger"
	"github.com/seesawlabs/go-flake"
)

var globalLogger logger.IRequestScopedLogger

func main() {

	globalLogger = logger.NewLogger()

	http.HandleFunc("/test", setLoggerMiddleware(handleAndLog))

	log.Fatal(http.ListenAndServe(":8181", nil))
}

// START1 OMIT
func setLoggerMiddleware(h http.HandlerFunc) http.HandlerFunc {
	// generate new 64 bits long unique flake id
	f, err := flake.New()
	if err != nil {
		panic("could not initialize flake")
	}

	return func(w http.ResponseWriter, r *http.Request) {
		// next values in real app will be retrieved from db, cache, http headers etc
		requestScopedLogger := globalLogger.GetRequestScoped( // HL
			fmt.Sprintf("%x", f.NextId()), // request id		// HL
			"web",    // application name						// HL
			int64(0), // current user id						// HL
		) // HL

		lctx := logger.NewContext(r.Context(), requestScopedLogger) // HL
		r = r.WithContext(lctx)                                     // HL

		h(w, r)
	}
}

// STOP1 OMIT

// START2 OMIT
func handleAndLog(w http.ResponseWriter, r *http.Request) {
	l := logger.MustFromContext(r.Context()) // HL

	l.Printf("Bazinga!")

	return
}

// STOP2 OMIT
