package maps

import "sync"

func Transform[Key comparable, In, Out any](input map[Key]In, transformer func(Key, In) Out) map[Key]Out {
	res := make(map[Key]Out, len(input))
	var mu sync.Mutex
	Each(input, func(k Key, v In) {
		out := transformer(k, v)
		mu.Lock()
		defer mu.Unlock()
		res[k] = out
	})
	return res
}
