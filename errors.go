package main

import "fmt"
import "github.com/pkg/errors"

func main() {
        err := errors.New("error")
        err = errors.Wrap(err, "open failed")
        err = errors.Wrap(err, "read config failed")

        fmt.Println(err) // read config failed: open failed: error
}


b := bufio.NewWriter(fd)
b.Write(p0)
b.Write(p1)
b.Write(p2)
// and so on
if b.Flush() != nil {
    return b.Flush()
}

r := fd.Write(p0).Then(func(_) {
		return fd.Write(p1)
	}).Then(func(_) {
		return fd.Write(p2)
	})
// and so on
if r.Error() != nil {
	return r.Error()
}
