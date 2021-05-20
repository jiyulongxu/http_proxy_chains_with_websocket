package server

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"net/url"
	"strings"

	"golang.org/x/net/websocket"
)

type Server struct {
	Port string `json:"port,omitempty"`

	CertFile string `json:"cert_file,omitempty"`
	KeyFile  string `json:"key_file,omitempty"`
}

func (self *Server) Start() {
	http.Handle("/http_proxy", websocket.Handler(self.HandleClientRequest))
	if err := http.ListenAndServeTLS(":"+self.Port, self.CertFile, self.KeyFile, nil); err != nil {
		panic(err)
	}
}

func (self *Server) HandleClientRequest(client *websocket.Conn) {
	if client == nil {
		log.Println("client is nil")
		return
	}
	defer client.Close()

	var b [2048]byte
	n, err := client.Read(b[:])
	if err != nil {
		log.Println("read err:", err)
		return
	}
	var address string
	arr := strings.Split(string(b[:bytes.IndexByte(b[:], '\n')]), " ")
	if len(arr) != 3 {
		log.Println("len(err) != 3")
		return
	}
	if strings.HasPrefix(arr[1], "http://") {
		if hostPortURL, errParse := url.Parse(arr[1]); errParse != nil {
			address = arr[1]
		} else {
			if (len(arr[1]) > 7) && (!strings.Contains(arr[1][7:], ":")) {
				address = hostPortURL.Host + ":80"
			} else {
				address = hostPortURL.Host
			}
		}
	} else {
		address = arr[1]
	}

	//获得了请求的host和port，就开始连接ws
	server, err := net.Dial("tcp", address)
	if err != nil {
		log.Println(err)
		return
	}
	if arr[0] == "CONNECT" {
		fmt.Fprint(client, strings.TrimSpace(arr[2])+" 200 Connection established\r\n\r\n")
	} else {
		server.Write(b[:n])
	} //进行转发
	go io.Copy(server, client)
	io.Copy(client, server)
}
