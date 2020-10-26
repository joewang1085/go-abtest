package sdk

import (
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
