package main

import (
	"vogsphere.42.fr/taskmaster.git/tmtp"
    "vogsphere.42.fr/taskmaster.git/tmconf"
    "fmt"
    "gopkg.in/readline.v1"
    "strings"
	"os"
)

var completion = readline.NewPrefixCompleter(
	readline.PcItem("start"),
    readline.PcItem("restart"),
    readline.PcItem("status"),
	readline.PcItem("stop"),
	readline.PcItem("shutdown"),
)

const socketPath = "/tmp/tm.sock"

func main() {
	// A AMELIORER
	if len(os.Args) == 1 {
		fmt.Fprintf(os.Stderr, "requires a config file")
		os.Exit(1)
	}
	_, err := tmconf.GetProcSettings(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "parsing error")
		os.Exit(1)
	}

    cli, err := tmtp.InitClient(socketPath)
    if err != nil {
        fmt.Println("error on creating client: ", err)
        fmt.Println("Make sure taskmaster server is running before launching client")
        return
    }
    defer cli.Close()
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
            fmt.Println("Trying to reload connection")
            cli.Close()
            cli, err = tmtp.InitClient(socketPath)
            if err != nil {
                fmt.Println("Coudlnt reload connection: ", err)
                fmt.Println("Make sure taskmaster server is running")
                return
            }
            fmt.Println("Connection was successfully reloaded")
            continue
        }
        if b != len(data) {
            fmt.Println("Data sent is corrupted!")
            continue
        }
        response := make([]byte, 512)
        b, err = cli.Read(response)
        if err != nil {
            fmt.Println("Error when reading response: ", err)
        }
        fmt.Printf("Response: %s\n", string(response))
    }
}
