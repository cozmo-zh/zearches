package treenode

import (
	"github.com/cozmo-zh/zearches/consts"
	"github.com/cozmo-zh/zearches/internal/pkg/tree/mocks"
	"github.com/cozmo-zh/zearches/pkg/bounds"
	"github.com/cozmo-zh/zearches/pkg/siface"
	"testing"

	"github.com/cozmo-zh/zearches/pkg/geo"
	"github.com/stretchr/testify/assert"
)

func TestNewTreeNode(t *testing.T) {
	maxDepth := 4
	capacity := 10
	b := bounds.NewBound(geo.NewVec3Int(0, 0, 0), geo.NewVec3Int(10, 10, 10))
	node, _ := NewTreeNode(consts.Dim3, nil, b, 0, 0, maxDepth, capacity)

	assert.NotNil(t, node)
	assert.Equal(t, 0, node.GetEntityList().Len())
	assert.Equal(t, b, node.Bound())
}

func TestAddEntity(t *testing.T) {
	maxDepth := 4
	capacity := 1
	b := bounds.NewBound(geo.NewVec3Int(0, 0, 0), geo.NewVec3Int(10, 10, 10))
	node, _ := NewTreeNode(consts.Dim3, nil, b, 0, 0, maxDepth, capacity)

	spatial1 := mocks.CreateMockSpatial(1, 2, 2, 2)
	spatial2 := mocks.CreateMockSpatial(2, 8, 8, 8)

	node.Add(spatial1)
	node.Add(spatial2)

	assert.Equal(t, 0, node.GetEntityList().Len())
	assert.Equal(t, 1, node.Children().GetChild(0).Size())
	assert.Equal(t, 1, node.Children().GetChild(7).Size())

}

func TestRemoveEntity(t *testing.T) {
	maxDepth := 2
	capacity := 1
	b := bounds.NewBound(geo.NewVec3Int(0, 0, 0), geo.NewVec3Int(10, 10, 10))
	node, _ := NewTreeNode(consts.Dim3, nil, b, 0, 0, maxDepth, capacity)

	spatial := mocks.CreateMockSpatial(1, 5, 5, 5)
	node.Add(spatial)
	removed := node.Remove(spatial.GetID())

	assert.True(t, removed)
	assert.Equal(t, 0, node.GetEntityList().Len())
}

func TestDivideIf(t *testing.T) {
	maxDepth := 4
	capacity := 1
	b := bounds.NewBound(geo.NewVec3Int(0, 0, 0), geo.NewVec3Int(10, 10, 10))
	node, _ := NewTreeNode(consts.Dim3, nil, b, 0, 0, maxDepth, capacity)

	spatial1 := mocks.CreateMockSpatial(1, 2, 2, 2)
	spatial2 := mocks.CreateMockSpatial(2, 8, 8, 8)
	node.Add(spatial1)
	node.Add(spatial2)
	assert.False(t, node.IsLeaf())
	assert.Equal(t, 8, node.Children().ChildrenCount())
}

func TestContains(t *testing.T) {
	maxDepth := 4
	capacity := 10
	b := bounds.NewBound(geo.NewVec3Int(0, 0, 0), geo.NewVec3Int(10, 10, 10))
	node, _ := NewTreeNode(consts.Dim3, nil, b, 0, 0, maxDepth, capacity)

	spatial := mocks.CreateMockSpatial(1, 5, 5, 5)
	contains := node.Contains(spatial)

	assert.True(t, contains)
}

func TestIsLeaf(t *testing.T) {
	maxDepth := 4
	capacity := 10
	b := bounds.NewBound(geo.NewVec3Int(0, 0, 0), geo.NewVec3Int(10, 10, 10))
	node, _ := NewTreeNode(consts.Dim3, nil, b, 0, 0, maxDepth, capacity)

	assert.True(t, node.IsLeaf())
}

func TestMergeIf(t *testing.T) {
	maxDepth := 2
	capacity := 2
	b := bounds.NewBound(geo.NewVec3Int(0, 0, 0), geo.NewVec3Int(10, 10, 10))
	parent, _ := NewTreeNode(consts.Dim3, nil, b, 0, 0, maxDepth, capacity)
	assert.True(t, parent.IsLeaf())
	spatials := []siface.ISpatial{
		mocks.CreateMockSpatial(1, 1, 1, 1),
		mocks.CreateMockSpatial(2, 2, 2, 2),
		mocks.CreateMockSpatial(3, 3, 3, 3),
		mocks.CreateMockSpatial(4, 4, 4, 4),
		mocks.CreateMockSpatial(5, 5, 5, 5),
		mocks.CreateMockSpatial(6, 6, 6, 6),
		mocks.CreateMockSpatial(7, 7, 7, 7),
		mocks.CreateMockSpatial(8, 8, 8, 8),
	}
	for _, spatial := range spatials {
		parent.Add(spatial)
	}
	assert.False(t, parent.IsLeaf())
	for _, spatial := range spatials {
		parent.Remove(spatial.GetID(), true)
	}
	assert.True(t, parent.IsLeaf())
}

