package main

import (
	"github.com/kataras/pio"
	"os"
)

func main(){
	p := pio.NewTextPrinter("test", os.Stdout)
	p.Println([]byte("test"))
}