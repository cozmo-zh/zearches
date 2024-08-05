// Package treenode .
package treenode

import (
	"github.com/cozmo-zh/zearches/consts"
	"github.com/cozmo-zh/zearches/internal/pkg/tree/mocks"
	"github.com/cozmo-zh/zearches/pkg/bounds"
	"github.com/cozmo-zh/zearches/pkg/geo"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDivideCreatesFourChildren(t *testing.T) {
	parentBound := bounds.NewBound(geo.NewVec3Int(0, 0, 0), geo.NewVec3Int(10, 0, 10))
	parent, _ := NewTreeNode(consts.Dim2, nil, parentBound, 0, 0, 4, 10)
	d2 := NewD2()
	d2.Divide(parent, 1)

	assert.Equal(t, 4, d2.ChildrenCount())
	assert.NotNil(t, d2.GetChild(0))
	assert.NotNil(t, d2.GetChild(1))
	assert.NotNil(t, d2.GetChild(2))
	assert.NotNil(t, d2.GetChild(3))
}

func TestDivideCorrectlySetsBounds(t *testing.T) {
	parentBound := bounds.NewBound(geo.NewVec3Int(0, 0, 0), geo.NewVec3Int(10, 0, 10))
	parent, _ := NewTreeNode(consts.Dim2, nil, parentBound, 0, 0, 4, 10)
	d2 := NewD2()
	d2.Divide(parent, 1)

	child0 := d2.GetChild(0)
	child1 := d2.GetChild(1)
	child2 := d2.GetChild(2)
	child3 := d2.GetChild(3)
	/*
	*   |1|2|
	*   --.--
	*   |0|3|
	 */
	assert.Equal(t, bounds.NewBound(geo.NewVec3Int(0, 0, 0), geo.NewVec3Int(5, 0, 5)), child0.Bound())
	assert.Equal(t, bounds.NewBound(geo.NewVec3Int(0, 0, 5), geo.NewVec3Int(5, 0, 10)), child1.Bound())
	assert.Equal(t, bounds.NewBound(geo.NewVec3Int(5, 0, 5), geo.NewVec3Int(10, 0, 10)), child2.Bound())
	assert.Equal(t, bounds.NewBound(geo.NewVec3Int(5, 0, 0), geo.NewVec3Int(10, 0, 5)), child3.Bound())
}

func TestClearRemovesAllChildren(t *testing.T) {
	d2 := NewD2()
	parentBound := bounds.NewBound(geo.NewVec3Int(0, 0, 0), geo.NewVec3Int(10, 0, 10))
	parent, _ := NewTreeNode(consts.Dim2, nil, parentBound, 0, 0, 4, 10)
	d2.Divide(parent, 1)

	d2.Clear()

	assert.Nil(t, d2.GetChild(0))
	assert.Nil(t, d2.GetChild(1))
	assert.Nil(t, d2.GetChild(2))
	assert.Nil(t, d2.GetChild(3))
}

func TestD2_Contains(t *testing.T) {
	d2 := NewD2()
	parentBound := bounds.NewBound(geo.NewVec3Int(0, 0, 0), geo.NewVec3Int(10, 0, 10))
	parent, _ := NewTreeNode(consts.Dim2, nil, parentBound, 0, 0, 4, 10)
	d2.Divide(parent, 1)
	flag := d2.Contains(d2.GetChild(0), mocks.CreateMockSpatial(1, 1, 0, 1))
	assert.True(t, flag)
	assert.False(t, d2.Contains(d2.GetChild(0), mocks.CreateMockSpatial(2, 6, 0, 6)))
	assert.True(t, d2.Contains(d2.GetChild(0), mocks.CreateMockSpatial(3, 1, 10, 1))) // ignore y
}

func TestD2_Intersects(t *testing.T) {
	d2 := NewD2()
	parentBound := bounds.NewBound(geo.NewVec3Int(0, 0, 0), geo.NewVec3Int(10, 0, 10))
	parent, _ := NewTreeNode(consts.Dim2, nil, parentBound, 0, 0, 4, 10)
	d2.Divide(parent, 1)
	flag := d2.Intersects(d2.GetChild(0), bounds.NewBound(geo.NewVec3Int(1, 0, 1), geo.NewVec3Int(3, 0, 3)))
	assert.True(t, flag)
	assert.False(t, d2.Intersects(d2.GetChild(0), bounds.NewBound(geo.NewVec3Int(6, 0, 6), geo.NewVec3Int(8, 0, 8))))
}
