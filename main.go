package main

import (
	"fmt"
	"go-realtime-dashboard/data"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true }, // dev用
}

func main() {
	r := gin.Default()

	r.Static("/static", "./static")

	r.GET("/ws", func(c *gin.Context) {
		conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			fmt.Println("WebSocket upgrade error:", err)
			return
		}
		defer conn.Close()

		for {
			dp := data.GenerateData()
			err := conn.WriteJSON(dp)
			if err != nil {
				fmt.Println("WebSocket write error:", err)
				break
			}
			time.Sleep(1 * time.Second) // 1秒ごとに送信
		}
	})

	r.Run(":8000")
}
