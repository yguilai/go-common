package maps

type Entry[K any, V any] struct {
	Key   K
	Value V
}

func Keys[K comparable, V any](m map[K]V) []K {
	keys := make([]K, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

func Values[K comparable, V any](m map[K]V) []V {
	values := make([]V, len(m))
	for _, v := range m {
		values = append(values, v)
	}
	return values
}

func Entries[K comparable, V any](m map[K]V) []Entry[K, V] {
	entries := make([]Entry[K, V], len(m))
	for k, v := range m {
		entries = append(entries, Entry[K, V]{
			Key:   k,
			Value: v,
		})
	}
	return entries
}

func Mapping[K comparable, V any, U any](m map[K]V, fn func(key K, value V) U) []U {
	us := make([]U, len(m))
	for k, v := range m {
		us = append(us, fn(k, v))
	}
	return us
}
