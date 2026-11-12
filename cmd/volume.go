package cmd

import (
	"log"
	"strconv"

	"github.com/marc-antoinegelinas/feishin-controls/internal/controls"
	"github.com/marc-antoinegelinas/feishin-controls/internal/websocket"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(volumeCmd)

	volumeCmd.Flags().BoolVarP(&relativeVolume, "relative", "r", false, "Increase volume relative to current volume i.e +10")
	volumeCmd.Flags().BoolVarP(&backwardVolume, "backward", "b", false, "When relative, decrease volume by given amount")
}

var (
	relativeVolume bool
	backwardVolume bool
)

var volumeCmd = &cobra.Command{
	Use:   "volume [volume 1-100]",
	Short: "Set Feishin's volume",
	Long:  "Set the volume of Feishin from 0 to 100%",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		volume, err := strconv.Atoi(args[0])
		if err != nil {
			log.Fatalf("Failed to parse volume %d: %s", volume, err)
		}

		if volume < 0 || volume > 100 {
			log.Fatalf("Volume must be between 0 and 100. %d is outside this range", volume)
		}

		if relativeVolume {
			var response websocket.ServerResponse
			err = websocket.GetServerMessage(&response)
			if err != nil {
				log.Fatal("Error unmarshal:", err)
			}

			currVolume := response.Data.Volume

			if backwardVolume {
				volume = currVolume - volume
			} else {
				volume = currVolume + volume
			}
			if volume < 0 {
				volume = 0
			} else if volume > 100 {
				volume = 100
			}
		}

		controls.Volume(volume)
	},
}
