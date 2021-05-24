package main

import (
	"fmt"
	"sync"
	"time"
)

func ImageFile(f string) string {
	time.Sleep(1000 * time.Millisecond)
	return f + "-ok"
}

func makeThumbnails3(filenames []string) []string {
	ch := make(chan string)
	for _, f := range filenames {
		go func(f string) {
			res := ImageFile(f) // NOTE: ignoring errors
			ch <- res
		}(f)
	}
	// Wait for goroutines to complete.
	result := []string{}
	for range filenames {
		result = append(result, <-ch)
	}
	return result
}

func makeThumbnails6(filenames <-chan string) int64 {
	sizes := make(chan int64)
	var wg sync.WaitGroup // number of working goroutines
	for f := range filenames {
		wg.Add(1) // Add是为计数器加一
		// worker
		go func(f string) {
			defer wg.Done() // Done没有任何参数,其实它和Add(-1)是等价的
			ImageFile(f)
			sizes <- 1
		}(f)
	}

	// closer
	go func() {
		wg.Wait()
		close(sizes)
	}()

	var total int64
	for size := range sizes {
		total += size
	}
	return total
}

func batch1() {
	start := time.Now()
	files := []string{"a", "b", "c", "d", "e", "f", "g"}
	result := makeThumbnails3(files)
	fmt.Println(result)
	secs := time.Since(start).Seconds()
	fmt.Printf("is %f seconds", secs)
}

func batch2() {
	start := time.Now()
	f := []string{"a", "b", "c", "d", "e", "f", "g"}
	files := make(chan string)

	go func() {
		for _, v := range f {
			files <- v
		}
		close(files)
	}()

	for ff := range files {
		fmt.Println(ff)
	}

	// result := makeThumbnails6(files)
	// fmt.Println(result)

	secs := time.Since(start).Seconds()
	fmt.Printf("is %f seconds", secs)
}

func main() {
	// batch1()
	batch2()
	// ch1 := make(chan string, 1)
	// // go func() {
	// ch1 <- "hello world"
	// // }()
	// fmt.Println(<-ch1)
}
