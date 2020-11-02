package plusplus

// 由于输入是一个升序排列的有序数组。
// 因此任意选择一点，将其作为根节点，其左部分左节点，其右部分右节点即可。
// 因此我们很容易写出递归代码。
// 而题目要求是高度平衡的二叉搜索树，因此我们必须要取中点。
// 不难证明：由于是中点，因此左右两部分差不会大于 1，也就是说其形成的左右子树节点数最多相差 1，因此左右子树高度差的绝对值不超过 1
func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	mid := len(nums) / 2
	root := &TreeNode{Val: nums[mid]}
	root.Left = sortedArrayToBST(nums[:mid])
	root.Right = sortedArrayToBST(nums[mid+1:])
	return root
}
