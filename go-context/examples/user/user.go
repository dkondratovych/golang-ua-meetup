package user

import "context"

type User struct {
	Name string
	Age  int64
}

type contextKey string

var userContextKey contextKey = "user"

// NewUserContext adds user to the context OMIT
func NewUserContext(ctx context.Context, user *User) context.Context {
	return context.WithValue(ctx, userContextKey, user)
}

// FromContext retrieves user from context OMIT
func UserFromContext(ctx context.Context) (*User, bool) {
	u, ok := ctx.Value(userContextKey).(*User)
	return u, ok
}

// UserMustFromContext retrieves user from context and panics if not found OMIT
func UserMustFromContext(ctx context.Context) *User {
	sp, ok := ctx.Value(userContextKey).(*User)
	if !ok {
		panic("user not found in context")
	}
	return sp
}

// STOP OMIT
