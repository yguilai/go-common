package tuple

type Pair[L any, R any] interface {
	Left() L
	Right() R
}
