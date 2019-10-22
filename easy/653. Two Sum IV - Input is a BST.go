package easy

// TreeNode is
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// FindTarget is
func FindTarget(root *TreeNode, k int) bool {
	var map1 = make(map[int]bool)
	return search(root, k, map1)
}

func search(root *TreeNode, k int, map1 map[int]bool) bool {
	if root == nil {
		return false
	}
	if _, ok := map1[k-root.Val]; ok {
		return ok
	}
	map1[root.Val] = true
	return search(root.Left, k, map1) || search(root.Right, k, map1)
}
