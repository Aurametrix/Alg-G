package main

import (
    "fmt"
    "time"
    "math/rand"
    )

var x, y, z = true, "no!", "what's your name?"

func add(x int, y int) int {
    return x + y
}

func main() {
    const You = "世界"
    var i int
    var name string
    
    fmt.Printf("hello, world\n")
    fmt.Println("Are you", You, x, y, z)
    fmt.Scan(&name );
    fmt.Println("Hello", name)
    
    fmt.Println("The time is", time.Now())
    fmt.Println("My favorite number is", add(4,rand.Intn(10)))
    fmt.Println("yours?")
    fmt.Scan(&i)
    fmt.Println("Your favorite number is", i, "!")
}
