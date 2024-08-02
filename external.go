// Package zearches provides functions to create and manage spatial partitioning structures like octrees and quadtrees, etc.
package zearches

import (
	"github.com/cozmo-zh/zearches/internal/pkg/tree"
	"github.com/cozmo-zh/zearches/internal/pkg/tree/octree"
	"github.com/cozmo-zh/zearches/internal/pkg/tree/quadtree"
	"github.com/cozmo-zh/zearches/pkg/bounds"
	"github.com/cozmo-zh/zearches/pkg/geo"
	"github.com/cozmo-zh/zearches/pkg/siface"
)

// OptionalSettings holds configuration options for creating spatial trees.
type OptionalSettings struct {
	MergeIf   bool                          // Flag to determine if nodes should be merged when removing an entity.
	ScaleFunc func(v []float32) geo.Vec3Int // Function to scale float32 slice to geo.Vec3Int.
}

// Option is a function type used to configure OptionalSettings.
type Option func(p *OptionalSettings)

// WithScale sets a custom scale function for the OptionalSettings.
// Parameters:
// - f: the custom scale function.
func WithScale(f func(v []float32) geo.Vec3Int) Option {
	return func(s *OptionalSettings) {
		s.ScaleFunc = f
	}
}

// WithMergeIf sets the mergeIf field of the OptionalSettings.
// Parameters:
// - merge: the flag to determine if nodes should be merged when removing an entity.
func WithMergeIf(merge bool) Option {
	return func(s *OptionalSettings) {
		s.MergeIf = merge
	}
}

// CreateOctree creates a new Octree with the specified parameters.
// Parameters:
// - bound: the spatial boundaries of the tree.
// - maxDepth: the maximum depth of the tree.
// - capacity: the maximum number of entities that a node can hold.
// - opt: variadic optional parameters to configure the octree.
// Returns an ISpatial search interface and an error if creation fails.
func CreateOctree(bound bounds.Bound, maxDepth, capacity int, opt ...Option) (siface.ISearch, error) {
	s := &OptionalSettings{
		ScaleFunc: func(v []float32) geo.Vec3Int {
			return geo.NewVec3Int(int32(v[0]), int32(v[1]), int32(v[2]))
		},
	}
	for _, o := range opt {
		o(s)
	}
	if ot, err := octree.NewOctree(
		bound,
		maxDepth,
		capacity,
		tree.WithMergeIf(s.MergeIf),
		tree.WithScale(s.ScaleFunc),
	); err == nil {
		return ot, nil
	} else {
		return nil, err
	}
}

// CreateQuadtree creates a new Quadtree with the specified parameters.
// Parameters:
// - bound: the spatial boundaries of the tree.
// - maxDepth: the maximum depth of the tree.
// - capacity: the maximum number of entities that a node can hold.
// - opt: variadic optional parameters to configure the quadtree.
// Returns an ISpatial search interface and an error if creation fails.
func CreateQuadtree(bound bounds.Bound, maxDepth, capacity int, opt ...Option) (siface.ISearch, error) {
	s := &OptionalSettings{
		ScaleFunc: func(v []float32) geo.Vec3Int {
			return geo.NewVec3Int(int32(v[0]), int32(v[1]), int32(v[2]))
		},
	}
	for _, op := range opt {
		op(s)
	}
	if qt, err := quadtree.NewQuadtree(
		bound,
		maxDepth,
		capacity,
		tree.WithMergeIf(s.MergeIf),
		tree.WithScale(s.ScaleFunc),
	); err == nil {
		return qt, nil
	} else {
		return nil, err
	}
}
