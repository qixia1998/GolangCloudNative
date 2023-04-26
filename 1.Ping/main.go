package main

import (
	"flag"
	"log"
	"net"
	"os"
	"time"
)

var (
	timeout int64
	count   int
	size    int
)

type nickType int64

var age nickType

type ICMP struct {
	Type     uint8
	Code     uint8
	CheckSum uint16
	ID       uint16
	SeqNum   uint16
}

func main() {
	getArgs()
	desIp := os.Args[len(os.Args)-1]
	conn, err := net.DialTimeout("ip:icmp", desIp, time.Duration(timeout)*time.Millisecond)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer conn.Close()
}

func getArgs() {
	flag.Int64Var(&timeout, "w", 1000, "请求超时时间")
	flag.IntVar(&count, "n", 4, "请求次数")
	flag.IntVar(&size, "l", 32, "请求次数")
	flag.Parse()
}
