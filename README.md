Alg-G
=====
Algorithms in Go Language

To install:
https://code.google.com/p/go/downloads/list
or
http://golang.org/doc/install

To learn:
http://golang.org/


To run

go run *.go


=====================
Usage:

go command [arguments]

The commands are:

build       compile packages and dependencies

clean       remove object files

doc         run godoc on package sources

env         print Go environment information

fix         run go tool fix on packages

fmt         run gofmt on package sources

get         download and install packages and dependencies

install     compile and install packages and dependencies

list        list packages

run         compile and run Go program

test        test packages

tool        run specified go tool

version     print Go version

vet         run go tool vet on packages

Use "go help [command]" for more information about a command.

======

Examples of syntax:
a + b => a.Add(b)
b*b - 4 * a * c => b.Mul(b).Sub(big.NewInt(4).Mul(a).Mul(c))
Go is a good choice for Python and Ruby developers for easier higher concurrency,
trading some flexibility, ambiguousness, or dynamism for better performance & safer code.
Go is more productive and maintainable than Lisp, but for many it's less "fun" 

Go’s structs and pointers are very similar to C, the main difference is that pointer arithmetic is illegal. 

    type Mat4 [16]float32

    type Entity struct {
      Id        uint32
      Transform Mat4
      Mesh      *gfx.Mesh
    }



Go Playground
http://play.golang.org/

Machine Learning Library:
https://github.com/sjwhitworth/golearn

A library for bringing generics-like functionality to Go
http://clipperhouse.github.io/gen/

Go Object Oriented Design
http://nathany.com/good/

Go google group
https://groups.google.com/forum/#!forum/go-qml

Static Site Generator built in GoLang
https://github.com/spf13/hugo

Go Resources
https://github.com/mindreframer/golang-stuff/

Book
http://www.golang-book.com/

Go for beginners:
http://kukuruku.co/hub/golang/go-language-for-beginners

Youtube
https://www.youtube.com/user/gocoding


[using Go templtes](https://blog.gopheracademy.com/advent-2017/using-go-templates/)

a few repos
https://github.com/clipperhouse/gen/tree/projection
https://github.com/droundy/gotgo
https://github.com/jbrukh/bayesian 
http://jsgoecke.github.io/go-wit/
https://github.com/StefanSchroeder/Golang-Regex-Tutorial

http://jozefg.bitbucket.org/posts/2013-08-23-leaving-go.html


[Hugo](http://themes.gohugo.io) -  open-source static site generator

[Neugram](https://neugram.io/blog/neugram-briefly)
