r := new(handlers.Regexp) // github.com/kevinburke/handlers
r.HandleFunc(regexp.MustCompile("^/$"), []string{"GET"}, func(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "text/html; charset=utf-8")
    io.WriteString(w, "<html><body><h1>Hello World</h1></body></html>")
})
