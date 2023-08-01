package activities

import (
	"context"
	"log/slog"
)

type contextKey int

const contextKeyActivity contextKey = iota

func NewContext(ctx context.Context, a *Activity) context.Context {
	if FromContext(ctx) != nil {
		slog.Warn("skipped new context with activity as it already exists")
		return ctx
	}
	return context.WithValue(ctx, contextKeyActivity, a)
}

func FromContext(ctx context.Context) *Activity {
	s, _ := ctx.Value(contextKeyActivity).(*Activity)
	return s
}
