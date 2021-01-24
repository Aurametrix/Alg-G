Alg-G
=====
Algorithms in Go Language

To install:
https://code.google.com/p/go/downloads/list
or
http://golang.org/doc/install

To learn:
http://golang.org/

[YC recommends](https://news.ycombinator.com/item?id=25812021)

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

GO Resources
https://github.com/mindreframer/golang-stuff/

[Using GO modules](https://blog.golang.org/using-go-modules)

Introdution to GO modules
https://roberto.selbach.ca/intro-to-go-modules/

Book
http://www.golang-book.com/

Go for beginners:
http://kukuruku.co/hub/golang/go-language-for-beginners

[REgular expressions](https://codesalad.dev/blog/go-regular-expressions-5)

Youtube:
https://www.youtube.com/user/gocoding

[Learn Go by writing tests](https://github.com/quii/learn-go-with-tests/tree/master/hello-world) /also for [other languages](http://exercism.io)

[using Go templtes](https://blog.gopheracademy.com/advent-2017/using-go-templates/)

[Scheduling](https://www.ardanlabs.com/blog/2018/08/scheduling-in-go-part2.html)

[Twirp, RPC framework for Go](https://blog.twitch.tv/twirp-a-sweet-new-rpc-framework-for-go-5f2febbf35f)


[Olivia](https://github.com/olivia-ai/olivia), open source [chatbot](https://olivia-ai.org)

[DICOM image parser](https://github.com/suyashkumar/dicom)

[Create system diagrams](https://github.com/blushft/go-diagrams)

a few repos

+ https://github.com/clipperhouse/gen/tree/projection
+ https://github.com/droundy/gotgo
+ https://github.com/jbrukh/bayesian 
+ http://jsgoecke.github.io/go-wit/
+ https://github.com/StefanSchroeder/Golang-Regex-Tutorial
+ [Virtual Kubernetes](https://github.com/ibuildthecloud/k3v)

+ [Parallel Image Processing Algorithms](https://github.com/anthonynsimon/bild)

+ [Automated trading](https://github.com/env0/terratag)

+ http://jozefg.bitbucket.org/posts/2013-08-23-leaving-go.html

+ [Bug tracker embedded in Git](https://github.com/MichaelMure/git-bug)

[Hugo](http://themes.gohugo.io) -  open-source static site generator

[RSS readerNeugram](https://neugram.io/blog/neugram-briefly)

[Ful-text search engine](https://artem.krylysov.com/blog/2020/07/28/lets-build-a-full-text-search-engine/) - [github]((https://github.com/akrylysov/simplefts)

[Uber's highest query per second service](http://eng.uber.com/go-geofence/)

[other data structures for geofencing](https://medium.com/@buckhx/unwinding-uber-s-most-efficient-service-406413c5871d)

[SQL builder and database toolkit](https://github.com/colinjfw/sqlkit)
[web development](https://freshman.tech/web-development-with-go/)

[RSS reader](https://github.com/dertuxmalwieder/rssfs) running as root

[MIDAS](https://github.com/steve0hh/midas)  - edge stream anomaly detection 

[GO+ for data science](https://github.com/qiniu/goplus)

[mmap package](https://github.com/edsrzf/mmap-go); see also blog on [usage](https://brunocalza.me/discovering-and-exploring-mmap-using-go/) & [Bolt](https://brunocalza.me/but-how-exactly-databases-use-mmap/)

[RDF data](https://github.com/owulveryck/gon3) - [ontology & graphs](https://blog.owulveryck.info/2020/11/17/ontology-graphs-and-turtles-part-ii.html)

[Hammer](https://github.com/ShaileshSurya/hammer) - REST API client

[io_uring](https://developers.mattermost.com/blog/hands-on-iouring-go/)

[Go users 2017 survey](https://blog.golang.org/survey2017-results)

[GO developer survey 2019](https://blog.golang.org/survey2019-results)

[11/10/2020](https://blog.golang.org/11years) - 11 years of GO

[Solkit](https://github.com/VladimirMarkelov/solkit) - Solitaire games

[Perkeep](https://github.com/perkeep/perkeep) - peersonal storage systeem



### Versioned Commands

    $ go list -f {{.Dir}} rsc.io/quote
    $ go list -f {{context.ReleaseTags}}
    $ vgo list -f {{.Dir}} rsc.io/quote
    $ vgo list -m
