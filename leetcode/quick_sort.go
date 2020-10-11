package leetcode

func sortArray(nums []int) []int {

	quickSort(nums, 0, len(nums)-1)
	return nums
}

func quickSort(arr []int, l, r int) {

	left, right := l, r

	if left <= right {
		pivot := arr[left]
		for left != right {
			for left < right && arr[right] >= pivot {
				right-- //从左往右扫描，直到找到一个第一个比基准元素小的值
			}
			arr[left] = arr[right]

			for left < right && arr[left] <= pivot { //从左往右扫描，找到第一个比基准元素大的值
				left++
			}
			arr[right] = arr[left]
		}
		arr[right] = pivot // 基准元素归位

		quickSort(arr, l, left-1)
		quickSort(arr, right+1, r)
	}

}
