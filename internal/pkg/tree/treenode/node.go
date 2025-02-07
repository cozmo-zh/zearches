// Package treenode provides the implementation of a tree node used in spatial indexing.
package treenode

import (
	"container/list"
	"fmt"
	"github.com/cozmo-zh/zearches/consts"
	"github.com/cozmo-zh/zearches/pkg/bounds"
	"github.com/cozmo-zh/zearches/pkg/geo"
	"github.com/cozmo-zh/zearches/pkg/siface"
	"github.com/cozmo-zh/zearches/util"
)

// TreeNode is a node in the tree.
//
// Not thread-safe, only works in a single thread(goroutine).
type TreeNode struct {
	depth       int // Depth of the node in the tree.
	maxDepth    int // Maximum depth of the tree.
	capacity    int // Maximum number of entities the node can hold.
	index       int
	bound       bounds.Bound            // Spatial boundaries of the node.
	entityList  *list.List              // List of entities in the node.
	entityIndex map[int64]*list.Element // Map of entity IDs to their list elements.
	parent      *TreeNode               // Parent node.
	children    IDimensionNode
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
func NewTreeNode(dim consts.Dim, parent *TreeNode, bound bounds.Bound, index, depth, maxDepth, capacity int) (*TreeNode, error) {
	var children IDimensionNode
	switch dim {
	case consts.Dim2:
		children = NewD2()
	case consts.Dim3:
		children = NewD3()
	default:
		return nil, fmt.Errorf("unsupported dimension: %v", dim)
	}
	if maxDepth < 1 {
		return nil, fmt.Errorf("maxDepth should be greater than 0")
	}
	if capacity < 1 {
		return nil, fmt.Errorf("capacity should be greater than 0")
	}
	return &TreeNode{
		parent:      parent,
		depth:       depth,
		maxDepth:    maxDepth,
		capacity:    capacity,
		bound:       bound,
		entityList:  list.New(),
		entityIndex: make(map[int64]*list.Element),
		children:    children,
		index:       index,
	}, nil
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
		for i := 0; i < n.children.ChildrenCount(); i++ {
			if _n.children.GetChild(i).Add(spatial) {
				return true
			}
		}
		return false
	}
	if n.IsLeaf() {
		if n.entityList.Len() < n.capacity {
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
		for i := 0; i < n.children.ChildrenCount(); i++ {
			if n.children.GetChild(i).Remove(spatialId, merge...) {
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

func (n *TreeNode) ClearChildren() {
	n.children.Clear()
}

// DivideIf divides the node into 8 children if the number of entities exceeds the capacity.
// if the depth of the node exceeds the maximum depth, the node will not be divided.
//
// Returns:
// - true if the node was divided, false otherwise.
func (n *TreeNode) DivideIf() bool {
	if n.depth >= n.maxDepth-1 {
		// Maximum depth reached.
		return false
	}
	if n.entityList.Len() < n.capacity {
		return false
	}
	n.children.Divide(n, n.depth+1)
	// Move entities to children.
	for e := n.entityList.Front(); e != nil; e = e.Next() {
		spatial := e.Value.(siface.ISpatial)
		for i := 0; i < n.children.ChildrenCount(); i++ {
			child := n.children.GetChild(i)
			if child.Contains(spatial) {
				child.Add(spatial)
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
	return n.children.Contains(n, spatial)
}

// Intersects checks if the bound intersects with the node.
func (n *TreeNode) Intersects(bound bounds.Bound) bool {
	return n.children.Intersects(n, bound)
}

// IsLeaf checks if the node is a leaf node.
func (n *TreeNode) IsLeaf() bool {
	for i := 0; i < n.children.ChildrenCount(); i++ {
		if n.children.GetChild(i) != nil {
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
	for i := 0; i < n.parent.Children().ChildrenCount(); i++ {
		child := n.parent.Children().GetChild(i)
		if !child.IsLeaf() {
			return false
		}
		count += child.Size()
	}
	if count >= n.capacity {
		return false
	}
	// Merge the node with its children
	add := make([]siface.ISpatial, 0)
	for i := 0; i < n.parent.children.ChildrenCount(); i++ {
		child := n.parent.children.GetChild(i)
		for e := child.GetEntityList().Front(); e != nil; e = e.Next() {
			spatial := e.Value.(siface.ISpatial)
			add = append(add, spatial)
		}
		child.Clear()
	}
	n.parent.ClearChildren()
	for _, spatial := range add {
		n.parent.Add(spatial)
	}
	return true
}

// FindEntities finds entities within a radius of a center point.
func (n *TreeNode) FindEntities(center geo.Vec3Int, radius float32, filters ...func(entity siface.ISpatial) bool) []siface.ISpatial {
	// build a cube bound for the search
	// the cube is a bounding box of the sphere
	cMin := geo.NewVec3Int(center.X()-int32(radius), center.Y()-int32(radius), center.Z()-int32(radius))
	cMax := geo.NewVec3Int(center.X()+int32(radius), center.Y()+int32(radius), center.Z()+int32(radius))
	cBound := bounds.NewBound(cMin, cMax)
	return n.FindEntitiesInBound(cBound, radius, filters...)
}

// FindEntitiesInBound finds entities within a bound.
func (n *TreeNode) FindEntitiesInBound(bound bounds.Bound, radius float32, filters ...func(entity siface.ISpatial) bool) []siface.ISpatial {
	// search entities
	ret := make([]siface.ISpatial, 0)
	if n.Intersects(bound) {
		if n.IsLeaf() {
			for e := n.entityList.Front(); e != nil; e = e.Next() {
				spatial := e.Value.(siface.ISpatial)
				if util.WithinDistance3D(spatial.GetLocation().ToFloat32(), bound.Center.ToFloat32(), radius) {
					if len(filters) == 0 {
						ret = append(ret, spatial)
					} else {
						for _, filter := range filters {
							if filter(spatial) {
								ret = append(ret, spatial)
								break
							}
						}
					}
				}
			}
		} else {
			// check children
			for i := 0; i < n.children.ChildrenCount(); i++ {
				child := n.children.GetChild(i)
				if child == nil {
					continue
				}
				ret = append(ret, child.FindEntitiesInBound(bound, radius, filters...)...)
			}
		}
	}
	return ret
}

// Bound returns the spatial boundaries of the node.
func (n *TreeNode) Bound() bounds.Bound {
	return n.bound
}

// MaxDepth returns the maximum depth of the tree.
func (n *TreeNode) MaxDepth() int {
	return n.maxDepth
}

// Capacity returns the maximum number of entities the node can hold.
func (n *TreeNode) Capacity() int {
	return n.capacity
}

// Parent returns the parent node.
func (n *TreeNode) Parent() *TreeNode {
	return n.parent
}

// Size returns the number of entities in the node.
func (n *TreeNode) Size() int {
	return n.entityList.Len()
}

// Children returns the children of the node.
func (n *TreeNode) Children() IDimensionNode {
	return n.children
}

func (n *TreeNode) Depth() int {
	return n.depth
}

func (n *TreeNode) Index() int {
	return n.index
}

// Range .
func (n *TreeNode) Range(f func(n *TreeNode) bool) {
	if !f(n) {
		return
	}
	if !n.IsLeaf() {
		for i := 0; i < n.children.ChildrenCount(); i++ {
			child := n.children.GetChild(i)
			if child != nil {
				child.Range(f)
			}
		}
	}
}

// RangeEntities ranges the entities in the node.
func (n *TreeNode) RangeEntities(f func(entity siface.ISpatial) bool) {
	for e := n.entityList.Front(); e != nil; e = e.Next() {
		if !f(e.Value.(siface.ISpatial)) {
			return
		}
	}
}
