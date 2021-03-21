package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanf()
}

// reader 读取一行数据
func reader() {
	inputReader := bufio.NewReader(os.Stdin)
	for {
		input, err := inputReader.ReadString('\n')
		if err != nil {
			panic(err)
		}

		if "exit" == strings.TrimSpace(strings.ToLower(input)) {
			break
		}

		fmt.Print(input)
	}

}

// scan 格式化读取
func scan() {
	var name string
	var age int

	fmt.Scan(&name, &age)
	fmt.Printf("name-> %s, age -> %d\n", name, age)
}

func scanf() {
	var (
		title    string
		price    float64
		quantity int64
	)

	_, err := fmt.Sscanf("malongshuai;23", "%s ; %d", &title, &quantity)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s %f %d\n", title, price, quantity)
}
