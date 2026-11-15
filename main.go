package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	SIZE        = 100_000_000
	CHUNKS      = 8
	randomRange = 10000
)

// generateRandomElements generates random elements.
func generateRandomElements(size int) []int {
	if size <= 0 {
		return nil
	}

	nums := make([]int, size)
	for i := range nums {
		nums[i] = rand.Intn(randomRange) + 1
	}
	return nums
}

// maximum returns the maximum number of elements.
func maximum(data []int) int {
	if len(data) == 0 {
		return 0
	}

	if len(data) == 1 {
		return data[0]
	}

	res := data[0]
	for _, v := range data {
		if v > res {
			res = v
		}
	}
	return res
}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) int {
	if len(data) == 0 {
		return 0
	}

	// Рассчитываем границы частей слайсов от исходного массива
	counter := 0
	chunkSize := len(data) / CHUNKS
	chunkIdx := make([]int, 0, CHUNKS+1)
	for range CHUNKS {
		chunkIdx = append(chunkIdx, counter)
		counter += chunkSize
	}
	chunkIdx = append(chunkIdx, len(data))

	var wg sync.WaitGroup
	// Буферизированный канал для записи ответов всех горутин
	chunkChannel := make(chan int, CHUNKS)

	for i := range CHUNKS {
		wg.Add(1)
		go func(data []int, ch chan int) {
			defer wg.Done()
			res := maximum(data)
			ch <- res
		}(data[chunkIdx[i]:chunkIdx[i+1]], chunkChannel)
	}

	wg.Wait()
	close(chunkChannel)
	channelValues := make([]int, 0, CHUNKS)
	for v := range chunkChannel {
		channelValues = append(channelValues, v)
	}

	return maximum(channelValues)
}

func main() {
	fmt.Printf("Генерируем %d целых чисел\n", SIZE)
	generatedSlice := generateRandomElements(SIZE)

	fmt.Printf("Ищем максимальное значение в один поток\n")
	start := time.Now()
	max := maximum(generatedSlice)
	t := time.Now()
	elapsed := t.Sub(start).Microseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d microseconds\n", max, elapsed)

	fmt.Printf("Ищем максимальное значение в %d потоков\n", CHUNKS)
	start = time.Now()
	max = maxChunks(generatedSlice)
	t = time.Now()
	elapsed = t.Sub(start).Microseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d microseconds\n", max, elapsed)
}
