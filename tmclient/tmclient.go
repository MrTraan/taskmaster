package main

import (
    "vogsphere.42.fr/taskmaster.git/tmtp"
    "fmt"
    "gopkg.in/readline.v1"
    "strings"
)

var completion = readline.NewPrefixCompleter(
	readline.PcItem("start"),
    readline.PcItem("restart"),
    readline.PcItem("status"),
	readline.PcItem("stop"),
	readline.PcItem("shutdown"),
)

func main() {
    cli, err := tmtp.InitClient("/tmp/tm.sock")
    if err != nil {
        fmt.Println("error on creating client: ", err)
        return
    }
    rl, err := readline.NewEx(&readline.Config{
        Prompt : "tm> ",
        AutoComplete: completion,
    })
    if err != nil {
        fmt.Println("Error on readline init: ", err)
        return
    }
    defer rl.Close()
    
    for {
        line, err := rl.Readline()
        if err != nil {
            fmt.Println("error on readline: ", err)
            return
        }
        if strings.HasPrefix(line, "exit") {
            break
        }
        data, err := tmtp.Encode(line)
        if err != nil {
            fmt.Println("Error on line encoding: ", err)
            continue
        }
        b, err := cli.Write(data)
        if err != nil {
            fmt.Println("Error on sending data: ", err)
            fmt.Println("Check if server is started")
            return
        }
        if b != len(data) {
            fmt.Println("Data sent is corrupted!")
            continue
        }
        fmt.Printf("sent %d bytes to server\n", b)
        response := make([]byte, 512)
        b, err = cli.Read(response)
        if err != nil {
            fmt.Println("Error when reading response: ", err)
        }
        fmt.Printf("got %d bytes from server\nResponse: %s\n", b, string(response))
    }
}