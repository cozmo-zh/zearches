// Package treenode .
package treenode

import (
	"github.com/cozmo-zh/zearches/pkg/bounds"
	"github.com/cozmo-zh/zearches/pkg/siface"
)

type IDimensionNode interface {
	Divide(parent *TreeNode, depth int)
	ChildrenCount() int
	GetChild(index int) *TreeNode
	Clear()
	Contains(n *TreeNode, spatial siface.ISpatial) bool
	Intersects(n *TreeNode, bound bounds.Bound) bool
}
