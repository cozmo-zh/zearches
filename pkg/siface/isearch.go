// Package siface .
package siface

import "github.com/cozmo-zh/zearches/pkg/geo"

type ISearch interface {
	Add(entity ISpatial) bool
	Remove(entityId int64) bool
	GetSurroundingEntities(center geo.Vec3Int, radius float32, filters ...func(entity ISpatial) bool) []ISpatial
}
