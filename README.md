# zearches
Zearches is a simple spatial segmentation/search toolkit that includes Octree, Quadtree, Rtree, and Grid-Based. It can be used to implement AOI, such as vision management in game projects, and also provides simple collision detection.

## octree

![octree](draws/octree1.png) 

## quadtree

![quadtree](draws/quadtree.png)


## Installation

```bash
go get -u github.com/cozmo-zh/zearches
```

## Usage

```go
func main() {
// create an octree
otree, _ := zearches.CreateOctree(
bounds.NewBound(geo.NewVec3Int(0, 0, 0), geo.NewVec3Int(100, 100, 100)), // bound, required
1,                          // maxDepth, required
1,                          // capacity, required
zearches.WithMergeIf(true), // Flag to determine if nodes should be merged when removing an entity , optional, default is false
zearches.WithScale(func(v []float32) geo.Vec3Int {
return geo.NewVec3Int(int32(v[0]), int32(v[1]), int32(v[2]))
}), // Function to scale float32 slice to geo.Vec3Int , optional, default is identity function
)

otree.GetSurroundingEntities([]float32{1, 1, 1}, 10, func(entity siface.ISpatial) bool {
return entity.GetID() == 999
})

// create a quadtree
qtree, _ := zearches.CreateQuadtree(
bounds.NewBound(geo.NewVec3Int(0, 0, 0), geo.NewVec3Int(100, 100, 100)), // bound, required
1, // maxDepth, required
1, // capacity, required
)
qtree.Remove(999)
}
```