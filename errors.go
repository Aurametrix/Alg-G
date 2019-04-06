package main

import "fmt"
import "github.com/pkg/errors"

func main() {
        err := errors.New("error")
        err = errors.Wrap(err, "open failed")
        err = errors.Wrap(err, "read config failed")

        fmt.Println(err) // read config failed: open failed: error
}


ew := &errWriter{w: fd}
ew.write(p0)
ew.write(p1)
ew.write(p2)
// and so on
if ew.err != nil {
    return ew.err
}
