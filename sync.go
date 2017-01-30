package main

import (
	"sync"
)

// Info sync struct
type Info struct {
	mu   sync.Mutex
	name string
}

func syncUpdate(info *Info) {
	info.mu.Lock()
	info.name = "" // new value
	info.mu.Unlock()
}
