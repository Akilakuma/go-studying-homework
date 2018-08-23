package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

func main() {

	testStart := time.Now()
	var waitgroup sync.WaitGroup
	var timeArr [100]float64

	for i := 0; i < 100; i++ {
		index := i
		waitgroup.Add(1)
		go func() {
			start := time.Now()
			execAPIRequest()
			// response := execAPIRequest()
			// log.Println(string(response))
			end := time.Since(start)

			timeArr[index] = end.Seconds()
			waitgroup.Done()
		}()
	}
	waitgroup.Wait()
	testEnd := time.Since(testStart)
	fmt.Println("花費時間:")
	fmt.Println(timeArr)

	var total float64
	for i := 0; i < 100; i++ {
		total += timeArr[i]
	}
	var max float64 = timeArr[0]
	var min float64 = timeArr[0]
	for _, value := range timeArr {
		if max < value {
			max = value
		}
		if min > value {
			min = value
		}
	}

	averageTime := total / 100
	fmt.Println("最大花費時間: " + strconv.FormatFloat(max, 'f', 4, 64))
	fmt.Println("最小花費時間: " + strconv.FormatFloat(min, 'f', 4, 64))
	fmt.Println("壓測總數花費時間: " + strconv.FormatFloat(testEnd.Seconds(), 'f', 4, 64))
	fmt.Println("平均時間: " + strconv.FormatFloat(averageTime, 'f', 4, 64))
}
