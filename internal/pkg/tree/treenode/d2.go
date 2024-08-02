// Package treenode .
package treenode

import (
	"github.com/cozmo-zh/zearches/consts"
	"github.com/cozmo-zh/zearches/pkg/bounds"
	"github.com/cozmo-zh/zearches/pkg/geo"
)

const (
	childrenCountD2 = 4
)

// D2 .
type D2 struct {
	children [childrenCountD2]*TreeNode
}

// NewD2 .
func NewD2() *D2 {
	return &D2{}
}

// Divide the node into 4 children and move entities to children.
func (d *D2) Divide(parent *TreeNode, depth int) {
	/*
	*   |1|2|
	*   --.--
	*   |0|3|
	 */
	min0 := geo.NewVec3Int(parent.Bound().Min.X(), 0, parent.Bound().Min.Z())
	max0 := geo.NewVec3Int(parent.Bound().Center.X(), 0, parent.Bound().Center.Z())
	bound0 := bounds.NewBound(min0, max0)

	min1 := geo.NewVec3Int(parent.Bound().Min.X(), 0, parent.Bound().Center.Z())
	max1 := geo.NewVec3Int(parent.Bound().Center.X(), 0, parent.Bound().Max.Z())
	bound1 := bounds.NewBound(min1, max1)

	min2 := geo.NewVec3Int(parent.Bound().Center.X(), 0, parent.Bound().Center.Z())
	max2 := geo.NewVec3Int(parent.Bound().Max.X(), 0, parent.Bound().Max.Z())
	bound2 := bounds.NewBound(min2, max2)

	min3 := geo.NewVec3Int(parent.Bound().Center.X(), 0, parent.Bound().Min.Z())
	max3 := geo.NewVec3Int(parent.Bound().Max.X(), 0, parent.Bound().Center.Z())
	bound3 := bounds.NewBound(min3, max3)

	// Create children.
	maxDepth := parent.MaxDepth()
	capacity := parent.Capacity()

	d.children[0], _ = NewTreeNode(consts.Dim2, parent, bound0, depth, maxDepth, capacity)
	d.children[1], _ = NewTreeNode(consts.Dim2, parent, bound1, depth, maxDepth, capacity)
	d.children[2], _ = NewTreeNode(consts.Dim2, parent, bound2, depth, maxDepth, capacity)
	d.children[3], _ = NewTreeNode(consts.Dim2, parent, bound3, depth, maxDepth, capacity)

}

// ChildrenCount returns the number of children.
func (d *D2) ChildrenCount() int {
	return childrenCountD2
}

// GetChild returns the child at the specified index.
func (d *D2) GetChild(index int) *TreeNode {
	return d.children[index]
}

// Clear removes all children.
func (d *D2) Clear() {
	d.children = [childrenCountD2]*TreeNode{}
}
