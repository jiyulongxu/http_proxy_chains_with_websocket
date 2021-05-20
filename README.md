# http_proxy_chains_with_websocket
基于websocket连接的http代理链路工具。


使用方式：

1、生成server.crt和server.key：

make openssl_key

2、在linux环境下编译指定目标环境的可执行程序

3、配置server.json之后在服务器上执行


# server.json
```json
{
	"cert_file": "./bin/server.crt",
	"key_file": "./bin/server.key",
	"port": "12345"
}
```

在服务器上执行：

./server ./server.json

4、配置client.json之后在本地执行，可以放在路由器上执行。

假设服务器工作在192.168.11.1上，监听的端口是12345。

# client.json
```json
{
	"out_addr": "https://192.168.11.1:12345",
	"out_ws_addr": "wss://192.168.11.1:12345/http_proxy",
	"port": "55555"
}
```

在本机或者局域网某台设备，或者路由器上执行。

./client ./client.json

client也会监听一个端口，55555。假设client的ip是192.168.11.5

局域网内的http代理https代理配置成，主机名：192.168.11.5，端口：55555,就可以使用代理上网了。
