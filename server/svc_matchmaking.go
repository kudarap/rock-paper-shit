package server

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/websocket"
	"github.com/kudarap/rockpapershit"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
} // use default options

func FindMatchWs(s *rockpapershit.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		playerID := r.URL.Query().Get("player_id")
		if playerID == "" {
			encodeJSONError(w, errors.New("player_id is required"), http.StatusBadRequest)
			log.Println("err playerID")
			return
		}

		if err := s.QueuePlayer(r.Context(), playerID); err != nil {
			log.Println("err QueuePlayer", err)
			encodeJSONError(w, err, http.StatusInternalServerError)
			return
		}

		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Print("upgrade:", err)
			return
		}
		defer conn.Close()

		for {
			select {
			case token := <-s.NotifQ:
				fmt.Println("token", token)

				tk := strings.Split(token, ".")
				gameID, playerID := tk[0], tk[1]
				if token == "" && token != playerID {
					continue
				}

				err = conn.WriteMessage(websocket.TextMessage, []byte(gameID))
				if err != nil {
					log.Println("write:", err)
					return
				}

				fmt.Println("sent", token)
				return
			}
		}

		//for {
		//	time.Sleep(2 * time.Second)
		//	err = conn.WriteMessage(websocket.TextMessage, []byte("token"))
		//	if err != nil {
		//		log.Println("write:", err)
		//		break
		//	}
		//}
	}
}

func FindMatch(s service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		playerID := r.URL.Query().Get("player_id")
		if playerID == "" {
			encodeJSONError(w, errors.New("player_id is required"), http.StatusBadRequest)
			return
		}

		if err := s.QueuePlayer(r.Context(), playerID); err != nil {
			encodeJSONError(w, err, http.StatusInternalServerError)
			return
		}

		response := map[string]string{"status": "queued"}
		encodeJSONResp(w, response, http.StatusOK)
	}
}
