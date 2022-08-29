package pool

import (
	"fmt"
	"github.com/panjf2000/ants/v2"
)

func MustSinglePool(opts ...ants.Option) *ants.Pool {
	pool, err := ants.NewPool(1, opts...)
	if err != nil {
		panic(fmt.Sprintf("new single pool error: %+v", err))
	}
	return pool
}
