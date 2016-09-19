package main

import (
	"errors"
	"io"
	"log"
	"net/http"

	"github.com/dkondratovych/golang-ua-meetup/go-context/examples/user"
)

var (
	ErrUserNotFound = errors.New("User not found")
)

func main() {
	http.HandleFunc("/test", setUserMiddleware(handleUserRequest))

	log.Fatal(http.ListenAndServe(":8181", nil))
}

func setUserMiddleware(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := user.NewUserContext(r.Context(), &user.User{
			Name: "Gopher",
			Age:  20,
		})

		// WithContext returns a shallow copy of r with its context changed to ctx.
		r = r.WithContext(ctx)

		h(w, r)
	}
}

// STOP1 OMIT

func handleUserRequest(w http.ResponseWriter, r *http.Request) {
	// Use EntityNameFromContext if value is not critical for function execution
	//u, ok := user.UserFromContext(r.Context())
	//if !ok {
	//	log.Print(ErrUserNotFound)
	//}

	// Use EntityNameMustFromContext if value from context
	// is critical for execution (logger, db =))
	u := user.UserMustFromContext(r.Context())

	io.WriteString(w, u.Name)
}

// STOP OMIT
