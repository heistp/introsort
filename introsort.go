package main

import "sync"

// ConcurrentCutoff sets the number of elements/sub-elements above which
// concurrency is used.
const ConcurrentCutoff = 8 * 1024

// Introsort implements a concurrent version of the introsort algorithm.
type Introsort struct {
	Concurrent bool
	wg         sync.WaitGroup
}

// Sort sorts the data in place using introsort.
func (s *Introsort) Sort(data []int) {
	maxdepth := maxDepth(len(data))
	s.introSort(data, maxdepth)
	if s.Concurrent {
		s.wg.Wait()
	}
}

// introSort is the recursive sort function
func (s *Introsort) introSort(data []int, maxdepth int) {
	// done
	if len(data) <= 1 {
		return
	}

	// heapsort
	if maxdepth == 0 {
		heapSort(data)
		return
	}

	// quicksort
	p := partitionLumuto(data)
	//p := partitionHoare(data)

	if s.Concurrent && len(data) > ConcurrentCutoff {
		s.wg.Add(2)
		go func() {
			defer s.wg.Done()
			s.introSort(data[:p], maxdepth-1) // Lumuto
			//s.introSort(data[:p+1], maxdepth-1) // Hoare
		}()
		go func() {
			defer s.wg.Done()
			s.introSort(data[p+1:], maxdepth-1)
		}()
	} else {
		s.introSort(data[:p], maxdepth-1) // Lumuto
		//s.introSort(data[:p+1], maxdepth-1) // Hoare
		s.introSort(data[p+1:], maxdepth-1)
	}
}

// maxDepth returns 2*ceil(log2(n+1))
func maxDepth(n int) (depth int) {
	for i := n; i > 0; i >>= 1 {
		depth++
	}
	depth *= 2
	return
}

// partitionLumuto uses the Lumuto partition scheme
func partitionLumuto(data []int) int {
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

// partitionHoare uses the Hoare partition scheme
func partitionHoare(data []int) int {
	hi := len(data) - 1
	pivot := data[hi/2]
	i := -1
	j := hi + 1
	for {
		i++
		for data[i] < pivot {
			i++
		}
		j--
		for data[j] > pivot {
			j--
		}
		if i >= j {
			return j
		}
		data[i], data[j] = data[j], data[i]
	}
}

// heapSort sorts a slice in-place using heap sort
func heapSort(data []int) {
	// build heap
	for i := len(data) / 2; i >= 0; i-- {
		heapify(data, i, len(data))
	}

	for length := len(data); length > 1; length-- {
		// remove top
		last := length - 1
		data[0], data[last] = data[last], data[0]
		heapify(data, 0, last)
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
