package slices

type (
	Matcher[T any]    func(index int, val T) bool
	MatcherErr[T any] func(index int, val T) error
)

func AnyMatch[T any](slice []T, matcher Matcher[T]) bool {
	for i, val := range slice {
		if matcher(i, val) {
			return true
		}
	}
	return false
}

func AnyError[T any](slice []T, matcher MatcherErr[T]) error {
	for i, val := range slice {
		if err := matcher(i, val); err != nil {
			return err
		}
	}
	return nil
}
