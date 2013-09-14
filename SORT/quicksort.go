package main

import (
    "fmt"
    "math/rand"
    "time"
)

type Item int
type Items []Item

// Algorithms and Data Structures, N. Wirth
// http://www.ethoberon.ethz.ch/WirthPubl/AD.pdf
// 2.3.3 Partition Sort: Quicksort, NonRecursiveQuickSort.
func qSort(a Items) {
    const M = 12
    var i, j, l, r int
    var x Item
    var low, high = make([]int, 0, M), make([]int, 0, M)

    low, high = append(low, 0), append(high, len(a)-1)
    for { // (*take top request from stack*)
        l, low = low[len(low)-1], low[:len(low)-1]
        r, high = high[len(high)-1], high[:len(high)-1]
        for { // (*partition a[l] ... a[r]*)
            i, j = l, r
            x = a[l+(r-l)/2]
            for {
                for ; a[i] < x; i++ {
                }
                for ; x < a[j]; j-- {
                }
                if i <= j {
                    a[i], a[j] = a[j], a[i]
                    i++
                    j--
                }
                if i > j {
                    break
                }
            }
            if i < r { // (*stack request to sort right partition*)
                low, high = append(low, i), append(high, r)
            }
            r = j // (*now l and r delimit the left partition*)
            if l >= r {
                break
            }
        }
        if len(low)+len(high) == 0 {
            break
        }
    }
}

func main() {
    nItems := 5
    //nItems := 4096
    a := make(Items, nItems)
    rand.Seed(time.Now().UnixNano())
    for i := range a {
        a[i] = Item(rand.Int31())
    }

    fmt.Println("unsorted:", a)

    qSort(a)
    fmt.Println("sorted:", a)
    for i := range a[1:] {
        if a[i] > a[i+1] {
            fmt.Println("(* sort error *)")
        }
    }



}
