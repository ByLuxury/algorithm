package other

const LEN = 10000000


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
