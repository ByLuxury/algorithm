package leetcode

// 编写一个函数来查找字符串数组中的最长公共前缀。
//如果不存在公共前缀，返回空字符串 ""。

//示例 1:

//输入: ["flower","flow","flight"]
//输出: "fl"

//示例 2:
//输入: ["dog","racecar","car"]
//输出: ""
//解释: 输入不存在公共前缀。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/longest-common-prefix
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

// LCP ...
type LCP struct {
	Input    []string
	Expected string
}

// LongestCommonPrefix .. 采用分治法
func (l *LCP) LongestCommonPrefix(strs []string) string {

	if len(strs) == 0 {
		return ""
	}

	return lcp(strs, 0, len(strs)-1)

}

// lcp 分治递归计算
func lcp(strs []string, l, r int) string {

	if l == r { // 说明分治到了相同位置
		return strs[l]
	}
	// 递归计算左右两边的公共前缀
	mid := (l + r) / 2
	left := lcp(strs, l, mid)
	right := lcp(strs, mid+1, r)
	return commonPrefix(left, right) // 合并2个子串的公共前缀

}

// commonPrefix 计算公共前缀
func commonPrefix(left, right string) string {
	if left == "" || right == "" {
		return ""
	}

	length := len(left)
	if length > len(right) {
		length = len(right) // 肯定是较小的那个
	}

	for i := 0; i < length; i++ {
		if left[i] != right[i] {
			if i == 0 {
				return ""
			}
			return string(left[:i])
		}
	}

	if len(left) > len(right) {
		return right
	}

	return left
}
