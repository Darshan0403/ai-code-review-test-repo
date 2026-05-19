package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"os"
)

// ProcessData reads a file and hashes it
func ProcessData(filename string) {
	// 1. RESOURCE LEAK: Opens a file but never closes it
	file, _ := os.Open(filename)

	// 2. CONCURRENCY BUG: This goroutine will leak because nobody reads from the channel
	ch := make(chan int)
	go func() {
		fmt.Println("Doing background work...")
		ch <- 42 // Will block forever here
	}()

	// 3. SECURITY VULNERABILITY: MD5 is cryptographically broken
	hash := md5.New()
	io.Copy(hash, file)

	fmt.Printf("File hash: %x\n", hash.Sum(nil))
}
