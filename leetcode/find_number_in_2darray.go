package leetcode

// 在一个 n * m 的二维数组中，每一行都按照从左到右递增的顺序排序，每一列都按照从上到下递增的顺序排序。
// 请完成一个函数，输入这样的一个二维数组和一个整数，判断数组中是否含有该整数。
// 思路：充分利用数组有序的特点，从二维数组的右上角开始查找，大于目标值往左移，小于则自动到下一行。
// findNumberIn2DArray ...
func findNumberIn2DArray(matrix [][]int, target int) bool {

	if len(matrix) ==0 {
		return false
	}

	row := len(matrix)
	column := len(matrix[0])
	currentRow := 0
	currentColumns := column -1
	for currentRow<row && currentColumns >=0 {
		temp := matrix[currentRow][currentColumns]
		if temp == target {
			return true
		}else if temp > target {
			currentColumns--
		}else {
			currentRow++
		}
	}
	return false
}
