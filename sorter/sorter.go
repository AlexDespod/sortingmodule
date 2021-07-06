package utils

import (
	"bufio"
	"io"
	"sort"

	"github.com/AlexDespod/sortingmodule/utils"
	"github.com/AlexDespod/sortingmodule/workerpool"
)

func SortFile(fileName, chunks string, PerChunk int) error {

	inFile, err := utils.GetFile(fileName)

	if err != nil {
		inFile.Close()
		return err
	}
	defer inFile.Close()

	reader := bufio.NewReader(inFile)

	i := 0

	for {

		intBuff, err := utils.GetSortedData(reader, PerChunk)

		if err != nil && err != io.EOF {

			return err
		}

		sort.Slice(intBuff, func(i, j int) bool {
			return intBuff[i].Num > intBuff[j].Num
		})

		err1 := utils.MakeChunkFile(i, &intBuff, chunks)

		if err1 != nil {
			return err1
		}

		i++

		if err == io.EOF {
			break
		}
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
