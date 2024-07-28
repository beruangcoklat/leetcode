/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

 var path strings.Builder

 func getDirections(root *TreeNode, startValue int, destValue int) string {
	 path.Reset()
	 findPath(root, startValue)
	 startPath := reverse(path.String())
 
	 path.Reset()
	 findPath(root, destValue)
	 destPath := reverse(path.String())
 
	 idx := 0
 
	 for idx < len(startPath) && idx < len(destPath) && startPath[idx] == destPath[idx] {
		 idx++
	 }
 
	 startPath = startPath[idx:]
	 destPath = destPath[idx:]
 
	 result := ""
	 for i := 0; i < len(startPath); i++ {
		 result += "U"
	 }
 
	 return result + destPath
 }
 
 func findPath(root *TreeNode, value int) bool {
	 if root == nil {
		 return false
	 }
 
	 if root.Val == value {
		 return true
	 }
 
	 if findPath(root.Left, value) {
		 path.WriteString("L")
		 return true
	 }
 
	 if findPath(root.Right, value) {
		 path.WriteString("R")
		 return true
	 }
 
	 return false
 }
 
 func reverse(str string) string {
	 var result strings.Builder
	 for i := len(str) - 1; i >= 0; i-- {
		 result.WriteByte(str[i])
	 }
	 return result.String()
 }
 