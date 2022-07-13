# WebSocket

https://github.com/nhooyr/websocket

https://github.com/gorilla/websocket

```go
package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024, // 读取存储空间大小
	WriteBufferSize: 1024, // 写入存储空间大小
	CheckOrigin: func(r *http.Request) bool {
		return true
	}, // 允许跨域
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	// conn类型表示WebSocket连接。服务器应用程序从HTTP请求处理程序调用Upgrader.Upgrade方法以获取* Conn：
	// 完成握手 升级为 WebSocket长连接，使用conn发送和接收消息。
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("upgrade:", err)
		return
	}
	defer conn.Close()

	//调用连接的WriteMessage和ReadMessage方法以一片字节发送和接收消息。实现如何回显消息：
	//p是一个[]字节，messageType是一个值为websocket.BinaryMessage或websocket.TextMessage的int。
	for {
		messageType, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("Reading error...", err)
			return
		}
		log.Printf("Read from client msg:%s \n", msg)

		if err := conn.WriteMessage(messageType, msg); err != nil {
			//if err := conn.WriteMessage(1, []byte("今天。。。"));err != nil {
			log.Println("Writeing error...", err)
			return
		}
		log.Printf("Write msg to client: recved: %s \n", msg)
	}
}

// 在线websocket测试： http://websocket.jsonin.com/
func main() {
	http.HandleFunc("/", wsHandler)
	err := http.ListenAndServe(":8888", nil) // ws://127.0.0.1:8888
	if err != nil {
		log.Fatal("ListenAndServe", err.Error())
	}
}
```