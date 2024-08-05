// Package quadtree provides an implementation of a quadtree data structure for spatial partitioning.
package quadtree

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

// QuadTree represents a quadtree data structure.
type QuadTree struct {
	root   *treenode.TreeNode // The root node of the quadtree.
	option *option.OptionalSettings
}

// NewQuadtree creates a new Octree.
// Parameters:
// - bound: the spatial boundaries of the tree.
// - maxDepth: the maximum depth of the tree.
// - capacity: the maximum number of entities that a node can hold.
// - optional: variadic optional parameters to configure the octree.
func NewQuadtree(bound bounds.Bound, maxDepth int, capacity int, optional ...option.Optional) (*QuadTree, error) {
	if root, err := treenode.NewTreeNode(consts.Dim2, nil, bound, 0, 0, maxDepth, capacity); err != nil {
		return nil, err
	} else {
		o := &QuadTree{
			root:   root,
			option: option.OptionalDefault(),
		}
		for _, opt := range optional {
			opt(o.option)
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
	return q.root.Remove(entityId, q.option.MergeIf())
}

// GetSurroundingEntities finds entities within a certain radius of a center point.
// Parameters:
// - center: the center point to search around.
// - radius: the radius within which to search for entities.
// - filters: optional filters to apply to the entities.
// Returns a slice of spatial entities within the specified radius.
func (q *QuadTree) GetSurroundingEntities(center []float32, radius float32, filters ...func(entity siface.ISpatial) bool) []siface.ISpatial {
	return q.root.FindEntities(q.option.ScaleFunc(center), radius, filters...)
}

func (q *QuadTree) ToDot() error {
	const fileName = "quadtree.dot"
	if q.option.DrawPath() == "" {
		return fmt.Errorf("draw path not set")
	}
	// 准备写文件
	if file, err := os.OpenFile(path.Join(q.option.DrawPath(), fileName), os.O_CREATE|os.O_WRONLY, 0644); err != nil {
		return err
	} else {
		defer file.Close()
		return tree.ToDot(tree.GetTemplate(), q.root, file)
	}
}
