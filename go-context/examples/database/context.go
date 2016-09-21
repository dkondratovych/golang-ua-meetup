package database

import "context"

type key int

const databaseKey key = 0

// FromContext retrieves database from context
func FromContext(ctx context.Context) (Database, bool) {
	sp, ok := ctx.Value(databaseKey).(Database)
	return sp, ok
}

// START1 OMIT

// NewContext adds database to context
func NewContext(ctx context.Context, database Database) context.Context {
	return context.WithValue(ctx, databaseKey, database)
}

// MustFromContext retrieves database from context and Panics if not found
func MustFromContext(ctx context.Context) Database {
	sp, ok := ctx.Value(databaseKey).(Database)
	if !ok {
		panic("database was not found in context")
	}
	return sp
}

// NewTransactionContext gets database connections from
// existing context, starts transaction and puts it back to context
func NewTransactionContext(ctx context.Context) (context.Context, Database) { // HL
	tx := MustFromContext(ctx).MustBeginTransaction() // HL

	return NewContext(ctx, tx), tx // HL
} // HL

// STOP1 OMIT
