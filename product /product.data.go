package product

import (
	"sync"
)

// used to hold our product list in memory
var productMap = struct {
	sync.RWMutex
	m map[int]Product
}{m: make(map[int]product)}