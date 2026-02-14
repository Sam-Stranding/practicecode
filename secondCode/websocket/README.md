# WebSocket(gorilla/websocket库)

**WebSocket特点：**

1. 一次握手，长久连接：客户端发送WebSocket请求，服务端同意后建立长久连接。

2. 双向通信：服务端可以主动向客户端推送数据，不需要客户端轮询

3. 低延迟：避免HTTP每次请求的头部开销和连接建立时间

4. 协议升级:由HTTP协议升级而来

**编程实例**

```go
package main

import (
    "fmt"
    "net/http"

    "github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
    CheckOrigin: func(r *http.Request) bool {
        return true
    },
}

func main() {
    http.HandleFunc("/ws", wsHandler)

    http.ListenAndServe(":8080", nil)
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
    //客户端发送websocket请求
    //http升级为WebSocket
    conn, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        http.Error(w, "Failed to upgrade to websocket", http.StatusBadRequest)
        fmt.Println("Upgrader failed")
        return
    }
    defer conn.Close()

    fmt.Println("New WebSocket: 连接成功")

    //持续监听
    for {
         //读取客户端发送信息
        messageType, message, err := conn.ReadMessage()
        if err != nil {
            fmt.Println("ReadMessage failed")
            return
        }
        //向客户端返回信息
        fmt.Println("Received messageType:", messageType)
        err = conn.WriteMessage(messageType, message)
        message = append(message, "已收到"...)
        if err != nil {
            fmt.Println("WriteMessage failed")
            return
        }
    }
    fmt.Println("连接已关闭")
}

```

**websocket流程总结**

1.连接建立

```go
//客户端发送websocket请求
//http升级为WebSocket
conn, err := upgrader.Upgrade(w, r, nil)
```

2.持续监听和处理

```go
 //读取客户端发送信息
 messageType, message, err := conn.ReadMessage()
 
 //处理逻辑
 ...
 
 //向客户端返回信息
 conn.WriteMessage(messageType, message)
```

3.关闭conn

```go
defer conn.Close()
```


