package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

const Elements = 1024 * 1024

// main function
func main() {
	data := make([]int, Elements)
	for i := 0; i < len(data); i++ {
		data[i] = rand.Int()
	}
	start := time.Now()
	introSort(data)
	end := time.Now()
	fmt.Println("elements:", Elements)
	fmt.Println("sorted:", sort.IntsAreSorted(data))
	fmt.Println("elapsed:", end.Sub(start))
}

// introSort is the main sort function
func introSort(data []int) {
	maxdepth := maxDepth(len(data))
	introSortDepth(data, maxdepth)
}

// introSortDepth is the recursive sort function
func introSortDepth(data []int, maxdepth int) {
	if len(data) <= 1 {
		return
	}
	if maxdepth == 0 {
		heapSort(data)
	} else {
		p := partition(data)
		introSortDepth(data[:p], maxdepth-1)
		introSortDepth(data[p+1:], maxdepth-1)
	}
}

// partition uses the Lumuto partition scheme
func partition(data []int) int {
	hi := len(data) - 1
	pivot := data[hi]
	i := 0
	for j := 0; j <= hi; j++ {
		if data[j] < pivot {
			data[i], data[j] = data[j], data[i]
			i++
		}
	}
	data[i], data[hi] = data[hi], data[i]
	return i
}

// maxDepth returns 2*ceil(log2(n+1))
func maxDepth(n int) int {
	depth := 0
	for i := n; i > 0; i >>= 1 {
		depth++
	}
	return depth * 2
}

// heapSort sorts the slice using heap sort
func heapSort(data []int) {
	// build heap
	for i := len(data) / 2; i >= 0; i-- {
		heapify(data, i, len(data))
	}

	for length := len(data); length > 1; length-- {
		// remove top
		lastIndex := length - 1
		data[0], data[lastIndex] = data[lastIndex], data[0]
		heapify(data, 0, lastIndex)
	}
}

// heapify builds heap from a list
func heapify(data []int, root, length int) {
	max := root
	l, r := 2*root+1, 2*root+2

	if l < length && data[l] > data[max] {
		max = l
	}

	if r < length && data[r] > data[max] {
		max = r
	}

	if max != root {
		data[root], data[max] = data[max], data[root]
		heapify(data, max, length)
	}
}
