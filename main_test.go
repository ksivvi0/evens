package main

import "testing"

func init() {
	data = make([]int, N)
	for i := 0; i < N; i++ {
		data[i] = i
	}
}

func BenchmarkEvenPointerWithAppend(b *testing.B) {
	copyData := make([]int, N)
	copy(copyData, data)
	results := make(chan result, 1)

	go evenPointerWithAppend(&copyData, results)
	<-results
}

func BenchmarkEvenMemoryWithIndex(b *testing.B) {
	results := make(chan result, 1)

	go evenMemoryWithIndex(data, results)

	<-results
}

func BenchmarkEvenMemoryWithAppend(b *testing.B) {
	b.StartTimer()
	results := make(chan result, 1)

	go evenMemoryWithAppend(data, results)

	<-results
}
