package maps

type DefaultMap[K comparable, V any] map[K]V

func (m DefaultMap[K, V]) Get(key K) V {
	return m[key]
}

func (m DefaultMap[K, V]) Put(key K, value V) V {
	m[key] = value
	return value
}

func (m DefaultMap[K, V]) Remove(key K) V {
	v := m[key]
	delete(m, key)
	return v
}

func (m DefaultMap[K, V]) Keys() []K {
	keys := make([]K, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

func (m DefaultMap[K, V]) Values() []V {
	values := make([]V, len(m))
	for _, v := range m {
		values = append(values, v)
	}
	return values
}

func (m DefaultMap[K, V]) Entries() []Entry[K, V] {
	entries := make([]Entry[K, V], len(m))
	for k, v := range m {
		entries = append(entries, Entry[K, V]{
			Key:   k,
			Value: v,
		})
	}
	return entries
}

func (m DefaultMap[K, V]) GetOrDefault(key K, defaultValue V) V {
	v := m[key]
	if v == nil {
		return defaultValue
	}
	return v
}

func (m DefaultMap[K, V]) PutIfAbsent(key K, value K) {
	//TODO implement me
	panic("implement me")
}

func (m DefaultMap[K, V]) ComputeIfAbsent(key K, mappingFn Mapping[K, V]) {
	//TODO implement me
	panic("implement me")
}

func (m DefaultMap[K, V]) ComputeIfPresent(key K, remappingFn ReMapping[K, V]) {
	//TODO implement me
	panic("implement me")
}

func Wrap[K comparable, V any](m map[K]V) Map[K, V] {
	return DefaultMap[K, V](m)
}
