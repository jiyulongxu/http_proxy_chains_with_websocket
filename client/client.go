package client

import (
	"crypto/tls"
	"fmt"
	"io"
	"net"

	"golang.org/x/net/websocket"
)

type Client struct {
	OutWsAddr string `json:"out_ws_addr,omitempty"`
	OutAddr   string `json:"out_addr,omitempty"`

	Port string `json:"port,omitempty"`
}

func (self *Client) Start() {
	l, err := net.Listen("tcp", ":"+self.Port)
	if err != nil {
		panic(err)
	}
	for {
		client, err := l.Accept()
		if err != nil {
			panic(err)
		}
		go self.HandleClientRequest(client)
	}
}

func (self *Client) HandleClientRequest(client net.Conn) {
	if client == nil {
		return
	}
	var config *websocket.Config
	var err error
	if config, err = websocket.NewConfig(self.OutWsAddr, self.OutAddr); err != nil {
		fmt.Println("ws new err:", err)
		return
	}
	config.TlsConfig = &tls.Config{
		InsecureSkipVerify: true,
	}
	var server *websocket.Conn
	if server, err = websocket.DialConfig(config); err != nil {
		fmt.Println("dial err:", err)
		return
	}
	go self.CopyIO(client, server)
	go self.CopyIO(server, client)
}

func (self *Client) CopyIO(src, dest io.ReadWriteCloser) {
	defer src.Close()
	defer dest.Close()
	io.Copy(src, dest)
}
