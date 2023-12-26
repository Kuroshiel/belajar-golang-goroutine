package belajar_golang_goroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// time.Timer

func TestTimer(t *testing.T) {
	timer := time.NewTimer(5 * time.Second)
	fmt.Println(time.Now())

	time := <-timer.C
	fmt.Println(time)
}

// time.After

func TestAfter(t *testing.T) {
	channel := time.After(2 * time.Second)
	fmt.Println(time.Now())

	tick := <-channel
	fmt.Println(tick)
}

// time.AfterFunc()

func TestAfterFunc(t *testing.T) {
	group := sync.WaitGroup{}
	group.Add(1)

	time.AfterFunc(5*time.Second, func() {
		fmt.Println(time.Now())
		group.Done()
	})
	fmt.Println(time.Now())

	group.Wait()
}
