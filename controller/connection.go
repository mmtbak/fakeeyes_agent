package controller

import "github.com/gorilla/websocket"

func Connect() (err error) {
	conn, err := deviceclient.Connect()
	if err != nil {
		return
	}

	defer func() {
		conn.Close()
	}()
	for {
		mt, p, err := conn.ReadMessage()
		if err != nil {
			// 自动重连
			continue
		}
		switch mt {
		case websocket.TextMessage:
			PrintAction(p)
		case websocket.BinaryMessage:
			Motion(p)
		default:
			continue
		}
	}
	return nil
}
