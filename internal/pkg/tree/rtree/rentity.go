// Package rtree .
package rtree

import (
	"github.com/cozmo-zh/zearches/pkg/siface"
	"github.com/dhconnelly/rtreego"
)

// REntity represents a spatial entity in RTree.
type REntity struct {
	siface.ISpatial
	rect rtreego.Rect
}

// NewREntity creates a new REntity.
func NewREntity(spatial siface.ISpatial) (*REntity, error) {
	const minL = 0.00001
	lx := float64(spatial.GetBound().Max.X() - spatial.GetBound().Min.X())
	ly := float64(spatial.GetBound().Max.Y() - spatial.GetBound().Min.Y())
	lz := float64(spatial.GetBound().Max.Z() - spatial.GetBound().Min.Z())
	if lx == 0 {
		lx = minL
	}
	if ly == 0 {
		ly = minL
	}
	if lz == 0 {
		lz = minL
	}
	l := []float64{lx, ly, lz}
	if bd, err := rtreego.NewRect(spatial.GetBound().Min.ToFloat64(), l); err == nil {
		return &REntity{
			ISpatial: spatial,
			rect:     bd,
		}, nil
	} else {
		return nil, err
	}

}

func (r *REntity) Bounds() rtreego.Rect {
	return r.rect
}
