package logger

import "context"

type key int

const loggerKey key = 0

// NewContext sets logger to context
func NewContext(ctx context.Context, logger Logger) context.Context {
	return context.WithValue(ctx, loggerKey, logger)
}

// FromContext retrieves logger from context
func FromContext(ctx context.Context) (Logger, bool) {
	l, ok := ctx.Value(loggerKey).(Logger)

	return l, ok
}

// MustFromContext retrieves logger from context. Panics if not found
func MustFromContext(ctx context.Context) Logger {
	l, ok := ctx.Value(loggerKey).(Logger)
	if !ok {
		panic("logger not found in context")
	}

	return l
}
