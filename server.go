package main

import (
	"net"
	"github.com/shellus/pkg/logs"
	"encoding/binary"
)

func main() {
	l, err := net.Listen("tcp", ":8001")
	if err != nil {
		logs.Fatal(err)
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			logs.Fatal(err)
		}

		logs.Info("onConnect %s", conn.RemoteAddr().String())
		go onConnect(conn)
	}
}

func onConnect(conn net.Conn){
	defer conn.Close()
	defer logs.Info("onClose %s", conn.RemoteAddr().String())

	for {
		var sz int64
		err := binary.Read(conn, binary.LittleEndian, &sz)
		if err != nil {
			logs.Error("read header err: %s", err)
			return
		}
		if sz > 1024 * 1024 {
			logs.Error("header body length %d > 1024 * 1024", sz)
			return
		}

		buffer := make([]byte, sz)
		n, err := conn.Read(buffer)
		if err != nil {
			logs.Error("read err: %s", err)
			return
		}

		if int64(n) != sz {
			logs.Error("Expected to read %d bytes, but only read %d", sz, n)
			return
		}

		onRecv(conn, buffer)
	}

}

func onRecv(conn net.Conn, buf []byte){
	logs.Info("recv %d byte from %s", len(buf), conn.RemoteAddr().String())
}