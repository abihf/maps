package maps

func EachErr[Key comparable, Val any](m map[Key]Val, fn func(Key, Val) error) error {
	for key, val := range m {
		err := fn(key, val)
		if err != nil {
			return &Error[Key, Val]{err, key, val}
		}
	}
	return nil
}
