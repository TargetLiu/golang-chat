// chat project main.go
package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Please input:\"chat [server|client] [:port|IP address:port]\"")
		os.Exit(-1)
	}

	if os.Args[1] == "server" {
		startServer(os.Args[2])
	} else if os.Args[1] == "client" {
		startClient(os.Args[2])
	} else {
		fmt.Println("Wrong param")
		os.Exit(-1)
	}
	fmt.Println(os.Args[1])
	fmt.Println("Hello World!")
}

//ScanLine 读取整行，不以空格为分隔符
func ScanLine() string {
	var c byte
	var err error
	var b []byte
	for err == nil {
		_, err = fmt.Scanf("%c", &c)

		if c != '\n' {
			b = append(b, c)
		} else {
			break
		}
	}

	return string(b)
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
