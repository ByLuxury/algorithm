package leetcode

// 题目描述
// 有一个数组(无序)：
// 找出数组中是否有两个数对（a,b）和（c,d）,使得a+b = c+d,其中，a,b,c和d是不同的元素。

func calcSum(arr []int) [][]int {

	res := make([][]int, 0)
	m := make(map[int][]int)

	for i := 0; i < len(arr); i++ {

		for j := i + 1; j < len(arr); j++ {
			sumAB := arr[i] + arr[j]
			if pair, ok := m[sumAB]; !ok {
				m[sumAB] = []int{arr[i], arr[j]}
			} else {
				res = append(res,append([]int{arr[i],arr[j]},pair...))
				return res
			}
		}

	}

	return res
}
