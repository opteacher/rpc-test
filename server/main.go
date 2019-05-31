package main

import (
	"fmt"
	"net"
	"log"
	"bufio"
	"encoding/json"

	"rpcserver/service"
	"rpcserver/domain"
)

var svcDsc *service.Discovery

func handleConnection(conn net.Conn) {
	log.Printf("收到来自%s的连接请求\n", conn.RemoteAddr().String())
	if data, err := bufio.NewReader(conn).ReadBytes('\n'); err != nil {
		conn.Write(wrapError("读取数据发生错误：%v", err))
	} else if req, err := parseReqParams(data); err != nil {
		conn.Write(wrapError("请求数据非Reqs格式：%v", err))
	} else if resp, err := svcDsc.CallMethod(req.Method, req.Params); err != nil {
		conn.Write(wrapError("调用服务错误：%v", err))
	} else {
		conn.Write(resp)
	}
}

func parseReqParams(data []byte) (domain.Reqs, error) {
	var req domain.Reqs
	log.Println("收到请求：", string(data))
	err := json.Unmarshal(data, &req)
	return req, err
	
}

func wrapError(fstr string, params ...interface{}) []byte {
	message := fmt.Sprintf(fstr, params...)
	log.Println(message)
	if res, err := json.Marshal(domain.Resp{
		Status: 400,
		Message: message,
	}); err != nil {
		res, _ = json.Marshal(domain.Resp{
			Status: 400,
			Message: fmt.Sprintf("打包错误失败：%v", err),
		})
		return res
	} else {
		return res
	}
}

func main()  {
	if ds, err := service.NewDiscovery(); err != nil {
		log.Fatalf("初始化服务发现者失败：%v", err)
	} else {
		svcDsc = ds
	}
	if ln, err := net.Listen("tcp", ":21700"); err != nil {
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