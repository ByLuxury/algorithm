package leetcode

// input:[1,2,3]
// Permutation 数组全排列
func Permutation(arr []int) [][]int {
	res := make([][]int, 0)
	l := len(arr)
	if len(arr) == 0 {
		return res
	}

	var arrange func(pos int)
	arrange = func(pos int) {
		if pos == l-1 {
			x := make([]int, l)
			copy(x, arr)
			res = append(res, x)
			return
		}

		for i := pos; i < l; i++ {
			swap(arr, i, pos)
			arrange(pos + 1)
			swap(arr, i, pos)
		}
	}
	arrange(0)
	return res

}

func swap(arr []int, i, j int) {
	if i != j && i < len(arr) && j < len(arr) {
		arr[i] = arr[i] ^ arr[j]
		arr[j] = arr[i] ^ arr[j]
		arr[i] = arr[i] ^ arr[j]
	}

}

