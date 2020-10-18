package leetcode

// canPlaceFlowers ...
func canPlaceFlowers(flowerbed []int, n int) bool {

	// 贪心算法
	if len(flowerbed) == 0 {
		return false
	}

	var count int
	// 1.当在第一个或最后一个时，只需判断目前值以及他下一个或上一个节点的值是否为0
	// 2.判断本身、上一个、下一个
	for i:=0;i<len(flowerbed);i++ {
		if flowerbed[i]==0 && ((i==0 || flowerbed[i-1]==0) &&
			(i==len(flowerbed)-1 || flowerbed[i+1]==0)) {
			flowerbed[i] = 1
			count++
		}
		if count >= n {
			return true
		}

	}
	return false
}
