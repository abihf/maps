package maps

func Merge[Key comparable, Val any](inputs ...map[Key]Val) map[Key]Val {
	res := map[Key]Val{}
	for _, input := range inputs {
		Copy(res, input)
	}
	return res
}
