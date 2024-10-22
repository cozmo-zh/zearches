// Package geo .
package geo

// Vec3Int represents a 3-dimensional vector with integer components.
type Vec3Int []int32

// NewVec3Int creates a new Vec3Int instance.
//
// Parameters:
// - x: The x-coordinate.
// - y: The y-coordinate.
// - z: The z-coordinate.
//
// Returns:
// - A new Vec3Int instance.
// if it's 2d, you can only use x and z (v.X(), v.Z())
func NewVec3Int(x, y, z int32) Vec3Int {
	return Vec3Int{x, y, z}
}

// X returns the x-coordinate of the vector.
//
// Returns:
// - The x-coordinate.
func (v Vec3Int) X() int32 {
	return v[0]
}

// Y returns the y-coordinate of the vector.
//
// Returns:
// - The y-coordinate.
func (v Vec3Int) Y() int32 {
	return v[1]
}

// Z returns the z-coordinate of the vector.
//
// Returns:
// - The z-coordinate.
func (v Vec3Int) Z() int32 {
	return v[2]
}

// ToFloat32 converts the vector to a float32 slice.
func (v Vec3Int) ToFloat32() []float32 {
	return []float32{float32(v[0]), float32(v[1]), float32(v[2])}
}

// ToFloat64 converts the vector to a float64 slice.
func (v Vec3Int) ToFloat64() []float64 {
	return []float64{float64(v[0]), float64(v[1]), float64(v[2])}
}
