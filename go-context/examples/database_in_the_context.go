package main

import (
	"log"
	"net/http"

	"context"

	"github.com/dkondratovych/golang-ua-meetup/go-context/examples/database"
)

func main() {
	var err error
	DB, err = database.NewDatabase(database.Config{
		IP:       "",
		User:     "",
		Password: "",
		Name:     "",
	})

	if err != nil {
		log.Fatal(err.Error())
	}

	http.HandleFunc("/test", setDatabaseMiddleware(handleAndQuery))

	log.Fatal(http.ListenAndServe(":8181", nil))
}

// START1 OMIT
var DB database.Database

func setDatabaseMiddleware(h http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		dctx := database.NewContext(r.Context(), DB) // HL

		r = r.WithContext(dctx) // HL

		h(w, r)
	}
}

// STOP1 OMIT

// START2 OMIT
func handleAndQuery(w http.ResponseWriter, r *http.Request) {
	// You can retrieve db from context
	// if you need to perform single db request
	db := database.MustFromContext(r.Context()) // HL
	_ = db                                      // OMIT

	// You can create new ctx with tx inside and pass it into functions
	txctx, tx := database.NewTransactionContext(r.Context()) // HL

	if err := foo(txctx); err != nil {
		tx.Rollback() // HL
		return
	}
	tx.Commit() // HL

	return
}

func foo(ctx context.Context) error {
	db := database.MustFromContext(ctx)
	_ = db //OMIT
	// Perform db operations
	return nil
}

// STOP2 OMIT
