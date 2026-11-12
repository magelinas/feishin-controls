package cmd

import (
	"log"
	"strconv"

	"github.com/marc-antoinegelinas/feishin-controls/internal/controls"
	"github.com/marc-antoinegelinas/feishin-controls/internal/websocket"
	"github.com/spf13/cobra"
)

var (
	relativePos bool
	backwardPos bool
)

func init() {
	rootCmd.AddCommand(positionCmd)

	positionCmd.Flags().BoolVarP(&relativePos, "relative", "r", false, "New position is relative to track current position.")
	positionCmd.Flags().BoolVarP(&backwardPos, "backward", "b", false, "Rewinds current position with new position.")
}

var positionCmd = &cobra.Command{
	Use:   "position [time]",
	Short: "Set track position",
	Long:  "Set the track current playing position",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		pos, err := strconv.Atoi(args[0])
		if err != nil {
			log.Fatalf("Failed to parse argument to int", pos, err.Error())
		}

		if relativePos {
			var response websocket.ServerResponse
			err := websocket.GetServerMessage(&response)
			if err != nil {
				log.Fatal("Error unmarshal:", err)
			}

			currPos := int(response.Data.Position)

			if backwardPos {
				pos = currPos - pos
			} else {
				pos = currPos + pos
			}
		}

		controls.Position(pos)
	},
}
