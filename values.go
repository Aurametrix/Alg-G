// What makes Go fast

// Values: compact data structures, better cache utilization, better performance
var gocon int32 = 2014
// consumes 4 bytes of memory vs 6x in Python (to track type, reference counting)
// vs the need to convert in into integer in Java that consumes 16-24 bytes

// Location is a point in a 3D space

type location struct {
  // 8 bytes per float64
  // 24 bytes in total
  X, Y, Z float64
}


// Location consumes 24 * 1000 bytes
var Locations [1000]Location 
