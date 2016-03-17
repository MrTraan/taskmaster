package tmtp

import (
	"encoding/json"
	"fmt"
	"strings"
    "net"
    "os"
)

const (
	ReqStart   = 0
	ReqStop    = 1
	ReqRestart = 2
	ReqStatus  = 3
    ReqShutdown = 4
)

type Request struct {
	//ContentLength int32
	RequestType int32
	Body        []string
}

func parseRequestType(s string) (int32, error) {
	switch s {
	case "start":
		return 0, nil
		break
	case "stop":
		return 1, nil
		break
	case "restart":
		return 2, nil
		break
	case "status":
		return 3, nil
		break
    case "shutdown":
		return 4, nil
		break
	default:
		return 0, fmt.Errorf("Wrong request type: %s", s)
	}
	return 0, fmt.Errorf("Wrong request type: %s", s)
}

func (r Request) String() string {
	return fmt.Sprintf("Request: %d\nBody: %v",
		r.RequestType, r.Body)
}

func Encode(line string) ([]byte, error) {
	args := strings.Split(line, " ")
	reqType, err := parseRequestType(args[0])
	if err != nil {
		return nil, err
	}
	buf, err := json.Marshal(Request{reqType, args[1:]})
	return buf, err
}

func Decode(rawData []byte) (Request, error) {
	var r Request
	err := json.Unmarshal(rawData, &r)
	return r, err
}

func InitClient(sPath string) (net.Conn, error) {
    return net.Dial("unix", sPath)
}

func InitServer(sPath string) (net.Listener, error) {
    return net.Listen("unix", sPath)
}

func CloseServer(sPath string) error {
    return os.Remove(sPath)
}
