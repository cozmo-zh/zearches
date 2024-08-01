// Package treenode .
package treenode

import (
	"github.com/cozmo-zh/zearches/pkg/bounds"
	"github.com/cozmo-zh/zearches/pkg/consts"
	"github.com/cozmo-zh/zearches/pkg/geo"
)

const (
	childrenCountD3 = 8
)

// D3 .
type D3 struct {
	children [childrenCountD3]INode
}

// NewD3 .
func NewD3() *D3 {
	return &D3{}
}

// Divide the node into 8 children and move entities to children.
func (d *D3) Divide(parent INode, depth int) {
	/*   3____7
	*  2/___6/|
	*  | 1__|_5
	*  0/___4/
	 */
	min0 := geo.NewVec3Int(parent.Bound().Min.X(), parent.Bound().Min.Y(), parent.Bound().Min.Z())
	max0 := geo.NewVec3Int(parent.Bound().Center.X(), parent.Bound().Center.Y(), parent.Bound().Center.Z())
	bound0 := bounds.NewBound(min0, max0)

	min1 := geo.NewVec3Int(parent.Bound().Min.X(), parent.Bound().Min.Y(), parent.Bound().Center.Z())
	max1 := geo.NewVec3Int(parent.Bound().Center.X(), parent.Bound().Center.Y(), parent.Bound().Max.Z())
	bound1 := bounds.NewBound(min1, max1)

	min2 := geo.NewVec3Int(parent.Bound().Min.X(), parent.Bound().Center.Y(), parent.Bound().Min.Z())
	max2 := geo.NewVec3Int(parent.Bound().Center.X(), parent.Bound().Max.Y(), parent.Bound().Center.Z())
	bound2 := bounds.NewBound(min2, max2)

	min3 := geo.NewVec3Int(parent.Bound().Min.X(), parent.Bound().Center.Y(), parent.Bound().Center.Z())
	max3 := geo.NewVec3Int(parent.Bound().Center.X(), parent.Bound().Max.Y(), parent.Bound().Max.Z())
	bound3 := bounds.NewBound(min3, max3)

	min4 := geo.NewVec3Int(parent.Bound().Center.X(), parent.Bound().Min.Y(), parent.Bound().Min.Z())
	max4 := geo.NewVec3Int(parent.Bound().Max.X(), parent.Bound().Center.Y(), parent.Bound().Center.Z())
	bound4 := bounds.NewBound(min4, max4)

	min5 := geo.NewVec3Int(parent.Bound().Center.X(), parent.Bound().Min.Y(), parent.Bound().Center.Z())
	max5 := geo.NewVec3Int(parent.Bound().Max.X(), parent.Bound().Center.Y(), parent.Bound().Max.Z())
	bound5 := bounds.NewBound(min5, max5)

	min6 := geo.NewVec3Int(parent.Bound().Center.X(), parent.Bound().Center.Y(), parent.Bound().Min.Z())
	max6 := geo.NewVec3Int(parent.Bound().Max.X(), parent.Bound().Max.Y(), parent.Bound().Center.Z())
	bound6 := bounds.NewBound(min6, max6)

	min7 := geo.NewVec3Int(parent.Bound().Center.X(), parent.Bound().Center.Y(), parent.Bound().Center.Z())
	max7 := geo.NewVec3Int(parent.Bound().Max.X(), parent.Bound().Max.Y(), parent.Bound().Max.Z())
	bound7 := bounds.NewBound(min7, max7)

	// Create children.
	maxDepth := parent.MaxDepth()
	capacity := parent.Capacity()

	d.children[0], _ = NewTreeNode(consts.Dim3, parent, bound0, depth, maxDepth, capacity)
	d.children[1], _ = NewTreeNode(consts.Dim3, parent, bound1, depth, maxDepth, capacity)
	d.children[2], _ = NewTreeNode(consts.Dim3, parent, bound2, depth, maxDepth, capacity)
	d.children[3], _ = NewTreeNode(consts.Dim3, parent, bound3, depth, maxDepth, capacity)
	d.children[4], _ = NewTreeNode(consts.Dim3, parent, bound4, depth, maxDepth, capacity)
	d.children[5], _ = NewTreeNode(consts.Dim3, parent, bound5, depth, maxDepth, capacity)
	d.children[6], _ = NewTreeNode(consts.Dim3, parent, bound6, depth, maxDepth, capacity)
	d.children[7], _ = NewTreeNode(consts.Dim3, parent, bound7, depth, maxDepth, capacity)
}

// ChildrenCount return the number of children.
func (d *D3) ChildrenCount() int {
	return childrenCountD3
}

// GetChild return the child at the given index.
func (d *D3) GetChild(index int) INode {
	return d.children[index]
}

// Clear the children.
func (d *D3) Clear() {
	d.children = [childrenCountD3]INode{}
}
