package treenode

import (
	"testing"

	"github.com/cozmo-zh/zearches/pkg/bound"
	"github.com/cozmo-zh/zearches/pkg/consts"
	"github.com/cozmo-zh/zearches/pkg/geo"
	"github.com/cozmo-zh/zearches/pkg/siface"
	"github.com/stretchr/testify/assert"
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

func TestNewTreeNode(t *testing.T) {
	maxDepth := 4
	capacity := 10
	b := bound.NewBound(geo.NewVec3Int(0, 0, 0), geo.NewVec3Int(10, 10, 10))
	node := NewTreeNode(consts.Dim3, nil, b, 0, &maxDepth, &capacity)

	assert.NotNil(t, node)
	assert.Equal(t, 0, node.GetEntityList().Len())
	assert.Equal(t, b, node.Bound())
}

func TestAddEntity(t *testing.T) {
	maxDepth := 4
	capacity := 1
	b := bound.NewBound(geo.NewVec3Int(0, 0, 0), geo.NewVec3Int(10, 10, 10))
	node := NewTreeNode(consts.Dim3, nil, b, 0, &maxDepth, &capacity)

	spatial1 := createMockSpatial(1, 2, 2, 2)
	spatial2 := createMockSpatial(2, 8, 8, 8)

	node.Add(spatial1)
	node.Add(spatial2)

	assert.Equal(t, 0, node.GetEntityList().Len())
	assert.Equal(t, 1, node.children[0].GetEntityList().Len())
	assert.Equal(t, 1, node.children[7].GetEntityList().Len())

}

func TestRemoveEntity(t *testing.T) {
	maxDepth := 4
	capacity := 10
	b := bound.NewBound(geo.NewVec3Int(0, 0, 0), geo.NewVec3Int(10, 10, 10))
	node := NewTreeNode(consts.Dim3, nil, b, 0, &maxDepth, &capacity)

	spatial := createMockSpatial(1, 5, 5, 5)
	node.Add(spatial)
	removed := node.Remove(spatial.GetID())

	assert.True(t, removed)
	assert.Equal(t, 0, node.GetEntityList().Len())
}

func TestDivideIf(t *testing.T) {
	maxDepth := 4
	capacity := 1
	b := bound.NewBound(geo.NewVec3Int(0, 0, 0), geo.NewVec3Int(10, 10, 10))
	node := NewTreeNode(consts.Dim3, nil, b, 0, &maxDepth, &capacity)

	spatial1 := createMockSpatial(1, 2, 2, 2)
	spatial2 := createMockSpatial(2, 8, 8, 8)
	node.Add(spatial1)
	node.Add(spatial2)
	assert.True(t, node.IsLeaf())

	divided := node.DivideIf()

	assert.True(t, divided)
	assert.Equal(t, 8, len(node.children))
	assert.False(t, node.IsLeaf())
}

func TestContains(t *testing.T) {
	maxDepth := 4
	capacity := 10
	b := bound.NewBound(geo.NewVec3Int(0, 0, 0), geo.NewVec3Int(10, 10, 10))
	node := NewTreeNode(consts.Dim3, nil, b, 0, &maxDepth, &capacity)

	spatial := createMockSpatial(1, 5, 5, 5)
	contains := node.Contains(spatial)

	assert.True(t, contains)
}

func TestIsLeaf(t *testing.T) {
	maxDepth := 4
	capacity := 10
	b := bound.NewBound(geo.NewVec3Int(0, 0, 0), geo.NewVec3Int(10, 10, 10))
	node := NewTreeNode(consts.Dim3, nil, b, 0, &maxDepth, &capacity)

	assert.True(t, node.IsLeaf())
}

func TestMergeIf(t *testing.T) {
	maxDepth := 4
	capacity := 10
	b := bound.NewBound(geo.NewVec3Int(0, 0, 0), geo.NewVec3Int(10, 10, 10))
	parent := NewTreeNode(consts.Dim3, nil, b, 0, &maxDepth, &capacity)
	node := NewTreeNode(consts.Dim3, parent, b, 1, &maxDepth, &capacity)
	parent.children[0] = node

	spatial := createMockSpatial(1, 5, 5, 5)
	node.Add(spatial)
	merged := node.MergeIf()

	assert.True(t, merged)
	assert.Equal(t, 1, parent.GetEntityList().Len())
}

func TestAddEntityAtMaxDepth(t *testing.T) {
	maxDepth := 3
	capacity := 2
	b := bound.NewBound(geo.NewVec3Int(0, 0, 0), geo.NewVec3Int(10, 10, 10))
	node := NewTreeNode(consts.Dim3, nil, b, 0, &maxDepth, &capacity)

	// Create spatial entities to fill the tree to its maximum depth
	spatials := []siface.ISpatial{
		createMockSpatial(1, 1, 1, 1),
		createMockSpatial(2, 2, 2, 2),
		createMockSpatial(3, 3, 3, 3),
		createMockSpatial(4, 4, 4, 4),
		createMockSpatial(5, 5, 5, 5),
		createMockSpatial(6, 6, 6, 6),
		createMockSpatial(7, 7, 7, 7),
		createMockSpatial(8, 8, 8, 8),
		createMockSpatial(9, 9, 9, 9),
	}

	// Add entities to the tree
	for _, spatial := range spatials {
		node.Add(spatial)
	}
}
