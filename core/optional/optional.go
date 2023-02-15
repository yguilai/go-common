package optional

type (
	Optional[T any] struct {
		value *T
	}

	Consumer[T any] func(value *T)

	Predicate[T any] func(value *T) bool
	Supplier[T any]  func() *T
)

func Of[T any](value *T) *Optional[T] {
	if value == nil {
		panic("value can not be nil")
	}
	return &Optional[T]{
		value: value,
	}
}

func OfNil[T any](value *T) *Optional[T] {
	if value == nil {
		return newEmpty[T]()
	}
	return Of(value)
}

func newEmpty[T any]() *Optional[T] {
	return &Optional[T]{}
}

func (o *Optional[T]) empty() bool {
	return o.value == nil
}

func (o *Optional[T]) IsPresent() bool {
	return !o.empty()
}

func (o *Optional[T]) IfPreset(consumer Consumer[T]) {
	if o.value != nil {
		consumer(o.value)
	}
}

func (o *Optional[T]) Filter(predicate Predicate[T]) *Optional[T] {
	if predicate == nil {
		panic("predicate is required")
	}
	if o.empty() {
		return o
	}
	if predicate(o.value) {
		return o
	}
	return newEmpty[T]()
}

func (o *Optional[T]) Get() *T {
	return o.value
}

func (o *Optional[T]) OrElse(value *T) *T {
	if o.empty() {
		return value
	}
	return o.value
}

func (o *Optional[T]) OrElseGet(supplier Supplier[T]) *T {
	if o.empty() {
		return supplier()
	}
	return o.value
}

func (o *Optional[T]) OrElsePanic() *T {
	if o.empty() {
		panic("value is nil")
	}
	return o.value
}

func (o *Optional[T]) OrElseError(supplier Supplier[error]) (*T, error) {
	if o.empty() {
		return nil, *supplier()
	}
	return o.value, nil
}
