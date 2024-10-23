package socket

import (
	"fmt"
	"log"
	"net/http"
	redisclient "score/redisClient"

	"github.com/fasthttp/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func StartWebSocket() {

	// ws://localhost:8081/ws
	http.HandleFunc("/ws", handleWebsocket)

	fmt.Println("socket is initialized on port 8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func handleWebsocket(w http.ResponseWriter, r *http.Request) {

	conn, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Fatal(err.Error())
	}

	pubsub := redisclient.RedisClient().Subscribe(r.Context(), "match_updates")
	defer pubsub.Close()

	for {

		msg, err := pubsub.ReceiveMessage(r.Context())
		if err != nil {
			log.Fatal(err.Error())
		}

		if err := conn.WriteMessage(websocket.TextMessage, []byte(msg.Payload)); err != nil {
			log.Fatal(err.Error())
		}
	}

}
