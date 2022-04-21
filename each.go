package maps

import "sync"

func Each[Key comparable, Val any](m map[Key]Val, fn func(Key, Val)) {
	var wg sync.WaitGroup
	EachSync(m, func(k Key, v Val) {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fn(k, v)
		}()
	})
	wg.Wait()
}
