package sdk

import (
	"fmt"
	"sync"
)

// labOutputCache for demo test, ABT lab output cache
var labOutputCache sync.Map

// Push ...
func outputPush(k, v interface{}) {
	labOutputCache.Store(k, v)
}

// Range ..
func outputRange(f func(k, v interface{}) bool) {
	labOutputCache.Range(f)
}

// DebugOutput ....
func DebugOutput() {
	labOutputCache.Range(
		func(k, v interface{}) bool {
			fmt.Println("debug labOutputCache", k, v)
			return true
		},
	)
}
