package main

import "time"
import "fmt"

//The first timer expires ~2s after the program is started
//The second is stopped before it has a chance to expire.

func main() {

        timer1 := time.NewTimer(time.Second * 2)

        <-timer1.C
        fmt.Println("1st timer expired")

        timer2 := time.NewTimer(time.Second)

        go func() {
                <-timer2.C
                fmt.Println("Second timer expired")
        }()
        stop2 := timer2.Stop()
        if stop2 {
                fmt.Println("Second timer stopped")
        }
}
