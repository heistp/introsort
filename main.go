package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sort"
	"time"
)

// Elements sets the number of elements to sort.
const Elements = 10 * 1024 * 1024

// main function
func main() {
	els := run(false)
	elc := run(true)
	fmt.Printf("cores: %d\n", runtime.NumCPU())
	fmt.Printf("concurrency speedup: %.3f\n", float64(els)/float64(elc))
}

// run runs and times a sort, with or without concurrency.
func run(concurrent bool) time.Duration {
	data := make([]int, Elements)
	for i := 0; i < len(data); i++ {
		data[i] = rand.Int()
	}
	s := Introsort{Concurrent: concurrent}
	start := time.Now()
	s.Sort(data)
	end := time.Now()
	elapsed := end.Sub(start)
	fmt.Println("number of elements:", Elements)
	fmt.Println("sorted successfully:", sort.IntsAreSorted(data))
	fmt.Println("concurrent sort:", s.Concurrent)
	fmt.Println("elapsed:", elapsed)
	fmt.Println("-----")
	return elapsed
}
