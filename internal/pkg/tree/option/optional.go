// Package option .
package option

import "github.com/cozmo-zh/zearches/pkg/geo"

// OptionalSettings holds configuration options for creating spatial trees.
type OptionalSettings struct {
	mergeIf   bool                          // Merge the node when removing an entity. not suggested to set to true
	scaleFunc func(v []float32) geo.Vec3Int // Scale the float32 slice to Vec3Int
	path      string                        // the path to draw the tree
}

// Optional is a function type used to configure optional parameters for the Octree.
type Optional func(o *OptionalSettings)

// OptionalDefault returns the default OptionalSettings.
func OptionalDefault() *OptionalSettings {
	return &OptionalSettings{
		mergeIf: false,
		scaleFunc: func(v []float32) geo.Vec3Int {
			return geo.NewVec3Int(int32(v[0]), int32(v[1]), int32(v[2]))
		},
	}
}

// WithScale sets a custom scale function for the Octree.
// The scale function converts a float32 slice to a geo.Vec3Int.
// For example, if you want to convert the float32 slice to a Vec3Int by dividing each element by 10,
// like 100.0, 100.0, 100.0 -> 10, 10, 10
func WithScale(f func(v []float32) geo.Vec3Int) Optional {
	return func(o *OptionalSettings) {
		o.scaleFunc = f
	}
}

// WithMergeIf sets the mergeIf field of the Octree.
// If you want to merge the node when removing an entity, you can set merge to true.
func WithMergeIf(merge bool) Optional {
	return func(o *OptionalSettings) {
		o.mergeIf = merge
	}
}

// WithDrawPath sets the path to draw the tree.
func WithDrawPath(path string) Optional {
	return func(o *OptionalSettings) {
		o.path = path
	}
}

// MergeIf returns the mergeIf field of the Octree.
func (o *OptionalSettings) MergeIf() bool {
	return o.mergeIf
}

// ScaleFunc returns the scaleFunc field of the Octree.
func (o *OptionalSettings) ScaleFunc(v []float32) geo.Vec3Int {
	if o.scaleFunc == nil {
		return geo.NewVec3Int(int32(v[0]), int32(v[1]), int32(v[2]))
	}
	return o.scaleFunc(v)
}

// DrawPath returns the path to draw the tree.
func (o *OptionalSettings) DrawPath() string {
	return o.path
}
