package main

import (
    "fmt"
    "writefile/writetxt"
    "os"
)

func main() {
    f, err := os.Create("file.txt")
    if err != nil {
        fmt.Fprintf(os.Stderr, "create: %s", err)
        os.Exit(1)
    }
    defer f.Close()

    n, err := writetxt.Copy(f, os.Stdin)
    if err != nil {
        fmt.Fprintf(os.Stderr, "write: %s", err)
        os.Exit(1)
    }
    fmt.Println(n)
}
