package main

import (
	"fmt"
	"math"
)

// 实现一个 atoi 函数，使其能将字符串转换成整数。
// 如果第一个非空字符为正或者负号时，则将该符号与之后面尽可能多的连续数字字符组合起来，形成一个有符号整数。
// 假如第一个非空字符是数字，则直接将其与之后连续的数字字符组合起来，形成一个整数。
// 该字符串在有效的整数部分之后也可能会存在多余的字符，那么这些字符可以被忽略，它们对函数不应该造成影响。
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/string-to-integer-atoi
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

// MyAtoiCase ...
type MyAtoiCase struct {
	Input    string
	Expected int
}


// MyAtoi ...
func (myCase *MyAtoiCase) MyAtoi(str string) int {

	if str == "" || rune(str[0]) == ' ' {

		return 0
	}
	valid := make([]int32, 0)
	for _, v := range str {
		if !atoiIsRange(v) && (v != '+' && v != '-') {
			break
		}
		valid = append(valid, v)
	}

	var result int
	switch rune(str[0]) {

	case '+':
		result = calcNumber(valid[1:])
	case '-':
		result = calcNumber(valid[1:]) * -1
	default:
		result = calcNumber(valid)

	}
	switch true {

	case result > math.MaxInt32:
		return math.MaxInt32
	case result < math.MinInt32:
		return math.MinInt32
	default:
		return result
	}

}

// atoiIsRange ...
func atoiIsRange(x int32) bool {
	numbers := []int32{48, 57} //ASCII 码范围
	return x >= numbers[0] && x <= numbers[1]
}

func calcNumber(valid []int32) int {
	var result int
	for i := 0; i < len(valid); i++ {
		result += (int(valid[i]) - 48) * int(math.Pow10(len(valid)-1-i))
	}

	return result
}


func main() {

	testCase := []MyAtoiCase{
		{
			Input:    "1234",
			Expected: 1234,
		},

		{
			Input:    "1234def3",
			Expected: 1234,
		},

		{
			Input:    "-1234def3",
			Expected: -1234,
		},

		{
			Input:    "2147483648",
			Expected: 2147483647,
		},
		{
			Input:    "-2147483649Abc",
			Expected: -2147483648,
		},
	}

	for _, c := range testCase {
		actual := c.MyAtoi(c.Input)
		if actual != c.Expected {
			fmt.Printf("test failed:,expected:%d,got:%d\n", c.Expected, actual)
		}
	}

}
