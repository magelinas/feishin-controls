package controls

import (
	"time"

	"github.com/marc-antoinegelinas/feishin-controls/internal/websocket"
)

type ServerResponse struct {
	Data struct {
		Status  string `json:"status"`
		Repeat  string `json:"repeat"`
		Shuffle bool   `json:"shuffle"`
		Volume  int    `json:"volume"`
		Song    struct {
			Album        string `json:"album"`
			AlbumID      string `json:"albumId"`
			AlbumArtists []struct {
				ID       string `json:"id"`
				ImageURL any    `json:"imageUrl"`
				Name     string `json:"name"`
			} `json:"albumArtists"`
			Artists []struct {
				ID       string `json:"id"`
				ImageURL any    `json:"imageUrl"`
				Name     string `json:"name"`
			} `json:"artists"`
			Participants struct{} `json:"participants"`
			ArtistName   string   `json:"artistName"`
			BitRate      int      `json:"bitRate"`
			Bpm          any      `json:"bpm"`
			Channels     int      `json:"channels"`
			Comment      any      `json:"comment"`
			Compilation  bool     `json:"compilation"`
			Container    string   `json:"container"`
			CreatedAt    string   `json:"createdAt"`
			DiscNumber   int      `json:"discNumber"`
			DiscSubtitle any      `json:"discSubtitle"`
			Duration     int      `json:"duration"`
			Gain         struct {
				Album float64 `json:"album"`
				Track float64 `json:"track"`
			} `json:"gain"`
			Genres []struct {
				ID       string `json:"id"`
				ImageURL any    `json:"imageUrl"`
				ItemType string `json:"itemType"`
				Name     string `json:"name"`
			} `json:"genres"`
			ID                  string `json:"id"`
			ImagePlaceholderURL any    `json:"imagePlaceholderUrl"`
			ImageURL            string `json:"imageUrl"`
			ItemType            string `json:"itemType"`
			LastPlayedAt        any    `json:"lastPlayedAt"`
			Lyrics              string `json:"lyrics"`
			Name                string `json:"name"`
			Path                string `json:"path"`
			Peak                struct {
				Album float64 `json:"album"`
				Track float64 `json:"track"`
			} `json:"peak"`
			PlayCount      int       `json:"playCount"`
			PlaylistItemID string    `json:"playlistItemId"`
			ReleaseDate    string    `json:"releaseDate"`
			ReleaseYear    int       `json:"releaseYear"`
			ServerID       string    `json:"serverId"`
			ServerType     string    `json:"serverType"`
			Size           int       `json:"size"`
			StreamURL      string    `json:"streamUrl"`
			Tags           struct {
				Disctotal      []string `json:"disctotal"`
				Genre          []string `json:"genre"`
				Media          []string `json:"media"`
				Recordlabel    []string `json:"recordlabel"`
				Releasecountry []string `json:"releasecountry"`
				Releasestatus  []string `json:"releasestatus"`
				Releasetype    []string `json:"releasetype"`
				Tracktotal     []string `json:"tracktotal"`
			} `json:"tags"`
			TrackNumber  int       `json:"trackNumber"`
			UniqueID     string    `json:"uniqueId"`
			UpdatedAt    time.Time `json:"updatedAt"`
			UserFavorite bool      `json:"userFavorite"`
			UserRating   int       `json:"userRating"`
		} `json:"song"`
		Position float64 `json:"position"`
	} `json:"data"`
	Event string `json:"event"`
}

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
