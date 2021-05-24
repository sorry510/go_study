package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
)

// 计算目录的文件大小
func walkDir(dir string, fileSizes chan<- int64) {
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			subdir := filepath.Join(dir, entry.Name())
			walkDir(subdir, fileSizes)
		} else {
			fileSizes <- entry.Size()
		}
	}
}

func dirents(dir string) []os.FileInfo {
	entries, err := ioutil.ReadDir(dir) // 返回一个os.FileInfo类型的slice
	if err != nil {
		fmt.Fprintf(os.Stderr, "du1: %v\n", err)
		return nil
	}
	return entries
}

func printDiskUsage(nfiles, nbytes int64) {
	fmt.Printf("%d files %.1f GB\n", nfiles, float64(nbytes)/1e9)
}

// func main() {
// 	flag.Parse()
// 	roots := flag.Args() // 接受命令行参数
// 	if len(roots) == 0 {
// 		roots = []string{"."}
// 	}

// 	fileSizes := make(chan int64)
// 	go func() {
// 		for _, root := range roots {
// 			walkDir(root, fileSizes)
// 		}
// 		close(fileSizes)
// 	}()

// 	var nfiles, nbytes int64
// 	for size := range fileSizes {
// 		nfiles++
// 		nbytes += size
// 	}
// 	printDiskUsage(nfiles, nbytes)
// }

// var verbose = flag.Bool("v", true, "show verbose progress messages")
// func main() {
// 	flag.Parse()
// 	roots := flag.Args() // 接受命令行参数
// 	if len(roots) == 0 {
// 		roots = []string{"."}
// 	}

// 	fileSizes := make(chan int64)
// 	go func() {
// 		for _, root := range roots {
// 			walkDir(root, fileSizes)
// 		}
// 		close(fileSizes)
// 	}()

// 	var tick <-chan time.Time
// 	if *verbose {
// 		tick = time.Tick(500 * time.Millisecond)
// 	}
// 	var nfiles, nbytes int64
// loop:
// 	for {
// 		select {
// 		case size, ok := <-fileSizes:
// 			if !ok {
// 				break loop // fileSizes was closed
// 			}
// 			nfiles++
// 			nbytes += size
// 		case <-tick:
// 			printDiskUsage(nfiles, nbytes)
// 		}
// 	}
// 	printDiskUsage(nfiles, nbytes) // final totals
// }

func main() {
	flag.Parse()
	roots := flag.Args() // 接受命令行参数
	if len(roots) == 0 {
		roots = []string{"."}
	}
	fileSizes := make(chan int64)
	var n sync.WaitGroup
	for _, root := range roots {
		n.Add(1)
		go walkDir2(root, &n, fileSizes)
	}
	go func() {
		n.Wait()
		close(fileSizes)
	}()
	var nfiles, nbytes int64
	for size := range fileSizes {
		nfiles++
		nbytes += size
	}
	printDiskUsage(nfiles, nbytes)
}

func walkDir2(dir string, n *sync.WaitGroup, fileSizes chan<- int64) {
	defer n.Done()
	for _, entry := range dirents2(dir) {
		if entry.IsDir() {
			n.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			go walkDir2(subdir, n, fileSizes)
		} else {
			fileSizes <- entry.Size()
		}
	}
}

var sema = make(chan struct{}, 20)

func dirents2(dir string) []os.FileInfo {
	sema <- struct{}{}                  // acquire token
	defer func() { <-sema }()           // release token
	entries, err := ioutil.ReadDir(dir) // 返回一个os.FileInfo类型的slice
	if err != nil {
		fmt.Fprintf(os.Stderr, "du1: %v\n", err)
		return nil
	}
	return entries
}
