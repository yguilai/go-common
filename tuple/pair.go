package tuple

type defaultPair[L any, R any] struct {
	left  L
	right R
}

func (p *defaultPair[L, R]) Left() L {
	return p.left
}

func (p *defaultPair[L, R]) Right() R {
	return p.right
}

func PairOf[L any, R any](l L, r R) Pair[L, R] {
	return &defaultPair[L, R]{l, r}
}
