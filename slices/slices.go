package slices

func AnyMatch[T any](slice []T, matcher func(index int, val T) bool) bool {
	for i, val := range slice {
		if matcher(i, val) {
			return true
		}
	}
	return false
}
