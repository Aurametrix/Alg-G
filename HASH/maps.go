package main

import "fmt"

/*
Like arrays and slices maps can be accessed using brackets

x := make(map[string]int)
x["key"] = 10
fmt.Println(x["key"])

OR

x := make(map[int]int)
x[1] = 10
fmt.Println(x[1])

*/

// an example program that uses a map

func main() {
    elements := map[string]map[string]string{
        "H": map[string]string{
            "name":"Hydrogen", 
            "state":"gas",
        },
        "He": map[string]string{
            "name":"Helium", 
            "state":"gas",
        },
        "Li": map[string]string{
            "name":"Lithium", 
            "state":"solid",
        },
        "Be": map[string]string{
            "name":"Beryllium", 
            "state":"solid",
        },
        "B":  map[string]string{
            "name":"Boron",
            "state":"solid",
        },
        "C":  map[string]string{
            "name":"Carbon",
            "state":"solid",
        },
        "N":  map[string]string{
            "name":"Nitrogen",
            "state":"gas",
        },
        "O":  map[string]string{
            "name":"Oxygen",
            "state":"gas",
        },
        "F":  map[string]string{
            "name":"Fluorine",
            "state":"gas",
        },
        "Ne":  map[string]string{
            "name":"Neon",
            "state":"gas",
        },
    }

    if el, ok := elements["Li"]; ok {    
        fmt.Println(el["name"], el["state"])
    }
}

