package main

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
	"time"
)

func almostEqual(a, b float64) bool {
	return math.Abs(a-b) <= 0.01
}

func TestSolution(t *testing.T) {
	a := make(chan string, 1)
	b := make(chan string, 1)
	Solution(2*time.Second, "ali", a, b)
	assert.Equal(t, "ali", <-a)
	assert.Equal(t, "ali", <-b)
}

func TestSolution2(t *testing.T) {
	a := make(chan string, 1)
	b := make(chan string, 1)
	a = nil
	start := time.Now()
	Solution(2*time.Second, "ali", a, b)
	delta := time.Since(start)
	assert.Equal(t, true, almostEqual(delta.Seconds(), 2.0))
	assert.Equal(t, "ali", <-b)
}

func TestSolution3(t *testing.T) {
	start := time.Now()
	a := make(chan string, 1)
	b := make(chan string, 1)
	a <- "full"
	go func() { time.Sleep(1 * time.Second); <-a }()
	Solution(2*time.Second, "ali", a, b)
	assert.Equal(t, "ali", <-b)
	delta := time.Since(start)
	assert.Equal(t, true, almostEqual(delta.Seconds(), 1.0))
}

func TestSolution4(t *testing.T) {
	start := time.Now()
	a := make(chan string, 1)
	b := make(chan string, 1)
	c := make(chan string, 0)
	d := make(chan string, 1)
	x := Solution(3*time.Second, "ali", a, b, c, d)
	assert.Equal(t, "ali", <-b)
	assert.Equal(t, "ali", <-a)
	assert.Equal(t, "ali", <-d)
	assert.Equal(t, 3, x)
	delta := time.Since(start)
	assert.Equal(t, true, almostEqual(delta.Seconds(), 3.0))
}

func TestSolution5(t *testing.T) {
	a := make(chan string, 10)
	b := make(chan string, 0)
	c := make(chan string, 5)
	x := Solution(1*time.Second, "salam", a, b, c)
	assert.Equal(t, 1, len(a))
	assert.Equal(t, 1, len(c))
	assert.Equal(t, 0, len(b))
	assert.Equal(t, "salam", <-a)
	assert.Equal(t, 2, x)
}

func TestSolution6(t *testing.T) {
	//okay we got three buffered channels
	a := make(chan string, 8)
	c := make(chan string, 4)
	b := make(chan string, 1)
	b <- "full"                                      //putting full in b
	go func() { time.Sleep(2 * time.Second); <-b }() //after two seconds removing full
	x := Solution(1*time.Second, "salam", a, b, c)   // a has salam. c has salam but b has full
	start := time.Now()
	Solution(2*time.Second, "hi", a, b, c) //now b has only hi in it. a has both.
	// actually we are working with queue so, it is a fifo.
	delta := time.Since(start)
	assert.Equal(t, "hi", <-b)
	assert.Equal(t, "salam", <-a)
	assert.Equal(t, 2, x)
	assert.Equal(t, true, almostEqual(delta.Seconds(), 1.0))
}
