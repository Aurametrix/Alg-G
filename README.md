Alg-G
=====
Algorithms in Go Language

To install:
https://code.google.com/p/go/downloads/list
or
http://golang.org/doc/install

To learn:
http://golang.org/

[Programming blueprints](https://mtlynch.io/book-reports/go-programming-blueprints/)

[GO is not an easy language](https://www.arp242.net/go-easy.html)

[YC recommends](https://news.ycombinator.com/item?id=25812021)

[Fast and reliable background jobs in Go](https://riverqueue.com/)

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

How to start a GO project in 2023
https://boyter.org/posts/how-to-start-go-project-2023/

Another Go playground powered by WASM
https://go-playground-wasm.vercel.app/

Better Structured Concurency
https://github.com/sourcegraph/conc

Machine Learning Library:
https://github.com/sjwhitworth/golearn

Higher order functions
https://eli.thegreenplace.net/2023/higher-order-functions-in-go/

A library for bringing generics-like functionality to Go
http://clipperhouse.github.io/gen/

Go Object Oriented Design
http://nathany.com/good/

Go google group
https://groups.google.com/forum/#!forum/go-qml

Go for data science
https://goplus.org/

Static Site Generator built in GoLang
https://github.com/spf13/hugo

GO Resources
https://github.com/mindreframer/golang-stuff/

[Golang news](https://golangnews.com/)

[Using GO modules](https://blog.golang.org/using-go-modules)

Introdution to GO modules
https://roberto.selbach.ca/intro-to-go-modules/

Book
http://www.golang-book.com/

Go for beginners:
http://kukuruku.co/hub/golang/go-language-for-beginners

[Regular expressions](https://codesalad.dev/blog/go-regular-expressions-5)

[Banging errors](https://flak.tedunangst.com/post/bango)

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

+ [Build Your Own Search Engine](https://wiby.me/about/guide.html)  - [repo](https://github.com/wibyweb/wiby/)
+ [An HTTP Server in Go From scratch](https://www.krayorn.com/posts/http-server-go-2/)
+ [OpenSERP - API for Search Engine](https://github.com/karust/openserp)
+ https://github.com/clipperhouse/gen/tree/projection
+ https://github.com/droundy/gotgo
+ https://github.com/jbrukh/bayesian 
+ http://jsgoecke.github.io/go-wit/
+ https://github.com/StefanSchroeder/Golang-Regex-Tutorial
+ [Virtual Kubernetes](https://github.com/ibuildthecloud/k3v)
+ [Weatherme CLI Tool](https://github.com/eoin-barr/weatherme)
+ [Polaris]([weight workflow orchestrator for Golang](https://github.com/harshadmanglani/polaris)) - extremely lightweight workflow orchestrator for Golang
+ [GO optimization guide](https://goperf.dev/)
+ [ClipCapsule](https://github.com/Victor-Evogor/clipcapsule) -  a minimalist clipboard manager for Linux, built with Go and WailsJS. It supercharges your productivity by allowing you to manage and switch clipboard entries using only keyboard shortcuts—no mouse or GUI required.

+ [Time series database](https://github.com/nakabonne/tstorage); [blog post](https://nakabonne.dev/posts/write-tsdb-from-scratch/)

+ [Parallel Image Processing Algorithms](https://github.com/anthonynsimon/bild)

+ [Automated trading](https://github.com/env0/terratag)

+ http://jozefg.bitbucket.org/posts/2013-08-23-leaving-go.html
+ [NBFM](https://github.com/notebox/nbfm) - file manager + CRDT Note Editor

+ [Bug tracker embedded in Git](https://github.com/MichaelMure/git-bug)
+ [Togomak](https://github.com/srevinsaju/togomak) - pipeline orchestrator

[Hugo](http://themes.gohugo.io) -  open-source static site generator

[RSS readerNeugram](https://neugram.io/blog/neugram-briefly)

[Purl](https://github.com/catatsuy/purl) - practical alternative to traditional tools like sed and grep, designed to address some of their common limitations

[Ful-text search engine](https://artem.krylysov.com/blog/2020/07/28/lets-build-a-full-text-search-engine/) - [github]((https://github.com/akrylysov/simplefts)

[Uber's highest query per second service](http://eng.uber.com/go-geofence/)

[A Study of Real-World Data Races in Golang](https://dl.acm.org/doi/pdf/10.1145/3519939.3523720)

[other data structures for geofencing](https://medium.com/@buckhx/unwinding-uber-s-most-efficient-service-406413c5871d)

[More powerful Go execution traces](https://go.dev/blog/execution-traces-2024)

[SQL builder and database toolkit](https://github.com/colinjfw/sqlkit)

[FeretDB](https://github.com/FerretDB/FerretDB) - MongoDB Protocol for SQLite

[GO Git](https://benhoyt.com/writings/gogit/) - a 400-line Git client that can create a repo and push itself to GitHub

[Layered design in Go](https://jerf.org/iri/post/2025/go_layered_design/)

[web development](https://freshman.tech/web-development-with-go/)

[Pushup](https://github.com/adhocteam/pushup) - page-oriented web framework for Go

[RSS reader](https://github.com/dertuxmalwieder/rssfs) running as root

[Ycombinator in Go](https://eli.thegreenplace.net/2022/the-y-combinator-in-go-with-generics/)

[MIDAS](https://github.com/steve0hh/midas)  - edge stream anomaly detection 

[GO+ for data science](https://github.com/qiniu/goplus)

[mmap package](https://github.com/edsrzf/mmap-go); see also blog on [usage](https://brunocalza.me/discovering-and-exploring-mmap-using-go/) & [Bolt](https://brunocalza.me/but-how-exactly-databases-use-mmap/)

[RDF data](https://github.com/owulveryck/gon3) - [ontology & graphs](https://blog.owulveryck.info/2020/11/17/ontology-graphs-and-turtles-part-ii.html)

[Hammer](https://github.com/ShaileshSurya/hammer) - REST API client

[io_uring](https://developers.mattermost.com/blog/hands-on-iouring-go/)

[Connet](https://github.com/connet-dev/connet) - peer-to-peer reverse proxy for NAT traversal

[Go users 2017 survey](https://blog.golang.org/survey2017-results)

[GO developer survey 2019](https://blog.golang.org/survey2019-results)

[Data race patterns](https://eng.uber.com/data-race-patterns-in-go/)

[From Python to Go](https://new.pythonforengineers.com/blog/learning-go-as-a-python-developer-the-good-the-bad-and-the-ugly/)

[A Guide to the Go Garbage Collector](https://tip.golang.org/doc/gc-guide)

[11/10/2020](https://blog.golang.org/11years) - 11 years of GO

[Faster Interpreters in Go](https://planetscale.com/blog/faster-interpreters-in-go-catching-up-with-cpp)

[Weather MPC server](https://github.com/TuanKiri/weather-mcp-server) - A lightweight Model Context Protocol (MCP) server that enables AI assistants like Claude to retrieve and interpret real-time weather data.

[1.1.9 beta](https://groups.google.com/g/golang-announce/c/SNruPJUSFz0); [release notes](https://tip.golang.org/doc/go1.19)
[1.19](https://go.dev/doc/go1.19)

[errtrace](https://github.com/bracesdev/errtrace) - trace error's return path

[graceful shutdown in go](https://victoriametrics.com/blog/go-graceful-shutdown/index.html)

[Google Maps Scraper](https://github.com/gosom/google-maps-scraper)

[Solkit](https://github.com/VladimirMarkelov/solkit) - Solitaire games

[Perkeep](https://github.com/perkeep/perkeep) - peersonal storage systeem

[Gokrazy](https://gokrazy.org/) - deploy your Go programs as appliances to a Raspberry Pi or PC

[PCP](https://github.com/dennis-tra/pcp) - peer-to-peer data transfr tool

[Command PATH security](https://blog.golang.org/path-security)

[Tilt](https://github.com/tilt-dev/tilt) - dev environment as code

[The universal loader](https://www.symbolcrash.com/2021/03/04/the-universal-loader-for-go/)

[Scripting with Go](https://bitfieldconsulting.com/golang/scripting)

[web scraping](https://www.scrapingbee.com/blog/web-scraping-go/)

[custom code search index](https://boyter.org/posts/how-i-built-my-own-index-for-searchcode/)

[Lmgrp](https://github.com/boyter/cs) -- [lucene grep](https://www.jocas.lt/blog/post/intro-to-lucene-grep/)

[Streaming Platform for Postgres](https://blog.peerdb.io/building-a-streaming-platform-in-go-for-postgres)

[Constraints in GO](https://bitfieldconsulting.com/posts/constraints)

[ML in GO](https://eli.thegreenplace.net/2024/ml-in-go-with-a-python-sidecar/)

[tcpulse](https://github.com/yuuki/tcpulse) - concurrent TCP/UDP load generator wthat provides fine-grained, flow-level control over connection establishment and data transfer.

[mcwig](https://github.com/firstrow/mcwig)   - editor

[SPOT](https://github.com/roblillack/spot) - cross-platform, reactive GUI toolkit for Go using native widgets where available.

[Bruin](https://github.com/bruin-data/bruin) - data pipeline tool that brings together data ingestion, data transformation with SQL & Python, and data quality into a single framework.

[Physics-GO](https://github.com/rudransh61/Physix-go) - A Simple Physics Engine in GoLang

[The Evolution of Caching Libraries](https://maypok86.github.io/otter/blog/cache-evolution/#on-heap-vs-off-heap)

[Goravel](https://www.goravel.dev/) - A Go framework inspired by Laravel

[event dispatcher](https://github.com/kelindar/event)

[Fasten](https://github.com/fastenhealth/fasten-onprem) - backend in go, front end in .ts - securely connects your healthcare providers together, creating a personal health record that never leaves your hands

e-paper gadgets. [Beepberry](https://github.com/beeper/beepberry) - a portable e-paper computer for hackers, designed for chatting on Beeper.
    [Beeper](https://beeper.com)
    [GOmuks](https://github.com/tulir/gomuks) - 
    [Maubot](https://github.com/maubot/maubot) - 
    
### Versioned Commands

    $ go list -f {{.Dir}} rsc.io/quote
    $ go list -f {{context.ReleaseTags}}
    $ vgo list -f {{.Dir}} rsc.io/quote
    $ vgo list -m
