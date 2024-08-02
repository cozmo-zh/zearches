// Package rtree .
package rtree

import (
	"github.com/cozmo-zh/zearches/pkg/geo"
	"github.com/cozmo-zh/zearches/pkg/siface"
	"github.com/dhconnelly/rtreego"
)

type RTree struct {
	origin *rtreego.Rtree
}

func (R *RTree) Add(entity siface.ISpatial) error {
	return nil
}

func (R *RTree) Remove(entityId int64) bool {
	return false
}

func (R *RTree) GetSurroundingEntities(center geo.Vec3Int, radius float32, filters ...func(entity siface.ISpatial) bool) []siface.ISpatial {
	return nil
}
