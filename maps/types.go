package maps

type (
	Map[K comparable, V any] interface {
		Get(key K) V
		Put(key K, value V) V
		Remove(key K) V
		Keys() []K
		Values() []V
		Entries() []Entry[K, V]
		GetOrDefault(key K, defaultValue V) V
		PutIfAbsent(key K, value K)
		ComputeIfAbsent(key K, mappingFn Mapping[K, V])
		ComputeIfPresent(key K, remappingFn ReMapping[K, V])
	}

	Entry[K any, V any] struct {
		Key   K
		Value V
	}

	Mapping[K comparable, V any] func(key K) V

	ReMapping[K comparable, V any] func(key K, oldValue V) V
)
