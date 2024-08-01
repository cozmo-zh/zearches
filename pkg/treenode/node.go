// Package treenode provides the implementation of a tree node used in spatial indexing.
package treenode

import (
	"container/list"
	"github.com/cozmo-zh/zearches/pkg/bound"
	"github.com/cozmo-zh/zearches/pkg/consts"
	"github.com/cozmo-zh/zearches/pkg/geo"
	"github.com/cozmo-zh/zearches/pkg/siface"
)

// TreeNode is a node in the tree.
//
// Not thread-safe, only works in a single thread(goroutine).
type TreeNode struct {
	depth       int                     // Depth of the node in the tree.
	maxDepth    *int                    // Maximum depth of the tree.
	capacity    *int                    // Maximum number of entities the node can hold.
	bound       bound.Bound             // Spatial boundaries of the node.
	children    [8]*TreeNode            // Child nodes.
	entityList  *list.List              // List of entities in the node.
	entityIndex map[int64]*list.Element // Map of entity IDs to their list elements.
	parent      *TreeNode               // Parent node.
}

// NewTreeNode creates a new tree node.
//
// Parameters:
// - dim: The dimension of the tree (e.g., 2D, 3D).
// - bound: The spatial boundaries of the node.
// - capacity: The maximum number of entities that the node can hold.
//
// Returns:
// - A pointer to the newly created TreeNode.
func NewTreeNode(dim consts.Dim, parent *TreeNode, bound bound.Bound, depth int, maxDepth, capacity *int) *TreeNode {
	return &TreeNode{
		parent:      parent,
		depth:       depth,
		maxDepth:    maxDepth,
		capacity:    capacity,
		bound:       bound,
		entityList:  list.New(),
		entityIndex: make(map[int64]*list.Element),
	}
}

// Add adds a spatial entity to the node.
//
// Parameters:
// - spatial: The spatial entity to add.
//
// Returns:
// - true if the entity was added successfully, false otherwise.
func (n *TreeNode) Add(spatial siface.ISpatial) bool {
	if !n.Contains(spatial) {
		return false
	}
	add := func(_n *TreeNode, spatial siface.ISpatial) {
		e := _n.entityList.PushBack(spatial)
		_n.entityIndex[spatial.GetID()] = e
	}
	add2Children := func(_n *TreeNode, spatial siface.ISpatial) bool {
		for i := 0; i < 8; i++ {
			if _n.children[i].Add(spatial) {
				return true
			}
		}
		return false
	}
	if n.IsLeaf() {
		if n.entityList.Len() < *n.capacity {
			add(n, spatial)
			return true
		} else {
			if n.DivideIf() {
				add2Children(n, spatial)
			} else {
				add(n, spatial)
				return true
			}
		}
	} else {
		add2Children(n, spatial)
	}
	return true
}

// Remove removes a spatial entity from the node by its ID.
//
// Parameters:
// - spatialId: The ID of the spatial entity to remove.
// - merge: Whether to merge the node with its children after removing the entity.
//
// Returns:
// - true if the entity was removed successfully, false otherwise.
func (n *TreeNode) Remove(spatialId int64, merge ...bool) bool {
	if n.IsLeaf() {
		if e, ok := n.entityIndex[spatialId]; ok {
			delete(n.entityIndex, spatialId)
			n.entityList.Remove(e)
			if len(merge) > 0 && merge[0] {
				n.MergeIf()
			}
			return true
		}
		return false
	} else {
		for i := 0; i < 8; i++ {
			if n.children[i].Remove(spatialId, merge...) {
				return true
			}
		}
		return false
	}
}

// GetEntityList returns the list of entities in the node.
//
// Returns:
// - A pointer to the list of entities.
func (n *TreeNode) GetEntityList() *list.List {
	return n.entityList
}

// Clear removes all entities from the node.
func (n *TreeNode) Clear() {
	clear(n.entityIndex)
	n.entityList.Init()
}

