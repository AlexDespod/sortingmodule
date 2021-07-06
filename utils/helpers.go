package utils

import (
	"bufio"
	"context"
	"io"
	"os"
	"sort"
	"strconv"

	"github.com/AlexDespod/sortingmodule/structs"
)

func GetSortedData(reader *bufio.Reader, size int) ([]structs.SortItem, error) {

	intBuff := make([]structs.SortItem, size)

	for i := 0; i < size; i++ {

		line, err1 := reader.ReadString('\n')

		if err1 != nil {
			if err1 == io.EOF {
				return intBuff[:i], err1
			}
			return nil, err1
		}

		dataLen := len(line)

		num, err := strconv.Atoi(line[:dataLen-1])

		if err != nil {
			panic(err)
		}

		intBuff[i] = structs.SortItem{Num: num, Line: line}

	}
	return intBuff, nil
}

func ProcessDataAsync(ctx context.Context, cancel context.CancelFunc, dataChan chan structs.DataChanItem, chunksDir string, size int, id int) error {

	intBuff := make([]structs.SortItem, size)
	var i int
	for i < size {

		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		item := <-dataChan

		//here a test of meaning error

		// if id == 9 && i == 1000000 {
		// 	fmt.Println("test error ", id)
		// 	return fmt.Errorf("test error %v", id)
		// }
		//
		if item.Err != nil {

			if item.Err == io.EOF {

				if item.Line != "" {

					intBuff[i] = item.SortItem
					break
				}
				break
			}
			cancel()
			return item.Err
		}
		intBuff[i] = item.SortItem
		i++
	}

	sort.Slice(intBuff, func(i, j int) bool {
		return intBuff[i].Num > intBuff[j].Num
	})

	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	err := MakeChunkFile(id, &intBuff, chunksDir)

	//here a test of meaning error
	//
	// if id == 10 && i == 1000000 {
	// 	fmt.Println("test error ", id)
	// 	return fmt.Errorf("test error %v", id)
	// }
	//
	if err != nil {
		return err
	}
	return nil
}

func ReadOneLine(reader *bufio.Reader) (structs.SortItem, error) {

	str, err := reader.ReadString('\n')
	if err != nil {

		if err == io.EOF {

			if str != "" {
				num, err := getNumber(str)

				if err != nil {
					return structs.SortItem{}, err
				}

				return structs.SortItem{Num: num, Line: str}, err
			}

		}
		return structs.SortItem{}, err
	}

	num, err := getNumber(str)
	if err != nil {
		return structs.SortItem{}, err
	}
	return structs.SortItem{Num: num, Line: str}, nil

}

func writeSortedData(intBuff *[]structs.SortItem, outFile *os.File) error {
	for _, val := range *intBuff {
		_, err := outFile.WriteString(val.Line)
		if err != nil {
			return err
		}
	}
	return nil
}

func MakeChunkFile(numName int, data *[]structs.SortItem, dir string) error {
	file, err := createFile(dir + "\\chunks\\" + strconv.Itoa(numName))
	if err != nil {
		return err
	}
	defer file.Close()
	err = writeSortedData(data, file)
	if err != nil {
		return err
	}
	return nil
}

func GetFile(name string) (*os.File, error) {
	dir, _ := os.Getwd()
	file, err := os.OpenFile(dir+"\\"+name, os.O_RDWR, 0666)
	if err != nil {
		file.Close()
		return nil, err
	}

	return file, nil
}
func createFile(name string) (*os.File, error) {
	dir, _ := os.Getwd()
	file, err := os.Create(dir + "\\" + name)
	if err != nil {
		file.Close()
		return nil, err
	}

	return file, nil
}
func getNumber(str string) (int, error) {
	dataLen := len(str)
	num, err := strconv.Atoi(str[:dataLen-1])
	if err != nil {
		return 0, err
	}
	return num, nil
}
