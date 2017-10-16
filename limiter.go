package main

import (
	"fmt"
	"sync"
	"time"
)

type Limiter struct {
	usersMax   uint
	interv     uint
	usersCur   uint
	resetSched bool
	cond       sync.Cond
}

func NewLimiter(max, interv uint) (l *Limiter) {
	return &Limiter{max, interv, 0, false, sync.Cond{L: &sync.RWMutex{}}}
}

func (l Limiter) Aquire() {
	if !l.resetSched {
		l.resetSched = true
		go l.reset()
	}
	fmt.Println("Aquired")
}

func (l Limiter) reset() {
	time.Sleep(time.Duration(l.interv) * time.Second)
	l.cond.L.Lock()
	defer l.cond.L.Unlock()
	l.usersCur = 0
	l.resetSched = false
	l.cond.Broadcast()
	fmt.Println("Reset")
}
