// Package octree .
package octree

import (
	"github.com/cozmo-zh/zearches/pkg/bound"
	"github.com/cozmo-zh/zearches/pkg/consts"
	"github.com/cozmo-zh/zearches/pkg/geo"
	"github.com/cozmo-zh/zearches/pkg/siface"
	"github.com/cozmo-zh/zearches/pkg/treenode"
)

// Octree .
type Octree struct {
	root *treenode.TreeNode
}

// NewOctree .
func NewOctree(dim consts.Dim, bound bound.Bound, maxDepth int, capacity int) *Octree {
	return &Octree{
		root: treenode.NewTreeNode(dim, nil, bound, 0, &maxDepth, &capacity),
	}
}

// Add .
func (o *Octree) Add(entity siface.ISpatial) bool {
	return o.root.Add(entity)
}

// Remove .
func (o *Octree) Remove(entityId int64) bool {
	return o.root.Remove(entityId)
}

// GetSurroundingEntities .
func (o *Octree) GetSurroundingEntities(center geo.Vec3Int, radius float32, filters ...func(entity siface.ISpatial) bool) []siface.ISpatial {

	return nil
}
