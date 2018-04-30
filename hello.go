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


  t := "Hello world!"
  c := make(chan string, 32)
  for i := 0; i < 32; i++ {
    go gopher(t, c)
  }
  for s := range c {
    fmt.Println(s)
    if s == t {
      break
    }
  }
}

func gopher(t string, c chan string) {
  s := []rune(t)
  for {
    rand.Shuffle(len(s), func(i, j int) {
      s[i], s[j] = s[j], s[i]
    })
    c <- string(s)
  }
}

