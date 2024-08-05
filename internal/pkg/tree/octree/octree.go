// Package octree provides an implementation of an octree data structure for spatial partitioning.
package octree

import (
	"fmt"
	"github.com/cozmo-zh/zearches/consts"
	"github.com/cozmo-zh/zearches/internal/pkg/tree"
	"github.com/cozmo-zh/zearches/internal/pkg/tree/option"
	"github.com/cozmo-zh/zearches/internal/pkg/tree/treenode"
	"github.com/cozmo-zh/zearches/pkg/bounds"
	"github.com/cozmo-zh/zearches/pkg/siface"
	"os"
	"path"
)

// Octree represents an octree data structure.
type Octree struct {
	root   *treenode.TreeNode // The root node of the octree.
	option *option.OptionalSettings
}

// NewOctree creates a new Octree.
// Parameters:
// - bound: the spatial boundaries of the tree.
// - maxDepth: the maximum depth of the tree.
// - capacity: the maximum number of entities that a node can hold.
// - optional: variadic optional parameters to configure the octree.
func NewOctree(bound bounds.Bound, maxDepth int, capacity int, optional ...option.Optional) (*Octree, error) {
	if root, err := treenode.NewTreeNode(consts.Dim3, nil, bound, 0, 0, maxDepth, capacity); err != nil {
		return nil, err
	} else {
		o := &Octree{
			root:   root,
			option: option.OptionalDefault(),
		}
		for _, opt := range optional {
			opt(o.option)
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
	return o.root.Remove(entityId, o.option.MergeIf())
}

// GetSurroundingEntities finds entities within a certain radius of a center point.
// Parameters:
// - center: the center point to search around.
// - radius: the radius within which to search for entities.
// - filters: optional filters to apply to the entities.
// Returns a slice of spatial entities within the specified radius.
func (o *Octree) GetSurroundingEntities(center []float32, radius float32, filters ...func(entity siface.ISpatial) bool) []siface.ISpatial {
	return o.root.FindEntities(o.option.ScaleFunc(center), radius, filters...)
}

// ToDot generates a dot file for the search tree.
func (o *Octree) ToDot() error {
	const fileName = "octree.dot"
	if o.option.DrawPath() == "" {
		return fmt.Errorf("draw path not set")
	}
	// 准备写文件
	if file, err := os.OpenFile(path.Join(o.option.DrawPath(), fileName), os.O_CREATE|os.O_WRONLY, 0644); err != nil {
		return err
	} else {
		defer file.Close()
		return tree.ToDot(tree.GetTemplate(), o.root, file)
	}
}
