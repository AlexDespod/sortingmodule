package main

import (
	"fmt"
	"time"

	sort "github.com/AlexDespod/sortingmodule/sorter"
)

func main() {
	now := time.Now()
	// err := utils.GenerateRandNums(60000000, "test.txt")
	// check(err)

	err := sort.WorkerPoolSort("test.txt", 5000000, 12)
	// file, err := os.OpenFile("outfile.txt", os.O_CREATE|os.O_WRONLY, 0666)
	check(err)

	// utils.MergeChunks(file)
	after := time.Since(now)
	fmt.Println(after)

}
func check(err error) {
	if err != nil {
		fmt.Println(err)
	}

}
