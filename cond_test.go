package belajar_golang_goroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// sync.Cond
var locker = sync.Mutex{}
var cond = sync.NewCond(&locker)
var group = &sync.WaitGroup{}

func WaitCondition(value int) {
	defer group.Done()
	group.Add(1)

	cond.L.Lock()
	cond.Wait()
	fmt.Println("Done", value)
	cond.L.Unlock()
}

func TestCond(t *testing.T) {
	for i := 0; i < 10; i++ {
		go WaitCondition(i)
	}
	// Goroutine di jalankan 1/1 dengan cond.Signal()
	go func() {
		for i := 0; i < 10; i++ {
			time.Sleep(1 * time.Second)
			cond.Signal()
		}
	}()
	// Goroutine di jalankan semuannya dengan cond.Broadcast()
	/* go func() {
		time.Sleep(1 * time.Second)
		cond.Broadcast()
	}() */

	group.Wait()
}
