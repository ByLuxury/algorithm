package leetcode

// TreeNode
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// mergeTrees 合并2个二叉树...
// 合并的规则是如果两个节点重叠，那么将他们的值相加作为节点合并后的新值
// 否则不为 NULL 的节点将直接作为新二叉树的节点。
func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {

	if t1 == nil {
		return t2
	}

	if t2 == nil {
		return t1
	}

	merge := &TreeNode{
		Val: t1.Val + t2.Val,
	}
	merge.Left = mergeTrees(t1.Left, t2.Left)
	merge.Right = mergeTrees(t1.Right, t2.Right)

	return merge
}

func MergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {

	return mergeTrees(t1, t2)
}
