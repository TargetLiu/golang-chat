package client

import (
	"bufio"
	"io"
	"net"

	"github.com/TargetLiu/golang-chat/protocol"
	"github.com/TargetLiu/golang-chat/scanline"
	"github.com/sirupsen/logrus"
)

func Start(addr string) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		logrus.Fatalf("dial %s err: %s", addr, err)
	}

	// 输入昵称
	logrus.Infof("Please input your nickname:")
	nickname := scanline.Read()
	logrus.Infof("Hello " + nickname)
	handshakeMsg := protocol.NewMessage()
	handshakeMsg.From = nickname
	handshakeMsg.Content = "handshake"
	handshakeMsg.Type = protocol.HANDSHAKE
	err = handshakeMsg.Write(conn)
	if err != nil {
		logrus.Fatalf("handshake err: %s", err)
	}

	// 开启协程处理消息
	go handle(conn, nickname)

	message := protocol.NewMessage()
	// 发送消息
	for {
		content := scanline.Read()
		message.From = nickname
		message.Content = content
		message.Type = protocol.SAY
		err := message.Write(conn)
		if err != nil {
			logrus.Fatalf("write message err: %s", err)
		}
		message.Reset()
	}
}

func handle(conn net.Conn, nickname string) {
	logrus.Infof("fetch message...")
	reader := bufio.NewReader(conn)
	message := protocol.NewMessage()
	for {
		err := message.Read(reader)
		if err != nil {
			if err == io.EOF {
				logrus.Fatalf("server close the connection")
			}
			logrus.Fatalf("read message err: %s", err)
		}
		if message.From != nickname {
			logrus.Infof("[%s]: %s\n", message.From, message.Content)
		}
		message.Reset()
	}
}
