// Package quadtree provides an implementation of a quadtree data structure for spatial partitioning.
package quadtree

import (
	"github.com/cozmo-zh/zearches/consts"
	"github.com/cozmo-zh/zearches/internal/pkg/tree"
	"github.com/cozmo-zh/zearches/internal/pkg/tree/treenode"
	"github.com/cozmo-zh/zearches/pkg/bounds"
	"github.com/cozmo-zh/zearches/pkg/geo"
	"github.com/cozmo-zh/zearches/pkg/siface"
)

// QuadTree represents a quadtree data structure.
type QuadTree struct {
	root    *treenode.TreeNode            // The root node of the quadtree.
	scale   func(v []float32) geo.Vec3Int // Function to scale float32 slice to geo.Vec3Int.
	mergeIf bool                          // Flag to determine if nodes should be merged when removing an entity.
}

// NewQuadtree creates a new Octree.
// Parameters:
// - bound: the spatial boundaries of the tree.
// - maxDepth: the maximum depth of the tree.
// - capacity: the maximum number of entities that a node can hold.
// - optional: variadic optional parameters to configure the octree.
func NewQuadtree(bound bounds.Bound, maxDepth int, capacity int, optional ...tree.Optional) (*QuadTree, error) {
	if root, err := treenode.NewTreeNode(consts.Dim3, nil, bound, 0, maxDepth, capacity); err != nil {
		return nil, err
	} else {
		o := &QuadTree{
			root: root,
		}
		for _, opt := range optional {
			opt(o)
		}
		if o.scale == nil {
			o.scale = func(v []float32) geo.Vec3Int {
				return geo.NewVec3Int(int32(v[0]), int32(v[1]), int32(v[2]))
			}
		}
		return o, nil
	}
}

// Add adds an entity to the quadtree.
// Parameters:
// - entity: the spatial entity to be added.
// Returns true if the entity was added successfully, false otherwise.
func (q *QuadTree) Add(entity siface.ISpatial) bool {
	return q.root.Add(entity)
}

// Remove removes an entity from the quadtree by its ID.
// Parameters:
// - entityId: the ID of the entity to be removed.
// Returns true if the entity was removed successfully, false otherwise.
func (q *QuadTree) Remove(entityId int64) bool {
	return q.root.Remove(entityId, q.mergeIf)
}

// GetSurroundingEntities finds entities within a certain radius of a center point.
// Parameters:
// - center: the center point to search around.
// - radius: the radius within which to search for entities.
// - filters: optional filters to apply to the entities.
// Returns a slice of spatial entities within the specified radius.
func (q *QuadTree) GetSurroundingEntities(center []float32, radius float32, filters ...func(entity siface.ISpatial) bool) []siface.ISpatial {
	return q.root.FindEntities(q.scale(center), radius, filters...)
}

// SetScale sets a custom scale function for the QuadTree.
// The scale function converts a float32 slice to a geo.Vec3Int.
// Parameters:
// - f: the custom scale function.
func (q *QuadTree) SetScale(f func(v []float32) geo.Vec3Int) {
	q.scale = f
}

// SetMergeIf sets the mergeIf field of the QuadTree.
// If you want to merge the node when removing an entity, you can set merge to true.
// Parameters:
// - merge: the flag to determine if nodes should be merged when removing an entity.
func (q *QuadTree) SetMergeIf(merge bool) {
	q.mergeIf = merge
}
