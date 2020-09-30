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
    c := make(chan string)
    n := 0
    for _, web := range sites {
        n++
        go crawl(web.URL, web.Name, c)
    }

    for i := 0; i < n; i++ {
        fmt.Print(<-c)
    }

    fmt.Printf("Total: [%.2fs]\n", time.Since(start).Seconds())
}

func crawl(url, name string, c chan<- string) {
    start := time.Now()
    r, err := http.Get(url)
    if err != nil {
        c <- fmt.Sprintf("error: %s -> %s\n", name, err)
        return
    }
    n,_ := io.Copy(ioutil.Discard, r.Body)
    r.Body.Close()
    dt := time.Since(start).Seconds()
    c <- fmt.Sprintf("%s %d [%.2fs]\n", name, n, dt)
}
