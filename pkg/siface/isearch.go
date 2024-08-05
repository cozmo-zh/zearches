// Package siface .
package siface

// ISearch  interface for search, like Octree, QuadTree, RTree, etc.
type ISearch interface {
	// Add adds an entity to the search tree.
	Add(entity ISpatial) bool
	// Remove removes an entity from the search tree by its ID.
	Remove(entityId int64) bool
	// GetSurroundingEntities finds entities within a certain radius of a center point.
	GetSurroundingEntities(center []float32, radius float32, filters ...func(entity ISpatial) bool) []ISpatial
	// ToDot generates a dot file for the search tree.
	ToDot() error
}
