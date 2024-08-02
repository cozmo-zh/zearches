// Package quadtree .
package quadtree

import (
	"github.com/cozmo-zh/zearches/pkg/bounds"
	"github.com/cozmo-zh/zearches/pkg/geo"
	"github.com/cozmo-zh/zearches/pkg/siface"
	"github.com/stretchr/testify/assert"
	"testing"
)

type MockSpatial struct {
	id       int64
	location geo.Vec3Int
}

func (m *MockSpatial) GetID() int64 {
	return m.id
}

func (m *MockSpatial) GetLocation() geo.Vec3Int {
	return m.location
}

func createMockSpatial(id int64, x, y, z int32) siface.ISpatial {
	return &MockSpatial{
		id:       id,
		location: geo.NewVec3Int(x, y, z),
	}
}

func TestQuadTree_Add(t *testing.T) {
	bound := bounds.NewBound(geo.NewVec3Int(0, 0, 0), geo.NewVec3Int(100, 100, 100))
	quad, _ := NewQuadtree(bound, 1, 1)
	entity1 := createMockSpatial(1, 10, 10, 10)
	added := quad.Add(entity1)
	assert.True(t, added)
	ret := quad.GetSurroundingEntities([]float32{10, 10, 10}, 1)
	assert.Equal(t, entity1, ret[0])
}

func TestQuadTree_Remove(t *testing.T) {
	bound := bounds.NewBound(geo.NewVec3Int(0, 0, 0), geo.NewVec3Int(100, 100, 100))
	quad, _ := NewQuadtree(bound, 5, 10)
	removed := quad.Remove(999)
	assert.False(t, removed)
	entity1 := createMockSpatial(1, 10, 10, 10)
	quad.Add(entity1)
	removed = quad.Remove(1)
	assert.True(t, removed)
}

func TestQuadTree_GetSurroundingEntities(t *testing.T) {
	bound := bounds.NewBound(geo.NewVec3Int(0, 0, 0), geo.NewVec3Int(100, 100, 100))
	quad, _ := NewQuadtree(bound, 5, 10)
	entities := quad.GetSurroundingEntities([]float32{10, 10, 10}, 10)
	assert.Len(t, entities, 0)
	entity1 := createMockSpatial(1, 10, 10, 10)
	entity2 := createMockSpatial(2, 15, 15, 15)
	quad.Add(entity1)
	quad.Add(entity2)
	filter := func(entity siface.ISpatial) bool {
		return entity.GetID() == 1
	}
	entities = quad.GetSurroundingEntities([]float32{10, 10, 10}, 10, filter)
	assert.Len(t, entities, 1)
	assert.Equal(t, entity1, entities[0])

}