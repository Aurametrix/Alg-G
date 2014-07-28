// initializing arrays
[3]int{1, 2, 3}	
[...]int{1, 2, 3}	
var a [3]int	

//initializing slices
a[:]	slice
var s []int	slice
make([]int, 3)	slice

//The zero value of a slice may be nil, but llen(nil) will return an error 
// unless you specify a type

var slice []int = nil
log.Println(len(slice))

//nil can be a valid parameter or field value and work with append

type Package struct {
    Name         string
    Dependencies []string
}

func main() {
    unsafe := Package{Name: "unsafe", Dependencies: nil}
    log.Println(unsafe, len(unsafe.Dependencies))
    // {unsafe []} 0 
}

var Primes []int
Primes = append(Primes, 2, 3, 5, 7)

//If the slice is the field of a struct

type Node struct {
    Name     string
    Children []Node
}

func Nodes() {
    A := Node{Name: "A"}
    B := Node{Name: "B"}
    A.Children = append(A.Children, B)
}

//named return value

func CreateFibonacci() (Fibonacci []int) {
    Fibonacci = append(Fibonacci, 1, 1, 2, 3, 5, 8)
    return
}

// the value of a map

regions := make(map[string][]string)
regions["Europe"] = append(regions["Europe"], "France", "Germany")
regions["Africa"] = append(regions["Africa"], "Djibouti", "Egypt")
But watch out, the JSON output of a slice will vary depending on the initialization method:

var a []int
jsonA, _ := json.Marshal(a)

b := make([]int, 0)
jsonB, _ := json.Marshal(b)

log.Printf("jsonA: %s\n", jsonA)
log.Printf("jsonB: %s\n", jsonB)
// jsonA: null
// jsonB: []

// arrays are hashable

xy := make(map[[2]int]string)
xy[[2]int{3, 4}] = "Battleship"
log.Println(xy)
// map[[3 4]:Battleship]

// Note: arrays are not valid keys in JSON objects:

output, err := json.Marshal(xy)
if err != nil {
    log.Fatal(err)
}
// json: unsupported type: map[[2]int]string
