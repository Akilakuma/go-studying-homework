package main

import (
	"fmt"
	"time"
)

func materialWoker(materialChannel chan float64) {

	var productMade = 0
	// 每秒1公噸的紙漿產能
	for productMade <= 6 {
		fmt.Printf("\n %c[1;40;32m%s%c[0m\n\n", 0x1B, "紙漿廠：準備中", 0x1B)

		select {
		case materialChannel <- 1:
			fmt.Printf("\n %c[1;40;32m%s%c[0m\n\n", 0x1B, "紙漿廠：製成->紙漿1頓", 0x1B)
		default:
			fmt.Printf("\n %c[1;40;32m%s%c[0m\n\n", 0x1B, "紙漿廠：buffer已滿，暫停製造", 0x1B)
		}
		productMade++
		// 提供紙漿

		time.Sleep(1 * time.Second)
	}

}

func paperManufactureWorker1(materialChannel chan float64, paperChannel chan int) {
	// 每秒0.5公噸的紙漿產能

	var stock float64
	for {
		select {
		case num := <-materialChannel: // 可以拿的時候！！！！！

			fmt.Printf("\n %c[1;40;33m%s%c[0m\n\n", 0x1B, "造紙廠1：可取得紙漿,紙漿成紙準備中", 0x1B)
			stock += num
			for ; stock >= 0.5; stock -= 0.5 {
				paperChannel <- 5000
				fmt.Printf("\n %c[1;40;33m%s%c[0m\n\n", 0x1B, "造紙廠1：製成->紙張5000", 0x1B)
				time.Sleep(1 * time.Second)
			}

		}

	}
}
func paperManufactureWorker2(materialChannel chan float64, paperChannel chan int) {
	// 每秒0.3公噸的紙漿產能

	var stock float64
	for {
		select {
		case num := <-materialChannel: // 可以拿的時候！！！！！
			fmt.Printf("\n %c[1;40;36m%s%c[0m\n\n", 0x1B, "造紙廠2：可取得紙漿,紙漿成紙準備中", 0x1B)
			stock += num
			for ; stock >= 0.3; stock -= 0.3 {
				paperChannel <- 3000
				fmt.Printf("\n %c[1;40;36m%s%c[0m\n\n", 0x1B, "造紙廠2：製成->紙張3000", 0x1B)
				time.Sleep(1 * time.Second)
			}

		}

	}
}

func printWorker(paperChannel chan int, printCannel chan int) {
	// 每秒列印6000張紙產能

	var stock int
	for {
		select {
		case num := <-paperChannel: // 可以拿的時候！！！！！
			fmt.Printf("\n %c[1;40;35m%s%c[0m\n\n", 0x1B, "列印廠：可取得紙,列印準備中", 0x1B)
			stock += num
			for ; stock >= 6000; stock -= 6000 {
				fmt.Printf("\n %c[1;40;35m%s%d%c[0m\n", 0x1B, "列印廠：有紙:", stock, 0x1B)
				printCannel <- 6000
				fmt.Printf("\n %c[1;40;35m%s%c[0m\n\n", 0x1B, "列印廠：製成->列印6000", 0x1B)
				time.Sleep(1 * time.Second)
			}
		}
	}
}

func main() {

	var materialChan chan float64 = make(chan float64, 100)

	var paperChan chan int = make(chan int, 100)
	var printChan chan int = make(chan int, 100)

	go materialWoker(materialChan)
	go paperManufactureWorker1(materialChan, paperChan)
	go paperManufactureWorker2(materialChan, paperChan)
	go printWorker(paperChan, printChan)

	totalPrint := 0
	for totalPrint < 60000 {
		select {
		case num := <-printChan: // 可以拿的時候！！！！！
			totalPrint += num
			fmt.Printf("已完成列印總數: %d\n", totalPrint)
		}
	}
}
