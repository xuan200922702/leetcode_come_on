package main


type TreeNode struct {
	   Val int
	   Left *TreeNode
	   Right *TreeNode
	 }
//深度优先
func levelOrder(root *TreeNode) [][]int {
	var res [][]int
	var _dfs func(root *TreeNode,level int)
	_dfs  = func(root *TreeNode, level int) {
		if root  ==  nil{
			return
		}
		if level ==len(res){
			res=append(res,[]int{})
		}
		res[level]  =  append(res[level],root.Val)
		_dfs(root.Left,level+1)
		_dfs(root.Right,level+1)
	}

	_dfs(root,0)
	return  res
}




