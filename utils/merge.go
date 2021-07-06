package utils

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"runtime"

	maxheap "github.com/AlexDespod/sortingmodule/max-heap"
)

type InFiles struct {
	file   *os.File
	reader *bufio.Reader
}

func MergeChunks(dirForChunks string, outFile *os.File) error {

	filesInfo, err := ioutil.ReadDir(dirForChunks)

	if err != nil {
		panic(err)
	}

	k := len(filesInfo)

	inFiles := make([]InFiles, k)

	Heap := maxheap.NewMaxHeap(k + 1)

	Heap.BuildMaxHeap()

	for i, file := range filesInfo {

		file, err := GetFile(dirForChunks + "\\" + file.Name())
		if err != nil {
			return err
		}
		defer file.Close()
		inFiles[i].file = file
		inFiles[i].reader = bufio.NewReader(file)

	}
	// memStat()
	err = insertOnce(Heap, &inFiles, k)
	// memStat()
	if err != nil {
		return err
	}

	count := 0

	for count != k {
		root, err := Heap.Remove()

		if err != nil {
			continue
		}

		outFile.WriteString(root.Line)

		item, err := ReadOneLine(inFiles[root.IndexOfFile].reader)
		if err != nil {
			if err == io.EOF {
				count++
				continue
			}
			return err
		}

		Heap.Insert(item, root.IndexOfFile)
	}

	return nil
}

func insertOnce(Heap *maxheap.Maxheap, inFiles *[]InFiles, k int) error {
	for j := 0; j < k; j++ {
		item, err := ReadOneLine((*inFiles)[j].reader)
		if err != nil {
			if err == io.EOF {
				continue
			}
			return err
		}
		err = Heap.Insert(item, j)
		if err != nil {
			continue
		}
	}
	return nil
}

func memStat() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Println(m.Alloc)

}
