package main

import (
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"go.uber.org/zap"
)

var logger *zap.Logger

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func main() {
	zc := zap.NewProductionConfig()
	var err error
	logger, err = zc.Build()
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		var ID = getID()
		header := make(http.Header)
		header.Add("ID", ID)
		conn, err := upgrader.Upgrade(w, r, header)
		defer func() {
			err = conn.Close()
			if err != nil {
				logger.Error("websocket connection close error", zap.Error(err))
			}
			err = r.Body.Close()
			if err != nil {
				logger.Error("request close error", zap.Error(err))
			}
		}()
		if err != nil {
			logger.Error("upgrade error", zap.Error(err))

		}

		processWebsocketConnection(conn, ID)
	})
	logger.Info("server start", zap.String("addr", ":1234"))
	logger.Sugar().Fatal(http.ListenAndServe(":1234", nil))
}

func processWebsocketConnection(conn *websocket.Conn, ID string) {
	done := make(chan struct{})
	go func() {
		for {
			_, bs, err := conn.ReadMessage()
			if err != nil {
				logger.Error("read message error", zap.Error(err))
				done <- struct{}{}
				return
			}
			logger.Info("receive", zap.String("msg", string(bs)))
		}
	}()

	t := time.NewTicker(time.Second)
	for {
		select {
		case <-done:
			return
		case <-t.C:
			err := conn.WriteMessage(1, []byte("你好"))
			if err != nil {
				logger.Error("write message error", zap.Error(err))
			}
		}
	}
}

var idMutex sync.Mutex
var id int

func getID() string {
	idMutex.Lock()
	defer idMutex.Unlock()
	id = id + 1
	return strconv.Itoa(id)
}
