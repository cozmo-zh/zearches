// Package siface .
package siface

import (
	"github.com/cozmo-zh/zearches/pkg/bounds"
	"github.com/cozmo-zh/zearches/pkg/geo"
)

// ISpatial spatial entity interface.
type ISpatial interface {
	GetID() int64             // GetID returns the ID of the spatial entity.
	GetLocation() geo.Vec3Int // returns the location of the spatial entity.
	// GetBound returns the boundary of the spatial entity.
	//
	//it is meaningless in octree and quadtree, because octree and quadtree treat all entities as a point,
	// but it is meaningful in rtree
	GetBound() bounds.Bound
}
