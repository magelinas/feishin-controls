package websocket

import (
	"context"
	"time"
	"encoding/base64"
	"fmt"
	"log"
	"net/url"
	"encoding/json"
	"runtime"
	"sync"

	"github.com/coder/websocket"
	"github.com/coder/websocket/wsjson"
	"github.com/spf13/viper"
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

var (
	connection *websocket.Conn
	connOnce   sync.Once
)

func Authenticate() {
	connOnce.Do(func() {
		u := url.URL{Scheme: "ws", Host: viper.GetString("URL"), Path: "/"}

		creds64 := base64.StdEncoding.EncodeToString(fmt.Appendf(nil, "%s:%s", viper.GetString("Username"), viper.GetString("Password")))

		authHeader := map[string]string{
			"event":  "authenticate",
			"header": fmt.Sprintf("Basic %s", creds64),
		}

		var err error

		ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
		defer cancel()
		connection, _, err = websocket.Dial(ctx, u.String(), nil)
		if err != nil {
			log.Fatal("Dial error:", err)
		}

		err = wsjson.Write(ctx, connection, authHeader)
		if err != nil {
			log.Fatal("Request error:", err)
		}
	})
}

func GetServerMessage(response *ServerResponse) error {
	if connection == nil {
		_, file, no, _ := runtime.Caller(1)
		log.Fatalf("Connection error in %s#%d\n", file, no)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	_, bytes, err := connection.Read(ctx)
	if err != nil {
		log.Fatal("Read:", err)
	}
	return json.Unmarshal(bytes, response)
}

func SendRequest(request map[string]any) {
	if connection == nil {
		_, file, no, _ := runtime.Caller(1)
		log.Fatalf("Connection error in %s#%d\n", file, no)
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	err := wsjson.Write(ctx, connection, request)
	if err != nil {
		log.Fatal("Request:", err)
	}

	connection.Close(websocket.StatusNormalClosure, "")
}
