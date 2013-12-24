/*
    Program reads text from 
    standard-input, and returns (prints) the total 
    number of distinct words found in the text. 
    http://unthought.net/c++/c_vs_c++.html
*/

package main

import (
    "bufio"
    "fmt"
    "os"
    "unicode"
)

func main() {
    rdr := bufio.NewReader(os.Stdin)
    words := make(map[string]bool, 1024*1024)
    word := make([]int, 0, 256)
    for {
        r, n, _ := rdr.ReadRune()
        if unicode.IsSpace(r) || n == 0 {
            if len(word) > 0 {
                words[string(word)] = true
                word = word[:0]
            }
            if n == 0 {
                break
            }
        } else {
            word = append(word, r)
        }
    }
    fmt.Println(len(words))
}

