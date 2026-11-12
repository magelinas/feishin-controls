package cmd

import (
	"log"

	"github.com/marc-antoinegelinas/feishin-controls/internal/controls"
	"github.com/marc-antoinegelinas/feishin-controls/internal/websocket"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(favoriteCmd)
}

var favoriteCmd = &cobra.Command{
	Use:   "favorite",
	Short: "Toggle track favorite status",
	Long:  "Toggle track favorite status between favorite or not",
	Run: func(cmd *cobra.Command, args []string) {
		var response websocket.ServerResponse
		err := websocket.GetServerMessage(&response)
		if err != nil {
			log.Fatal("Error unmarshal:", err)
		}

		songId := response.Data.Song.ID
		favorite := !response.Data.Song.UserFavorite

		controls.Favorite(favorite, songId)
	},
}
