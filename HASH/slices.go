package main

import "fmt"

/*
example of an array
func main() {
    var x [5]int
    x[4] = 100
    fmt.Println(x)
}
*/

// there are two built-in functions to assist with slices: append and copy

/*
example of append
func main() {
    slice1 := []int{1,2,3}
    slice2 := append(slice1, 4, 5)
    fmt.Println(slice1, slice2)
}
*/

// example of copy
func main() {
    slice1 := []int{1,2,3}
    slice2 := make([]int, 2)
    copy(slice2, slice1)
    fmt.Println(slice1, slice2)
}