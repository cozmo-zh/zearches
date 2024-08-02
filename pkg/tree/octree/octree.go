// Package octree provides an implementation of an octree data structure for spatial partitioning.
package octree

import (
	"github.com/cozmo-zh/zearches/pkg/bounds"
	"github.com/cozmo-zh/zearches/pkg/consts"
	"github.com/cozmo-zh/zearches/pkg/geo"
	"github.com/cozmo-zh/zearches/pkg/siface"
	"github.com/cozmo-zh/zearches/pkg/tree"
	"github.com/cozmo-zh/zearches/pkg/tree/treenode"
)

// Octree represents an octree data structure.
type Octree struct {
	root    *treenode.TreeNode            // The root node of the octree.
	scale   func(v []float32) geo.Vec3Int // Function to scale float32 slice to geo.Vec3Int.
	mergeIf bool                          // Flag to determine if nodes should be merged when removing an entity.
}

func (o *Octree) SetScale(f func(v []float32) geo.Vec3Int) {
	o.scale = f
}

func (o *Octree) SetMergeIf(merge bool) {
	o.mergeIf = merge
}

// NewOctree creates a new Octree.
// Parameters:
// - bound: the spatial boundaries of the tree.
// - maxDepth: the maximum depth of the tree.
// - capacity: the maximum number of entities that a node can hold.
// - optional: variadic optional parameters to configure the octree.
func NewOctree(bound bounds.Bound, maxDepth int, capacity int, optional ...tree.Optional) (*Octree, error) {
	if root, err := treenode.NewTreeNode(consts.Dim3, nil, bound, 0, maxDepth, capacity); err != nil {
		return nil, err
	} else {
		o := &Octree{
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

// Add adds an entity to the octree.
// Parameters:
// - entity: the spatial entity to be added.
// Returns true if the entity was added successfully, false otherwise.
func (o *Octree) Add(entity siface.ISpatial) bool {
	return o.root.Add(entity)
}

// Remove removes an entity from the octree by its ID.
// Parameters:
// - entityId: the ID of the entity to be removed.
// Returns true if the entity was removed successfully, false otherwise.
func (o *Octree) Remove(entityId int64) bool {
	return o.root.Remove(entityId, o.mergeIf)
}

// GetSurroundingEntities finds entities within a certain radius of a center point.
// Parameters:
// - center: the center point to search around.
// - radius: the radius within which to search for entities.
// - filters: optional filters to apply to the entities.
// Returns a slice of spatial entities within the specified radius.
func (o *Octree) GetSurroundingEntities(center []float32, radius float32, filters ...func(entity siface.ISpatial) bool) []siface.ISpatial {
	return o.root.FindEntities(o.scale(center), radius, filters...)
}
