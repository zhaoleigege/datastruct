package main

import (
    "fmt"
)

func main(){
    var a = 2
    fmt.Printf("%t\n",a)
    fmt.Println(float64(a) / 100.0)


    fmt.Printf("%f\n", float64(int64(20) / 100.0))
	fmt.Println("5.34.1" < "5.35.0")
	fmt.Println(len([]byte("中文")))
}