func TestAddEntityAtMaxDepth(t *testing.T) {
	maxDepth := 3
	capacity := 2
	b := bounds.NewBound(geo.NewVec3Int(0, 0, 0), geo.NewVec3Int(10, 10, 10))
	node, _ := NewTreeNode(consts.Dim3, nil, b, 0, 0, maxDepth, capacity)

	// Create spatial entities to fill the tree to its maximum depth
	spatials := []siface.ISpatial{
		mocks.CreateMockSpatial(1, 1, 1, 1),
		mocks.CreateMockSpatial(2, 2, 2, 2),
		mocks.CreateMockSpatial(3, 3, 3, 3),
		mocks.CreateMockSpatial(4, 4, 4, 4),
		mocks.CreateMockSpatial(5, 5, 5, 5),
		mocks.CreateMockSpatial(6, 6, 6, 6),
		mocks.CreateMockSpatial(7, 7, 7, 7),
		mocks.CreateMockSpatial(8, 8, 8, 8),
		mocks.CreateMockSpatial(9, 9, 9, 9),
	}

	// Add entities to the tree
	for _, spatial := range spatials {
		node.Add(spatial)
	}
}

func TestFindEntitiesReturnsEntitiesWithinRadius(t *testing.T) {
	maxDepth := 1
	capacity := 1
	b := bounds.NewBound(geo.NewVec3Int(0, 0, 0), geo.NewVec3Int(10, 10, 10))
	node, _ := NewTreeNode(consts.Dim3, nil, b, 0, 0, maxDepth, capacity)

	spatial1 := mocks.CreateMockSpatial(1, 4, 4, 4)
	spatial2 := mocks.CreateMockSpatial(2, 8, 8, 8)
	spatial3 := mocks.CreateMockSpatial(3, 5, 5, 5)
	node.Add(spatial1)
	node.Add(spatial2)
	node.Add(spatial3)

	center := geo.NewVec3Int(5, 5, 5)
	radius := float32(5)
	entities := node.FindEntities(center, radius)

	assert.Contains(t, entities, spatial1)
	assert.NotContains(t, entities, spatial2)
	assert.Contains(t, entities, spatial3)
}

func TestTreeNode_Range(t *testing.T) {
	maxDepth := 2
	capacity := 1
	b := bounds.NewBound(geo.NewVec3Int(0, 0, 0), geo.NewVec3Int(10, 10, 10))
	node, _ := NewTreeNode(consts.Dim3, nil, b, 0, 0, maxDepth, capacity)

	spatial1 := mocks.CreateMockSpatial(1, 4, 4, 4)
	spatial2 := mocks.CreateMockSpatial(2, 8, 8, 8)
	spatial3 := mocks.CreateMockSpatial(3, 5, 5, 5)
	node.Add(spatial1)
	node.Add(spatial2)
	node.Add(spatial3)

	entities := make([]siface.ISpatial, 0)
	nodes := make([]*TreeNode, 0)
	node.Range(func(n *TreeNode) bool {
		nodes = append(nodes, n)
		n.RangeEntities(
			func(s siface.ISpatial) bool {
				entities = append(entities, s)
				return true
			},
		)
		return true
	})
	assert.Len(t, nodes, 9)
	assert.Equal(t, nodes[0], node)
	assert.Equal(t, nodes[0].depth, 0)
	assert.Equal(t, nodes[1].depth, 1)
	assert.Equal(t, nodes[2].depth, 1)
	assert.Equal(t, nodes[3].depth, 1)
	assert.Equal(t, nodes[4].depth, 1)
	assert.Equal(t, nodes[5].depth, 1)
	assert.Equal(t, nodes[6].depth, 1)
	assert.Equal(t, nodes[7].depth, 1)
	assert.Equal(t, nodes[8].depth, 1)
	assert.Contains(t, entities, spatial1)
	assert.Contains(t, entities, spatial2)
	assert.Contains(t, entities, spatial3)

}
