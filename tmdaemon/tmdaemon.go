package main

import (
	"fmt"
	"os"
	"net"
	"vogsphere.42.fr/taskmaster.git/tmtp"
	"vogsphere.42.fr/taskmaster.git/tmconf"
	"vogsphere.42.fr/taskmaster.git/tmexec"
    "time"
)

func main() {
	// init log file
	logfile, err := CreateLogFile("log.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Logfile: %v\n", err)
		os.Exit(1)
	}
	logfile.Printf("log file created\n")
	if len(os.Args) == 1 {
		fmt.Fprintf(os.Stderr, "requires a config file")
		os.Exit(1)
	}

	// init config file
	conf, err := tmconf.GetProcSettings(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "parsing error")
		os.Exit(1)
	}
	procW, err := tmexec.InitCmd(conf)
	if err != nil {
		fmt.Fprintf(os.Stderr, "cmd error")
		os.Exit(1)
	}
	for i, _ := range procW {
		fmt.Println("launching", procW[i].Name)
		go tmexec.StartCmd(&procW[i])
	}

	// init server connection
	serv, err := tmtp.InitServer("/tmp/tm.sock")
	if err != nil {
		fmt.Println("Coudlnt launch server: ", err)
		return
	}
    defer serv.Close()
	cReq := make(chan tmtp.Request)
    cCon := make(chan net.Conn)
    cResponse := make(chan string)
    go acceptConnections(serv, cCon)
	for {
		select {
        case newConnection := <- cCon :
            go handleConnection(newConnection, cReq, cResponse)
		case req := <-cReq:
			if req.RequestType == tmtp.ReqShutdown {
				fmt.Println("Turning off daemon")
                cResponse <- "You turned me off!"
				logfile.Printf("daemon exited\n")
                time.Sleep(100 * time.Millisecond)
				return
			}
			fmt.Println("Main received request: ", req)
            cResponse <- "Heard you bro"
		}
	}
}

func acceptConnections(serv net.Listener, c chan net.Conn) {
    for {
        conn, err := serv.Accept()
        if err != nil {
		  fmt.Println("Server accept error: ", err)
		  return
        }
        fmt.Println("New client connection")
        c <- conn
    }
}

func handleConnection(conn net.Conn, cIn chan tmtp.Request, cOut chan string){
	for {
		data := make([]byte, 512)
		n, err := conn.Read(data)
        if n == 0 {
            fmt.Println("Connection closed by client")
            return
        }
		if err != nil {
            if err.Error() == "EOF" {
                fmt.Println("Connection closed by client")
                return
            }
			fmt.Println("error on reading data: ", err)
		}
		req, err := tmtp.Decode(data[:n])
		if err != nil {
			fmt.Println("error on decoding data: ", err)
		}
        cIn <- req

        response := []byte(<- cOut)
        n, err = conn.Write(response)
        if err != nil {
            fmt.Println("Error on writing response: ", err)
        }
	}
}
