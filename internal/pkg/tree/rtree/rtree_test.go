// Package rtree .
package rtree

import (
	"github.com/cozmo-zh/zearches/consts"
	"github.com/cozmo-zh/zearches/internal/pkg/tree/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_RTree_creation(t *testing.T) {
	rtree := NewRTree(consts.Dim3, 1, 10)
	assert.NotNil(t, rtree)
}

func Test_RTree_AddAndRemoveEntity(t *testing.T) {
	rtree := NewRTree(consts.Dim3, 1, 10)
	entity := mocks.CreateMockSpatial(1, 10, 10, 10)
	added := rtree.Add(entity)
	assert.True(t, added)
	removed := rtree.Remove(entity.GetID())
	assert.True(t, removed)
}

func Test_RTree_GetSurroundingEntities(t *testing.T) {
	rtree := NewRTree(consts.Dim3, 1, 10)
	entity1 := mocks.CreateMockSpatial(1, 10, 10, 10)
	entity2 := mocks.CreateMockSpatial(2, 20, 20, 20)
	rtree.Add(entity1)
	rtree.Add(entity2)
	ret := rtree.GetSurroundingEntities([]float32{10, 10, 10}, 1)
	assert.True(t, len(ret) == 1)
	assert.Equal(t, entity1, ret[0])
	ret = rtree.GetSurroundingEntities([]float32{20, 20, 20}, 1)
	assert.True(t, len(ret) == 1)
	assert.Equal(t, entity2, ret[0])
	ret = rtree.GetSurroundingEntities([]float32{1, 1, 1}, 1)
	assert.True(t, len(ret) == 0)
}
