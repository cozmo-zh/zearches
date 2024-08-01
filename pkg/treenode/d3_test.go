// Package treenode .
package treenode

import (
	"github.com/cozmo-zh/zearches/pkg/bounds"
	"github.com/cozmo-zh/zearches/pkg/consts"
	"github.com/cozmo-zh/zearches/pkg/geo"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDivideCreatesEightChildren(t *testing.T) {
	parentBound := bounds.NewBound(geo.NewVec3Int(0, 0, 0), geo.NewVec3Int(10, 10, 10))
	parent, _ := NewTreeNode(consts.Dim3, nil, parentBound, 0, 4, 10)
	d3 := NewD3()
	d3.Divide(parent, 1)

	assert.Equal(t, 8, d3.ChildrenCount())
	assert.NotNil(t, d3.GetChild(0))
	assert.NotNil(t, d3.GetChild(1))
	assert.NotNil(t, d3.GetChild(2))
	assert.NotNil(t, d3.GetChild(3))
	assert.NotNil(t, d3.GetChild(4))
	assert.NotNil(t, d3.GetChild(5))
	assert.NotNil(t, d3.GetChild(6))
	assert.NotNil(t, d3.GetChild(7))
}

func TestDivideCorrectlySetsBoundsForD3(t *testing.T) {
	parentBound := bounds.NewBound(geo.NewVec3Int(0, 0, 0), geo.NewVec3Int(10, 10, 10))
	parent, _ := NewTreeNode(consts.Dim3, nil, parentBound, 0, 4, 10)
	d3 := NewD3()
	d3.Divide(parent, 1)

	child0 := d3.GetChild(0)
	child1 := d3.GetChild(1)
	child2 := d3.GetChild(2)
	child3 := d3.GetChild(3)
	child4 := d3.GetChild(4)
	child5 := d3.GetChild(5)
	child6 := d3.GetChild(6)
	child7 := d3.GetChild(7)

	assert.Equal(t, bounds.NewBound(geo.NewVec3Int(0, 0, 0), geo.NewVec3Int(5, 5, 5)), child0.Bound())
	assert.Equal(t, bounds.NewBound(geo.NewVec3Int(0, 0, 5), geo.NewVec3Int(5, 5, 10)), child1.Bound())
	assert.Equal(t, bounds.NewBound(geo.NewVec3Int(0, 5, 0), geo.NewVec3Int(5, 10, 5)), child2.Bound())
	assert.Equal(t, bounds.NewBound(geo.NewVec3Int(0, 5, 5), geo.NewVec3Int(5, 10, 10)), child3.Bound())
	assert.Equal(t, bounds.NewBound(geo.NewVec3Int(5, 0, 0), geo.NewVec3Int(10, 5, 5)), child4.Bound())
	assert.Equal(t, bounds.NewBound(geo.NewVec3Int(5, 0, 5), geo.NewVec3Int(10, 5, 10)), child5.Bound())
	assert.Equal(t, bounds.NewBound(geo.NewVec3Int(5, 5, 0), geo.NewVec3Int(10, 10, 5)), child6.Bound())
	assert.Equal(t, bounds.NewBound(geo.NewVec3Int(5, 5, 5), geo.NewVec3Int(10, 10, 10)), child7.Bound())
}

func TestClearRemovesAllChildrenForD3(t *testing.T) {
	d3 := NewD3()
	parentBound := bounds.NewBound(geo.NewVec3Int(0, 0, 0), geo.NewVec3Int(10, 10, 10))
	parent, _ := NewTreeNode(consts.Dim3, nil, parentBound, 0, 4, 10)
	d3.Divide(parent, 1)

	d3.Clear()

	assert.Nil(t, d3.GetChild(0))
	assert.Nil(t, d3.GetChild(1))
	assert.Nil(t, d3.GetChild(2))
	assert.Nil(t, d3.GetChild(3))
	assert.Nil(t, d3.GetChild(4))
	assert.Nil(t, d3.GetChild(5))
	assert.Nil(t, d3.GetChild(6))
	assert.Nil(t, d3.GetChild(7))
}
