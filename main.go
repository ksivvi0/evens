package main

import (
	"fmt"
	"time"
)

const (
	N               int = 100000
	workerPoolCount int = 3
)

var data []int

type result struct {
	name string
	t    float32
	data []int
}

func init() {
	/*
		fill in the data array
		just 0,1,2,3...N
	*/
	data = make([]int, N)
	for i := 0; i < N; i++ {
		data[i] = i
	}
}

//evenMemoryWithAppend - counting even numbers with append to an empty array
func evenMemoryWithAppend(arr []int, done chan result) {
	out := make([]int, 0)
	start := time.Now()
	for i := 0; i < len(arr); i++ {
		if arr[i]%2 == 0 {
			out = append(out, arr[i])
		}
	}
	done <- result{"memory with appending", float32(time.Since(start).Seconds()), out}
}

//evenMemoryWithIndex - counting even numbers with an access by index to array of zeros
func evenMemoryWithIndex(arr []int, done chan result) {
	out := make([]int, len(arr))
	start := time.Now()
	curIndex := 0
	for i := 0; i < len(arr); i++ {
		if arr[i]%2 == 0 {
			out[curIndex] = arr[i]
			curIndex++
		}
	}
	out = out[:curIndex]
	done <- result{"memory with index access", float32(time.Since(start).Seconds()), out}
}

//evenPointerWithAppend - counting even numbers with deleting elements from array received by reference
func evenPointerWithAppend(arr *[]int, done chan result) {
	start := time.Now()
	for i := 0; i < len(*arr); i++ {
		if (*arr)[i]%2 != 0 {
			*arr = append((*arr)[:i], (*arr)[i+1:]...)
		}
	}
	done <- result{"with pointer and appending", float32(time.Since(start).Seconds()), nil}
}

func main() {
	results := make(chan result, workerPoolCount)

	copyData := make([]int, N)
	copy(copyData, data)

	go evenMemoryWithAppend(data, results)
	go evenMemoryWithIndex(data, results)
	go evenPointerWithAppend(&copyData, results)

	resultCount := 0
	for r := range results {
		resultCount++
		if resultCount == workerPoolCount {
			close(results)
		}
		if r.data != nil {
			fmt.Printf("Function: %s, time: %f (%v...)\n", r.name, r.t, r.data[:10])
			continue
		}
		fmt.Printf("Function: %s, time: %f (%v...) \n", r.name, r.t, data[:10])
	}
}
