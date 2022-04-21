package maps

func Copy[Key comparable, Val any](dest, src map[Key]Val) {
	for name, value := range src {
		dest[name] = value
	}
}
