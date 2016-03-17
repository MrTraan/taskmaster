package main

import (
    "fmt"
    "vogsphere.42.fr/taskmaster.git/tmtp"
)

func main() {
    req := "status lol"
    fmt.Println("Before encoding: ", req)
    data, err := tmtp.Encode(req)
    if err != nil {
        panic(err)
    }
    out, err := tmtp.Decode(data)
    if err != nil {
        panic(err)
    }
    fmt.Println("After decoding: ", out)
}