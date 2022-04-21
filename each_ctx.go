package maps

import "context"

func EachCtx[Key comparable, Val any](ctx context.Context, m map[Key]Val, fn func(context.Context, Key, Val) bool) {
	ctx, cancel := context.WithCancel(ctx)
	Each(m, func(k Key, v Val) {
		if !fn(ctx, k, v) {
			cancel()
		}
	})
}
