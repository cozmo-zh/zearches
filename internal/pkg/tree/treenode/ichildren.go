// Package treenode .
package treenode

type IChildren interface {
	Divide(parent *TreeNode, depth int)
	ChildrenCount() int
	GetChild(index int) *TreeNode
	Clear()
}