// DivideIf divides the node into 8 children if the number of entities exceeds the capacity.
// if the depth of the node exceeds the maximum depth, the node will not be divided.
//
// Returns:
// - true if the node was divided, false otherwise.
func (n *TreeNode) DivideIf() bool {
	if n.depth >= *n.maxDepth {
		// Maximum depth reached.
		return false
	}
	if n.entityList.Len() < *n.capacity {
		return false
	}
	// Divide the node into 8 children and move entities to children.
	// Divide first.
	min0 := geo.NewVec3Int(n.bound.Min.X(), n.bound.Min.Y(), n.bound.Min.Z())
	max0 := geo.NewVec3Int(n.bound.Center.X(), n.bound.Center.Y(), n.bound.Center.Z())
	bound0 := bound.NewBound(min0, max0)

	min1 := geo.NewVec3Int(n.bound.Min.X(), n.bound.Min.Y(), n.bound.Center.Z())
	max1 := geo.NewVec3Int(n.bound.Center.X(), n.bound.Center.Y(), n.bound.Max.Z())
	bound1 := bound.NewBound(min1, max1)

	min2 := geo.NewVec3Int(n.bound.Min.X(), n.bound.Center.Y(), n.bound.Min.Z())
	max2 := geo.NewVec3Int(n.bound.Center.X(), n.bound.Max.Y(), n.bound.Center.Z())
	bound2 := bound.NewBound(min2, max2)

	min3 := geo.NewVec3Int(n.bound.Min.X(), n.bound.Center.Y(), n.bound.Center.Z())
	max3 := geo.NewVec3Int(n.bound.Center.X(), n.bound.Max.Y(), n.bound.Max.Z())
	bound3 := bound.NewBound(min3, max3)

	min4 := geo.NewVec3Int(n.bound.Center.X(), n.bound.Min.Y(), n.bound.Min.Z())
	max4 := geo.NewVec3Int(n.bound.Max.X(), n.bound.Center.Y(), n.bound.Center.Z())
	bound4 := bound.NewBound(min4, max4)

	min5 := geo.NewVec3Int(n.bound.Center.X(), n.bound.Min.Y(), n.bound.Center.Z())
	max5 := geo.NewVec3Int(n.bound.Max.X(), n.bound.Center.Y(), n.bound.Max.Z())
	bound5 := bound.NewBound(min5, max5)

	min6 := geo.NewVec3Int(n.bound.Center.X(), n.bound.Center.Y(), n.bound.Min.Z())
	max6 := geo.NewVec3Int(n.bound.Max.X(), n.bound.Max.Y(), n.bound.Center.Z())
	bound6 := bound.NewBound(min6, max6)

	min7 := geo.NewVec3Int(n.bound.Center.X(), n.bound.Center.Y(), n.bound.Center.Z())
	max7 := geo.NewVec3Int(n.bound.Max.X(), n.bound.Max.Y(), n.bound.Max.Z())
	bound7 := bound.NewBound(min7, max7)

	// Increase the depth.
	depth := n.depth + 1
	// Create children.
	n.children[0] = NewTreeNode(consts.Dim3, n, bound0, depth, n.maxDepth, n.capacity)
	n.children[1] = NewTreeNode(consts.Dim3, n, bound1, depth, n.maxDepth, n.capacity)
	n.children[2] = NewTreeNode(consts.Dim3, n, bound2, depth, n.maxDepth, n.capacity)
	n.children[3] = NewTreeNode(consts.Dim3, n, bound3, depth, n.maxDepth, n.capacity)
	n.children[4] = NewTreeNode(consts.Dim3, n, bound4, depth, n.maxDepth, n.capacity)
	n.children[5] = NewTreeNode(consts.Dim3, n, bound5, depth, n.maxDepth, n.capacity)
	n.children[6] = NewTreeNode(consts.Dim3, n, bound6, depth, n.maxDepth, n.capacity)
	n.children[7] = NewTreeNode(consts.Dim3, n, bound7, depth, n.maxDepth, n.capacity)

	// Move entities to children.
	for e := n.entityList.Front(); e != nil; e = e.Next() {
		spatial := e.Value.(siface.ISpatial)
		for i := 0; i < 8; i++ {
			if n.children[i].Contains(spatial) {
				n.children[i].Add(spatial)
				break
			}
		}
	}
	// Clear the entity list.
	n.Clear()
	return true
}

// Contains checks if the spatial entity is within the bounds of the node.
//
// Parameters:
// - spatial: The spatial entity to check.
//
// Returns:
// - true if the entity is within the bounds, false otherwise.
func (n *TreeNode) Contains(spatial siface.ISpatial) bool {
	if n.bound.Min.X() <= spatial.GetLocation().X() && spatial.GetLocation().X() <= n.bound.Max.X() &&
		n.bound.Min.Y() <= spatial.GetLocation().Y() && spatial.GetLocation().Y() <= n.bound.Max.Y() &&
		n.bound.Min.Z() <= spatial.GetLocation().Z() && spatial.GetLocation().Z() <= n.bound.Max.Z() {
		return true
	}
	return false
}

// IsLeaf checks if the node is a leaf node.
func (n *TreeNode) IsLeaf() bool {
	for _, child := range n.children {
		if child != nil {
			return false
		}
	}
	return true
}

// MergeIf merges the node with its children if the number of entities in the node is less than the capacity.
// not suggested to use this method, it's not efficient.
func (n *TreeNode) MergeIf() bool {
	if !n.IsLeaf() || n.parent == nil {
		return false
	}
	// check other siblings
	count := 0
	for _, sibling := range n.parent.children {
		if sibling != nil {
			count += sibling.entityList.Len()
		}
	}
	if count >= *n.capacity {
		return false
	}
	// Merge the node with its children
	for _, sibling := range n.parent.children {
		if sibling != nil {
			for e := sibling.entityList.Front(); e != nil; e = e.Next() {
				spatial := e.Value.(siface.ISpatial)
				n.parent.Add(spatial)
			}
			sibling.Clear()
		}
	}
	n.parent.children = [8]*TreeNode{}
	return true
}

func (n *TreeNode) Bound() bound.Bound {
	return n.bound
}
