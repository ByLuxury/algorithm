package leetcode


// BinarySearch ...
func BinarySearch(arr []int, low, high, dst int) (index int) {

	if low > high {
		return -1
	}

	mid := low + (high-low)/2

	if dst > arr[mid] { // 从后半部分找
		return BinarySearch(arr, mid+1, high, dst)
	} else if dst < arr[mid] {
		return BinarySearch(arr, low, mid-1, dst)
	}

	return mid
}
