// Package bounds provides the implementation of spatial boundaries.
package bounds

import (
	"github.com/cozmo-zh/zearches/pkg/geo"
	"github.com/cozmo-zh/zearches/pkg/util"
)

// Bound represents a spatial boundary.
type Bound struct {
	Min    geo.Vec3Int // The minimum point of the spatial boundary.
	Max    geo.Vec3Int // The maximum point of the spatial boundary.
	Center geo.Vec3Int // The center point of the spatial boundary.
	Length float32     // The length of the spatial boundary.
}

// NewBound creates a new Bound instance.
//
// Parameters:
// - min: The minimum point of the spatial boundary.
// - max: The maximum point of the spatial boundary.
//
// Returns:
// - A new Bound instance.
func NewBound(min, max geo.Vec3Int) Bound {
	return Bound{
		Min:    min,
		Max:    max,
		Center: geo.NewVec3Int((min.X()+max.X())/2, (min.Y()+max.Y())/2, (min.Z()+max.Z())/2),
		Length: util.Distance3D(min.ToFloat32(), max.ToFloat32()),
	}
}
