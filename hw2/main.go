package main

import (
	"fmt"
	"sort"
)

func main() {
	practice1()
	practice2()
}

func practice1() {
	// 先建立一個長度為 5 的陣列來儲存這些測驗成績
	var score [5]int

	// 接著將分數填入每個元素中
	score[0] = 88
	score[1] = 99
	score[2] = 76
	score[3] = 82
	score[4] = 91

	// 再來使用一個迴圈來計算成績的總和
	var total = 0
	for k, v := range score {
		total += v
		fmt.Printf("score[%d]=%d\n", k, v)
	}

	// 最後我們將成績的總和除以元素的數量，以取得平均值。
	average := total / len(score)

	fmt.Printf("平均分數是： %d \n\n", average)
}

func practice2() {
	// 寫個程式找出串列中最小的數字
	x := []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}
	fmt.Println("slice內容為： ", x)
	sort.Ints(x)

	fmt.Printf("最小的數字是： %d \n", x[0])
}
