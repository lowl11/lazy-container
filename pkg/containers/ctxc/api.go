package ctxc

import "context"

func Store(parent context.Context, key string, value any) context.Context {
	parentCtx := parent
	if parentCtx == nil {
		parentCtx = context.Background()
	}

	return context.WithValue(parent, key, value)
}

func Get(ctx context.Context, key string) any {
	return ctx.Value(key)
}
