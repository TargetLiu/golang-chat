package protocol

import (
	"bytes"
	"fmt"
	"io"
	"strconv"
)

// MsgType 消息类型
type Type uint8

const (
	HANDSHAKE Type = iota
	SAY

	MessageHeadLen     = 2
	MessageTypeLen     = 1
	MessageFromSizeLen = 4
	MessageSizeLen     = 8
	MessageTailLen     = 2

	MessageHeadMust = "##"
	MessageTailMust = "##"
)

// Message 消息协议
type Message struct {
	Type    Type
	From    string
	Content string
	Raw     bytes.Buffer
}

func NewMessage() *Message {
	return &Message{}
}

func (m *Message) Read(r io.Reader) error {
	// read head
	head := make([]byte, MessageHeadLen)
	_, err := io.ReadFull(r, head)
	if err != nil {
		return err
	}
	if string(head) != MessageHeadMust {
		return fmt.Errorf("parse head err, read: %s", head)
	}
	m.Raw.Write(head)

	// read msg type
	msgTypeB := make([]byte, MessageTypeLen)
	_, err = io.ReadFull(r, msgTypeB)
	if err != nil {
		return err
	}
	m.Type = Type(msgTypeB[0])
	m.Raw.Write(msgTypeB)

	// read from size
	fromSizeB := make([]byte, MessageFromSizeLen)
	_, err = io.ReadFull(r, fromSizeB)
	if err != nil {
		return err
	}
	fromSize, err := strconv.ParseInt(string(fromSizeB), 10, 32)
	if err != nil {
		return err
	}
	m.Raw.Write(fromSizeB)

	// read from
	fromB := make([]byte, fromSize)
	_, err = io.ReadFull(r, fromB)
	if err != nil {
		return err
	}
	m.From = string(fromB)
	m.Raw.Write(fromB)

	// read msg size
	msgSizeB := make([]byte, MessageSizeLen)
	_, err = io.ReadFull(r, msgSizeB)
	if err != nil {
		return err
	}
	m.Raw.Write(msgSizeB)
	msgSize, err := strconv.ParseInt(string(msgSizeB), 10, 64)
	if err != nil {
		return err
	}

	// read msg
	msgB := make([]byte, msgSize)
	_, err = io.ReadFull(r, msgB)
	if err != nil {
		return err
	}
	m.Content = string(msgB)
	m.Raw.Write(msgB)

	// read tail
	tail := make([]byte, MessageTailLen)
	_, err = io.ReadFull(r, tail)
	if err != nil {
		return err
	}
	if string(tail) != MessageTailMust {
		return fmt.Errorf("parse tail err, read: %s", head)
	}
	m.Raw.Write(tail)

	return nil
}

func (m *Message) Write(w io.Writer) error {
	m.Raw.Write([]byte(MessageHeadMust))
	m.Raw.Write([]byte{byte(m.Type)})
	m.Raw.Write([]byte(fmt.Sprintf("%04d", len(m.From))))
	m.Raw.Write([]byte(m.From))
	m.Raw.Write([]byte(fmt.Sprintf("%08d", len(m.Content))))
	m.Raw.Write([]byte(m.Content))
	m.Raw.Write([]byte(MessageTailMust))
	_, err := w.Write(m.Raw.Bytes())
	m.Raw.Reset()
	return err
}

func (m *Message) Reset() {
	m.Content = ""
	m.From = ""
	m.Type = 0
	m.Raw.Reset()
}
