package main

import (
	"fmt"
)

func main() {

	ch := make(chan int, 1)
	for {
		select { //在0～2中产生随机数
		case ch <- 0:
		case ch <- 1:
		case ch <- 2:
		case ch <- 3:
		}
		i := <-ch
		fmt.Println("\n 产生的随机数为：", i)
	}
}
