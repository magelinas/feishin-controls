package cmd

import (
	"encoding/json"
	"log"

	"github.com/marc-antoinegelinas/feishin-controls/internal/controls"
	"github.com/marc-antoinegelinas/feishin-controls/internal/websocket"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(playpauseCmd)
}

var playpauseCmd = &cobra.Command{
	Use:   "playpause",
	Short: "Play the current track, or pause if already playing",
	Long:  "Start playing the currently selected track if it is currently paused, or pause the currently selected track if it is currently playing",
	Run: func(cmd *cobra.Command, args []string) {
		var response controls.ServerResponse
		err := json.Unmarshal(websocket.GetServerMessage(), &response)
		if err != nil {
			log.Fatal("Error unmarshal:", err)
		}

		if response.Data.Status == "playing" {
			controls.ClientSimpleEvent(controls.Pause)
		} else {
			controls.ClientSimpleEvent(controls.Play)
		}
	},
}
