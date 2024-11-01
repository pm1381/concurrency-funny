package main

import (
	"fmt"
	"math"
	"sync"
	"sync/atomic"
	"time"
)

func almostEqualMock(a, b float64) float64 {
	return math.Abs(a - b)
}

func main() {
	a := make(chan string, 1)
	b := make(chan string, 1)
	a = nil
	start := time.Now()
	Solution(2, "ali", a, b)
	delta := time.Since(start)
	fmt.Println(<-b)
	fmt.Println(almostEqualMock(delta.Seconds(), 2.0))
}

func Solution(d time.Duration, message string, ch ...chan string) (numberOfAccesses int) {
	ticker := time.NewTicker(d)
	defer ticker.Stop()
	var count int64
	var wg sync.WaitGroup
	for _, c := range ch {
		wg.Add(1)
		go func(eachChannel chan string, count *int64, ticker *time.Ticker) {
			select {
			case <-ticker.C:
				wg.Done()
			case eachChannel <- message:
				atomic.AddInt64(count, 1)
				wg.Done()
			}
		}(c, &count, ticker)
	}
	wg.Wait()
	return int(count)
}
