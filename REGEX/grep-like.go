// Grep-like searches for regular expressions in text files 
// Every line is matched against the pattern provided on the command line
// if match, it's printed.

package main

import (
    "flag"
    "regexp"
    "bufio"
    "fmt"
    "os"
)

func grep(re, filename string) {
    regex, err := regexp.Compile(re)
    if err != nil {
        return // there was a problem with the regular expression.
    }

    fh, err := os.Open(filename)
    f := bufio.NewReader(fh)

    if err != nil {
        return // there was a problem opening the file.
    }
    defer fh.Close()

    buf := make([]byte, 1024)
    for {
        buf, _ , err = f.ReadLine()
        if err != nil {
            return
        }

        s := string(buf)
        if regex.MatchString(s) {
            fmt.Printf("%s\n", string(buf))
        }
    }
}

func main() {
    flag.Parse()
    if flag.NArg() == 2 {
        grep(flag.Arg(0), flag.Arg(1))
    } else {
        fmt.Printf("Wrong number of arguments.\n")
    }
}
