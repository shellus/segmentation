package main

import (
	"net"
	"github.com/shellus/pkg/logs"
	"crypto/rand"
	"encoding/binary"
	"time"
)

func main(){
	conn, err := net.Dial("tcp", "127.0.0.1:8001")
	if err != nil {
		logs.Fatal(err)
	}
	buf := make([]byte, 1024*10)
	n, err := rand.Read(buf)
	if err != nil {
		logs.Fatal(err)
	}
	if n == 0 {
		logs.Fatal("rand.Read 0")
	}

	err = binary.Write(conn, binary.LittleEndian, int64(n))
	if err != nil {
		logs.Fatal(err)
	}
	n, err = conn.Write(buf)
	if err != nil {
		logs.Fatal(err)
	}
	logs.Info(n)
	time.Sleep(time.Second)
}
