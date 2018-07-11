package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	// 加總1~20階乘用array
	calculateFactorial20withArray()

	// 加總1~20階乘用slice
	calculateFactorial20withSlice()
	// 加總1~100階乘
	calculateFactorial100()

}

func calculateFactorial20withArray() {

	start := time.Now()
	var n = 20
	number := make([]int, 21)
	number[0] = 1
	total := 0

	for i := 1; i <= n; i++ {
		mNow := i * number[i-1]
		number[i] = mNow

		total += mNow

	}

	elapsed := time.Since(start)

	fmt.Printf("1階乘相加到20階乘 - array")
	fmt.Println("\n 1! + 2! + ... + 20! ，總數:" + strconv.Itoa(total))
	fmt.Printf("耗時(奈秒)")
	fmt.Println(elapsed.Nanoseconds())
	fmt.Println("")
}

func calculateFactorial20withSlice() {

	start := time.Now()
	var n = 20
	number := []int{1}
	total := 0

	for i := 1; i <= n; i++ {
		mNow := i * number[i-1]
		// fmt.Println(strconv.Itoa(mNow))
		number = append(number, mNow)

		total += mNow

	}

	elapsed := time.Since(start)
	fmt.Printf("1階乘相加到20階乘 - slice")
	fmt.Println("\n 1! + 2! + ... + 20! ，總數:" + strconv.Itoa(total))
	fmt.Printf("耗時(奈秒)")
	fmt.Println(elapsed.Nanoseconds())
	fmt.Println("")
}

func calculateFactorial100() {

	start := time.Now()

	var n float64 = 100
	var i float64 = 1
	number := []float64{1}
	var total float64 = 0

	for i = 1; i <= n; i++ {
		mNow := i * number[int(i)-1]
		// fmt.Println(mNow)

		number = append(number, mNow)

		total += mNow

	}
	elapsed := time.Since(start)

	fmt.Printf("1階乘相加到100階乘 - slice")
	fmt.Println("\n 1! + 2! + ... + 100! ，總數:" + strconv.FormatFloat(total, 'g', 1, 64))
	fmt.Printf("耗時(奈秒)")
	fmt.Println(elapsed.Nanoseconds())
	fmt.Println("")
}
