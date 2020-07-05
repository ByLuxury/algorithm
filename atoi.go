package main

import (
	"fmt"
	"math"
	"strings"
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

func (myCase *MyAtoiCase) MyAtoi(str string) int {

	if str == "" {

		return 0
	}
	flag := 1
	result := 0

	// 找到第一个除空和+-的有效数字，第一个是无效的字符串直接返回0
	str = strings.TrimSpace(str) // 去掉空格，筛选出有效的数

Loop:
	for i, v := range str {
		switch true {
		case v == '+' && i == 0:
		case v == '-' && i == 0:
			flag = -1
		case atoiIsRange(v):
			result = result*10 + int(v-'0')
			if flag == -1 && result*flag < math.MinInt32 {
				return math.MinInt32
			}

			if flag == 1 && result*flag > math.MaxInt32 {
				return math.MaxInt32
			}

		default:
			break Loop

		}
	}

	return result * flag

}

// atoiIsRange ...
func atoiIsRange(x int32) bool {
	numbers := []int32{48, 57} //ASCII 码范围
	return x >= numbers[0] && x <= numbers[1]
}

func main() {

	testCase := []MyAtoiCase{
		{
			Input:    "  -44",
			Expected: -44,
		},
		{
			Input:    "H124",
			Expected: 0,
		},
		{
			Input:    "102 this code",
			Expected: 102,
		},
		{
			Input:    "+-12",
			Expected: 0,
		},
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
			Input:    "-91283472332",
			Expected: -2147483648,
		},
		{
			Input:    "18446744073709551617",
			Expected: 2147483647,
		},
	}

	for _, c := range testCase {
		actual := c.MyAtoi(c.Input)
		if actual != c.Expected {
			fmt.Printf("test failed:,expected:%d,got:%d\n", c.Expected, actual)
		}
	}

}
