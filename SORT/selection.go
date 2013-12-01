package main
 
import "fmt"
 
var a = []int{721, 54, 75, -90, -802, 2, 24, 66, 1}
 
func main() {
    fmt.Println("before:", a)
    selectionSort(a)
    fmt.Println("after: ", a)
}
 
func selectionSort(a []int) {
    last := len(a) - 1
    for i := 0; i < last; i++ {
        aMin := a[i]
        iMin := i
        for j := i + 1; j < len(a); j++ {
            if a[j] < aMin {
                aMin = a[j]
                iMin = j
            }
        }
        a[i], a[iMin] = aMin, a[i]
    }
}
