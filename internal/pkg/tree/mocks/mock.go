// Package mocks .
package mocks

import (
	"github.com/cozmo-zh/zearches/pkg/bounds"
	"github.com/cozmo-zh/zearches/pkg/geo"
	"github.com/cozmo-zh/zearches/pkg/siface"
)

type MockSpatial struct {
	ID       int64
	Location geo.Vec3Int
	Bound    bounds.Bound
}

func (m *MockSpatial) GetBound() bounds.Bound {
	return m.Bound
}

func (m *MockSpatial) GetID() int64 {
	return m.ID
}

func (m *MockSpatial) GetLocation() geo.Vec3Int {
	return m.Location
}

func CreateMockSpatial(id int64, x, y, z int32, bound ...bounds.Bound) siface.ISpatial {
	m := &MockSpatial{
		ID:       id,
		Location: geo.NewVec3Int(x, y, z),
	}
	if len(bound) > 0 {
		m.Bound = bound[0]
	} else {
		m.Bound = bounds.NewBound(geo.NewVec3Int(x, y, z), geo.NewVec3Int(x, y, z))
	}
	return m
}
