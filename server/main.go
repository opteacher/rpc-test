package main

import (
	"net"
	"log"
	"bufio"

	"rpcserver/service"
)

func handleConnection(conn net.Conn) {
	log.Printf("收到来自%s的连接请求\n", conn.RemoteAddr().String())
	if data, err := bufio.NewReader(conn).ReadString('\n'); err != nil {
		log.Fatalf("读取数据发生错误：%v", err)
	} else {
		log.Println(data)
		conn.Write([]byte("欢迎光临！"))
	}
}

func main()  {
	if _, err := service.NewDiscovery(); err != nil {
		log.Fatalf("初始化服务发现者失败：%v", err)
	} else if ln, err := net.Listen("tcp", ":21700"); err != nil {
		log.Fatalf("监听端口异常：%v", err)
	} else {
		log.Printf("端口%s监听中...\n", ln.Addr().String())
		for {
			if conn, err := ln.Accept(); err != nil {
				log.Fatalf("传输数据发生错误：%v", err)
			} else {
				go handleConnection(conn)
			}
		}
	}
}