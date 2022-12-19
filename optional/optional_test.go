package optional

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOf(t *testing.T) {
	var foo int64 = 1
	op := OfNil(&foo)

	assert.True(t, *op.Get() == 1)
}
