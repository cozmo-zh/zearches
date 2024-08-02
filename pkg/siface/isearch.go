// Package siface .
package siface

type ISearch interface {
	Add(entity ISpatial) bool
	Remove(entityId int64) bool
	GetSurroundingEntities(center []float32, radius float32, filters ...func(entity ISpatial) bool) []ISpatial
}
