package utils

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func GenerateRandNums(count int, filaname string) error {
	var i int
	dir, _ := os.Getwd()
	file, err := os.OpenFile(dir+"\\"+filaname, os.O_CREATE|os.O_WRONLY, 0777)

	if err != nil {

		file.Close()
		return err
	}

	defer file.Close()

	rand.Seed(time.Now().Unix())

	for i = 0; i <= count; i++ {

		temp := strconv.Itoa(rand.Int())

		fmt.Fprintln(file, temp)
		if err != nil {

			return err
		}
	}
	return nil

}
