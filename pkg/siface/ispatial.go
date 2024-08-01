// Package siface .
package siface

import "github.com/cozmo-zh/zearches/pkg/geo"

type ISpatial interface {
	GetID() int64
	GetLocation() geo.Vec3Int
}
