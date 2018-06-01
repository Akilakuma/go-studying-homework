package main

import (
	"fmt"
	"strconv"

	lib "./lib"
)

func main() {

	// 呼叫main.go的function
	localHello()

	// 預設變數
	var a int = 1
	var b int32 = 2
	var c int64 = 3
	var d string = "999"
	var e float32 = 88.8
	var f float64 = 99.9
	var x string = "I Love Golang"

	// a + b
	fmt.Println(a + int(b))

	// a + b + c
	fmt.Println(a + int(b) + int(c))

	// f /c
	mod := f / float64(e)
	fmt.Printf("%f\n", mod)

	// a + d
	dToStr, _ := strconv.Atoi(d)
	fmt.Println(a + dToStr)

	// x + a
	aToStr := strconv.Itoa(a)
	fmt.Println(x + aToStr)

	// 呼叫package的function
	lib.Hello()
}

// 在main.go 裡面的function
func localHello() {
	fmt.Println("local hello")
}
