package main

import (
	"crypto/tls"
	"fmt"

	"golang.org/x/net/websocket"
)

func main() {
	var config *websocket.Config
	var err error
	fmt.Println("begin ws:")
	if config, err = websocket.NewConfig("wss://127.0.0.1:12345/http_proxy", "https://127.0.0.1:12345"); err != nil {
		fmt.Println(err)
		return
	}
	config.TlsConfig = &tls.Config{
		InsecureSkipVerify: true,
	}
	var server *websocket.Conn
	if server, err = websocket.DialConfig(config); err != nil {
		fmt.Println("DialConfig:", err)
		return
	}
	fmt.Println("conn ws ok:")
	server.Write([]byte(`CONNECT server.example.com:80 HTTP/1.1
Host: server.example.com:80
Proxy-Authorization: basic aGVsbG86d29ybGQ=
`))
}
