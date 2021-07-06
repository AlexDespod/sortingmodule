package utils

import (
	"bufio"
	"fmt"
	"io"
	"sort"
	"sync"

	"github.com/AlexDespod/sortingmodule/utils"
	"github.com/AlexDespod/sortingmodule/workerpool"
)

func SortFile(fileName string, PerChunk int) error {

	inFile, err := utils.GetFile(fileName)

	if err != nil {
		inFile.Close()
		return err
	}
	defer inFile.Close()

	fail := make(chan error, 1)

	reader := bufio.NewReader(inFile)

	wg := new(sync.WaitGroup)

	i := 0

	for {

		intBuff, err := utils.GetSortedData(reader, PerChunk)

		if err != nil && err != io.EOF {

			return err
		}

		sort.Slice(intBuff, func(i, j int) bool {
			return intBuff[i].Num > intBuff[j].Num
		})

		go func(i int) {
			err1 := utils.MakeChunkFile(i, &intBuff)

			if err1 != nil {
				fail <- err1
				wg.Done()
				return
			}
			wg.Done()
		}(i)

		i++

		if err != io.EOF {
			break
		}
	}
	wg.Add(i + 1)

	wg.Wait()

	close(fail)

	errState := len(fail)

	fmt.Println(errState)

	if errState != 0 {

		err = <-fail

		fmt.Println(err)

		return err
	}
	return nil
}

func WorkerPoolSort(fileName string, perChunk int, concurrency int) error {
	config := workerpool.Config{Filename: fileName, PerChunk: perChunk, Concurrency: concurrency}
	pool := workerpool.NewPool(config)
	err := pool.Run()
	if err != nil {
		return err
	}
	return nil
}
