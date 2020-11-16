// chat project main.go
package main

import (
	"flag"
	"fmt"

	"github.com/TargetLiu/golang-chat/client"
	"github.com/TargetLiu/golang-chat/server"
	"github.com/sirupsen/logrus"
)

var (
	startType string
	host      string
	port      int
)

func main() {
	flag.StringVar(&startType, "type", "server", "start type [server|client]")
	flag.StringVar(&host, "host", "127.0.0.1", "host")
	flag.IntVar(&port, "port", 6666, "port")
	flag.Parse()

	addr := fmt.Sprintf("%s:%d", host, port)
	switch startType {
	case "server":
		server.Start(addr)
	case "client":
		client.Start(addr)
	default:
		logrus.Fatalf("wrong param")
	}
}
