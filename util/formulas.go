// Package util provides utility functions for distance calculations.
package util

import "math"

// Distance2D calculates the 2D distance between two points p1 and p2.
func Distance2D(p1, p2 []float32) float32 {
	dx := p1[0] - p2[0]
	dz := p1[2] - p2[2]
	return float32(math.Hypot(float64(dx), float64(dz)))
}

// Distance3D calculates the 3D distance between two points p1 and p2.
func Distance3D(p1, p2 []float32) float32 {
	dx := p1[0] - p2[0]
	dy := p1[1] - p2[1]
	dz := p1[2] - p2[2]
	return float32(math.Sqrt(float64(dx*dx + dy*dy + dz*dz)))
}

// WithinDistance2D checks if the 2D distance between two points p1 and p2
func WithinDistance2D(p1, p2 []float32, distance float32) bool {
	dx := p1[0] - p2[0]
	dz := p1[2] - p2[2]
	return dx*dx+dz*dz <= distance*distance
}

// WithinDistance3D checks if the 3D distance between two points p1 and p2
func WithinDistance3D(p1, p2 []float32, distance float32) bool {
	dx := p1[0] - p2[0]
	dy := p1[1] - p2[1]
	dz := p1[2] - p2[2]
	return dx*dx+dy*dy+dz*dz <= distance*distance
}

// Normalize normalizes a 3D vector v.
func Normalize(v []float32) []float32 {
	lens := math.Sqrt(float64(v[0]*v[0] + v[1]*v[1] + v[2]*v[2]))
	return []float32{float32(float64(v[0]) / lens), float32(float64(v[1]) / lens), float32(float64(v[2]) / lens)}
}

// Normalize2D normalizes a 2D vector v.
func Normalize2D(v []float32) []float32 {
	lens := math.Sqrt(float64(v[0]*v[0] + v[2]*v[2]))
	return []float32{float32(float64(v[0]) / lens), 0, float32(float64(v[2]) / lens)}
}
