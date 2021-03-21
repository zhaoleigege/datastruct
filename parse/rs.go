package main

import (
	"fmt"
	"reflect"
)

func main(){
	tag := reflect.StructTag(`species:"gopher" color:"red"`)
	fmt.Println(tag.Get("color"))
}
