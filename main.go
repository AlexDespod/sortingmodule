package main

import (
	"fmt"
	"os"
	"time"

	sort "github.com/AlexDespod/sortingmodule/sorter"
	"github.com/AlexDespod/sortingmodule/utils"
)

func main() {
	now := time.Now()

	err := utils.GenerateRandNums(100, "tests/test.txt")
	check(err)
	err = sort.SortFile("tests\\test.txt", "tests", 10)
	//err := sort.WorkerPoolSort("test.txt", 5000000, 12)
	check(err)
	file, err := os.OpenFile("tests\\outfile.txt", os.O_CREATE|os.O_WRONLY, 0666)
	check(err)

	utils.MergeChunks("tests\\chunks", file)
	after := time.Since(now)
	fmt.Println(after)

}
func check(err error) {
	if err != nil {
		panic(err)
	}

}
