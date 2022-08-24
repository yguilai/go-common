package routines

import (
	"fmt"
	"runtime/debug"
)

func Go(fn func()) {
	go func() {
		defer func() {
			if err := recover(); err != nil {
				// TODO use logger
				fmt.Printf("%s\n%s", err, string(debug.Stack()))
			}
		}()
		fn()
	}()
}
