package main

import (
    "fmt"
    "time"
    "io"
    "net/http"
    "io/ioutil"
)

type Web struct {
    Name string
    URL string
}

var sites = []Web {
    { "Go", "https://golang.org/" },
    { "Python", "https://python.org/" },
    { "C#", "https://www.java.com/" },
    { "Ruby", "https://www.ruby-lang.org/" },
    { "Erlang", "https://www.erlang.org/" },
}

func main() {
    start := time.Now()
    for _, web := range sites {
        crawl(web.URL, web.Name)
    }
    fmt.Printf("Total: [%.2fs]\n", time.Since(start).Seconds())
}

// START OMIT
func crawl(url, name string) {
    start := time.Now()
    r, err := http.Get(url)
    if err != nil {
        fmt.Printf("error: %s -> %s\n", name, err)
        return
    }
    n, _ := io.Copy(ioutil.Discard, r.Body)
    r.Body.Close()
    fmt.Printf("%s %d [%.2fs]\n", name, n, time.Since(start).Seconds())
}
// END OMIT
