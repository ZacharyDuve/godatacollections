package tree

func containsInSlice[T comparable](t T, ts []T) bool {
	for _, curT := range ts {
		if curT == t {
			return true
		}
	}

	return false
}
