package server

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"strings"
	"sync"

	"github.com/TargetLiu/golang-chat/scanline"

	"github.com/TargetLiu/golang-chat/protocol"
	"github.com/sirupsen/logrus"
)

var (
	conns sync.Map
	chMsg chan *protocol.Message
)

func init() {
	chMsg = make(chan *protocol.Message, 10)
}

func Start(addr string) {
	listen, err := net.Listen("tcp", addr)
	if err != nil {
		logrus.Fatalf("listen err: %s", err)
	}

	logrus.Infof("start listen on: %s", addr)

	//启动协程广播消息
	go broadcast()

	//启动协程处理服务端消息及命令
	go servermsg()

	for {
		//接收客户端连接
		conn, err := listen.Accept()
		if err != nil {
			if opErr, ok := err.(*net.OpError); ok && opErr.Temporary() {
				logrus.Errorf("accept err: %s", opErr)
				continue
			}
			logrus.Fatalf("accept err: %s", err)
		}
		//每个连接放到单独协程进行处理
		go handler(conn)
	}
}

func handler(conn net.Conn) {
	var nickname string
	defer func() {
		conns.Delete(nickname)
		conn.Close()
	}()

	writeMsg := protocol.NewMessage()
	//初次连接发送欢迎消息
	writeMsg.From = "server"
	writeMsg.Content = "Welcome to join."
	writeMsg.Type = protocol.SAY
	writeMsg.Write(conn)
	writeMsg.Reset()

	//读取客户端消息
	reader := bufio.NewReader(conn)
	message := protocol.NewMessage()
	for {
		err := message.Read(reader)
		if err != nil {
			if opErr, ok := err.(*net.OpError); (ok && opErr.Err.Error() == "use of closed network connection") || err == io.EOF {
				return
			}
			logrus.Errorf("read message from %s err: %s", nickname, err)
			return
		}
		logrus.Infof("[%s]: %s\n", message.From, message.Content)
		writeMsg.Type = protocol.SAY
		switch message.Type {
		case protocol.HANDSHAKE:
			//判断是否有同名昵称并保存客户端连接
			nickname = message.From
			if _, ok := conns.LoadOrStore(nickname, conn); ok || nickname == "server" {
				return
			}
			writeMsg.From = "server"
			writeMsg.Content = fmt.Sprintf("[%s] join.", message.From)
		case protocol.SAY:
			writeMsg.From = message.From
			writeMsg.Content = message.Content
		}
		//向通道发送消息
		chMsg <- writeMsg
	}

}

func broadcast() {
	for {
		//从通道中接收消息
		msg := <-chMsg
		//循环客户端连接并发送消息
		conns.Range(func(key interface{}, value interface{}) bool {
			conn := value.(net.Conn)
			err := msg.Write(conn)
			if err != nil {
				conns.Delete(key.(string))
			}
			return true
		})
		msg.Reset()
	}
}

func servermsg() {
	serverMsg := protocol.NewMessage()
	for {
		serverMsg.From = "server"
		serverMsg.Type = protocol.SAY
		//解析命令
		command := scanline.Read()
		cmd := strings.Split(string(command), "|")
		if len(cmd) > 1 {
			switch cmd[0] {
			case "kick":
				if conn, ok := conns.LoadAndDelete(cmd[1]); ok {
					//关闭对应客户端连接
					conn.(net.Conn).Close()
					serverMsg.Content = fmt.Sprintf("kick [%s]", cmd[1])
				}
			default:
				logrus.Errorf("command err")
			}
		} else {
			serverMsg.Content = command
		}
		chMsg <- serverMsg
	}
}
