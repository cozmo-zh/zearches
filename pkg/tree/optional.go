// Package tree .
package tree

import "github.com/cozmo-zh/zearches/pkg/geo"

type IOptional interface {
	SetScale(f func(v []float32) geo.Vec3Int)
	SetMergeIf(merge bool)
}

// Optional is a function type used to configure optional parameters for the Octree.
type Optional func(o IOptional)

// WithScale sets a custom scale function for the Octree.
// The scale function converts a float32 slice to a geo.Vec3Int.
// For example, if you want to convert the float32 slice to a Vec3Int by dividing each element by 10,
// like 100.0, 100.0, 100.0 -> 10, 10, 10
func WithScale(f func(v []float32) geo.Vec3Int) Optional {
	return func(o IOptional) {
		o.SetScale(f)
	}
}

// WithMergeIf sets the mergeIf field of the Octree.
// If you want to merge the node when removing an entity, you can set merge to true.
func WithMergeIf(merge bool) Optional {
	return func(o IOptional) {
		o.SetMergeIf(merge)
	}
}
