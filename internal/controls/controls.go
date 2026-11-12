package controls

import (
	"github.com/marc-antoinegelinas/feishin-controls/internal/websocket"
)


type SimpleEvent string

const (
	Next     SimpleEvent = "next"
	Pause    SimpleEvent = "pause"
	Play     SimpleEvent = "play"
	Previous SimpleEvent = "previous"
	Proxy    SimpleEvent = "proxy"
	Repeat   SimpleEvent = "repeat"
	Shuffle  SimpleEvent = "shuffle"
)

func ClientSimpleEvent(event SimpleEvent) {
	request := map[string]any{
		"event": string(event),
	}
	websocket.SendRequest(request)
}

func Position(position int) {
	request := map[string]any{
		"event":    "position",
		"position": position,
	}
	websocket.SendRequest(request)
}

func Favorite(favorite bool, songId string) {
	request := map[string]any{
		"event":    "favorite",
		"favorite": favorite,
		"id":       songId,
	}
	websocket.SendRequest(request)
}

func Rate(songId string, rating int) {
	request := map[string]any{
		"event":  "rating",
		"id":     songId,
		"rating": rating,
	}
	websocket.SendRequest(request)
}

func Volume(volume int) {
	request := map[string]any{
		"event":  "volume",
		"volume": volume,
	}
	websocket.SendRequest(request)
}
