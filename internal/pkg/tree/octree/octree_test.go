package octree

import (
	"github.com/cozmo-zh/zearches/internal/pkg/tree"
	"github.com/cozmo-zh/zearches/internal/pkg/tree/mocks"
	"testing"

	"github.com/cozmo-zh/zearches/pkg/bounds"
	"github.com/cozmo-zh/zearches/pkg/geo"
	"github.com/stretchr/testify/assert"
)

func TestOctree_creation(t *testing.T) {
	bound := bounds.NewBound(geo.NewVec3Int(0, 0, 0), geo.NewVec3Int(100, 100, 100))
	oct, err := NewOctree(
		bound,
		1,
		1,
		tree.WithMergeIf(false),
		tree.WithScale(func(v []float32) geo.Vec3Int {
			return geo.NewVec3Int(int32(v[0])+1, int32(v[1])+1, int32(v[2])+1)
		},
		))
	assert.Nil(t, err)
	assert.NotNil(t, oct)
	assert.False(t, oct.mergeIf)
	assert.Equal(t, oct.scale([]float32{1.0, 2.0, 3.0}), geo.NewVec3Int(2, 3, 4))
}

func TestOctree_createFail(t *testing.T) {
	bound := bounds.NewBound(geo.NewVec3Int(0, 0, 0), geo.NewVec3Int(100, 100, 100))
	oct, err := NewOctree(bound, 0, 0)
	assert.Nil(t, oct)
	assert.NotNil(t, err)
}

func TestOctree_Add(t *testing.T) {
	bound := bounds.NewBound(geo.NewVec3Int(0, 0, 0), geo.NewVec3Int(100, 100, 100))
	oct, _ := NewOctree(bound, 1, 1) // maxDepth and capacity set to 1
	entity1 := mocks.CreateMockSpatial(1, 10, 10, 10)
	added := oct.Add(entity1)
	assert.True(t, added)
	ret := oct.GetSurroundingEntities([]float32{10, 10, 10}, 1)
	assert.True(t, len(ret) == 1)
	assert.Equal(t, entity1, ret[0])
}

func TestOctree_Remove(t *testing.T) {
	bound := bounds.NewBound(geo.NewVec3Int(0, 0, 0), geo.NewVec3Int(100, 100, 100))
	oct, _ := NewOctree(bound, 1, 1) // maxDepth and capacity set to 1
	entity1 := mocks.CreateMockSpatial(1, 10, 10, 10)
	oct.Add(entity1)
	ret := oct.GetSurroundingEntities([]float32{10, 10, 10}, 1)
	assert.True(t, len(ret) == 1)
	succ := oct.Remove(1)
	assert.True(t, succ)
	ret = oct.GetSurroundingEntities([]float32{10, 10, 10}, 1)
	assert.True(t, len(ret) == 0)
}

func TestOctree_GetSurroundingEntities(t *testing.T) {
	bound := bounds.NewBound(geo.NewVec3Int(0, 0, 0), geo.NewVec3Int(100, 100, 100))
	oct, _ := NewOctree(bound, 1, 1) // maxDepth and capacity set to 1
	entity1 := mocks.CreateMockSpatial(1, 10, 10, 10)
	entity2 := mocks.CreateMockSpatial(2, 20, 20, 20)
	oct.Add(entity1)
	oct.Add(entity2)
	ret := oct.GetSurroundingEntities([]float32{10, 10, 10}, 1)
	assert.True(t, len(ret) == 1)
	assert.Equal(t, entity1, ret[0])
	ret = oct.GetSurroundingEntities([]float32{20, 20, 20}, 1)
	assert.True(t, len(ret) == 1)
	assert.Equal(t, entity2, ret[0])
	ret = oct.GetSurroundingEntities([]float32{1, 1, 1}, 1)
	assert.True(t, len(ret) == 0)
}
