package main

import "fmt"

const LEN = 10000000

func main() {
	var a [LEN]int
	//var a [LEN]int=[LEN]int{1,2,3,4,5,6,7,8,9,7}
	for i := 0; i < len(a); i++ {
		a[i] = i + 1
	}
	a[len(a) - 1] = a[6]
	result := search(a)
	fmt.Println(result)
	//fmt.Println(a)

}

/*func search(data [LEN]int64) int64 {
	var result int64
	//for x := 0; x < len(data); x++ {
	//	result ^= data[x]
	//}
	//for x := 1; x < len(data); x++ {
	//	result ^= x
	//}
	for x := 0; x < LEN; x++ {
		for y := x + 1; y < LEN; y++ {
			if data[x] == data[y] {
				result = data[x]
				return result
			}
		}
	}
	return -1
}*/


func search(data [LEN]int) int {
	var result int
	for x := 0; x < len(data); x++ {
		result ^= data[x]
	}
	for x := 1; x < len(data); x++ {
		result ^= x
	}
	return result
}
