package maps

func EachSync[Key comparable, Val any](m map[Key]Val, fn func(Key, Val)) {
	for key, val := range m {
		fn(key, val)
	}
}
