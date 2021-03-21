package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	loadFile()
}
func csvRead() {
	file, err := os.Open("csv.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	titleSlice := make([]string, 0)
	priceSlice := make([]float64, 0)
	quantitySlice := make([]int64, 0)

	for {
		str, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			panic(err)
		}
		if err == io.EOF {
			break
		}

		strs := strings.Split(strings.TrimSpace(str), ";")

		title := strs[0]
		titleSlice = append(titleSlice, title)

		price, err := strconv.ParseFloat(strs[1], 64)
		if err != nil {
			panic(err)
		}
		priceSlice = append(priceSlice, price)

		quantity, err := strconv.ParseInt(strs[2], 10, 64)
		if err != nil {
			panic(err)
		}
		quantitySlice = append(quantitySlice, quantity)
	}

	fmt.Println(titleSlice)
	fmt.Println(priceSlice)
	fmt.Println(quantitySlice)
}

type Page struct {
	Title string
	Body  []byte
}

func readAllFile() {
	buf, err := ioutil.ReadFile("file.txt")
	if err != nil {
		panic(err)
	}

	fmt.Print(string(buf))
}

func loadFile() {
	file, err := os.Open("file.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			panic(err)
		}
		if err == io.EOF && len(str) == 0 {
			break
		}

		fmt.Print(str)
	}

}
