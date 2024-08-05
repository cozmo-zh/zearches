// Package rtree .
package rtree

import (
	"fmt"
	"github.com/cozmo-zh/zearches/consts"
	"github.com/cozmo-zh/zearches/pkg/siface"
	"github.com/dhconnelly/rtreego"
)

// RTree .
type RTree struct {
	origin   *rtreego.Rtree
	entities map[int64]*REntity
}

// NewRTree .
func NewRTree(dim consts.Dim, min, max int) *RTree {
	return &RTree{
		origin:   rtreego.NewTree(int(dim), min, max),
		entities: make(map[int64]*REntity),
	}
}

// Add .
func (r *RTree) Add(entity siface.ISpatial) bool {
	if e, err := NewREntity(entity); err != nil {
		//TODO  log error
		return false
	} else {
		r.entities[entity.GetID()] = e
		r.origin.Insert(e)
		return true
	}
}

// Remove .
func (r *RTree) Remove(entityId int64) bool {
	if e, ok := r.entities[entityId]; ok {
		r.origin.Delete(e)
		delete(r.entities, entityId)
		return true
	}
	return false
}

func (r *RTree) GetSurroundingEntities(center []float32, radius float32, filters ...func(entity siface.ISpatial) bool) []siface.ISpatial {
	ret := make([]siface.ISpatial, 0)
	// build a search rect with the center and radius
	rmin := []float64{float64(center[0] - radius), float64(center[1] - radius), float64(center[2] - radius)}
	length := float64(radius * 2)
	l := []float64{length, length, length}

	if rect, err := rtreego.NewRect(rmin, l); err != nil {
		return ret
	} else {
		entities := r.origin.SearchIntersect(rect)
	outer:
		for _, e := range entities {
			if re, ok := e.(*REntity); ok {
				for _, f := range filters {
					if !f(re.ISpatial) {
						continue outer
					}
				}
				ret = append(ret, re.ISpatial)
			}
		}
	}
	return ret
}

func (r *RTree) ToDot() error {
	return fmt.Errorf("rtree not support draw")
}
