const mod = int64(1000000007)

var (
	total        int64
	result       int64
	memo         map[*TreeNode]int64
	originalRoot *TreeNode
)

func getSum(root *TreeNode) int64 {
	if root == nil {
		return 0
	}
	if mem, ok := memo[root]; ok {
		return mem
	}
	res := int64(root.Val) + getSum(root.Right) + getSum(root.Left)
	memo[root] = res
	return res
}

func getMax(arr ...int64) int64 {
	res := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] > res {
			res = arr[i]
		}
	}
	return res
}

func traverse(root *TreeNode) {
	if root == originalRoot {
		traverse(root.Left)
		traverse(root.Right)
		return
	}

	if root == nil {
		return
	}
	a := getSum(root)
	b := (total - a) * a
	result = getMax(result, a, b)
	traverse(root.Left)
	traverse(root.Right)
}

func maxProduct(root *TreeNode) int {
	originalRoot = root
	memo = make(map[*TreeNode]int64)
	total = getSum(root)
	result = 0
	traverse(root)
	return int(result % mod)
}
