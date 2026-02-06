package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	PartitionImage()
	ProcessImageFrames()

	elapsed := time.Since(start)
	fmt.Println("execution time: \n", elapsed)
}
