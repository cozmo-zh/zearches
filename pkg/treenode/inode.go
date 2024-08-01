// Package treenode .
package treenode

import (
	"container/list"
	"github.com/cozmo-zh/zearches/pkg/bounds"
	"github.com/cozmo-zh/zearches/pkg/siface"
)

type INode interface {
	Add(spatial siface.ISpatial) bool
	Bound() bounds.Bound
	MaxDepth() int
	Capacity() int
	Contains(spatial siface.ISpatial) bool
	Remove(spatialId int64, merge ...bool) bool
	FindEntitiesInBound(bound bounds.Bound, filters ...func(entity siface.ISpatial) bool) []siface.ISpatial
	Clear()
	ClearChildren()
	Parent() INode
	Size() int
	Children() IChildren
	IsLeaf() bool
	GetEntityList() *list.List
}

type IChildren interface {
	Divide(parent INode, depth int)
	ChildrenCount() int
	GetChild(index int) INode
	Clear()
}
