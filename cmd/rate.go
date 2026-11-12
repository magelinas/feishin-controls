package cmd

import (
	"log"
	"strconv"

	"github.com/marc-antoinegelinas/feishin-controls/internal/controls"
	"github.com/marc-antoinegelinas/feishin-controls/internal/websocket"
	"github.com/spf13/cobra"
)

var (
	relativeRating bool
	backwardRating bool
)

func init() {
	rootCmd.AddCommand(rateCmd)

	rateCmd.Flags().BoolVarP(&relativeRating, "relative", "r", false, "Increase rating relative to current rating i.e. +1")
	rateCmd.Flags().BoolVarP(&backwardRating, "backward", "b", false, "When relative, decrease rating by given rating")
}

var rateCmd = &cobra.Command{
	Use:       "rate [1-5]",
	Short:     "Rate (star) the current track",
	Long:      "Set the rating for the currently playing song from 1 to 5 star",
	Args:      cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
	ValidArgs: []string{"1", "2", "3", "4", "5"},
	Run: func(cmd *cobra.Command, args []string) {
		rating, err := strconv.Atoi(args[0])
		if err != nil {
			log.Fatalf("Failed to parse rating %d: %s", rating, err)
		}
		var response websocket.ServerResponse
		err = websocket.GetServerMessage(&response)
		if err != nil {
			log.Fatal("Error unmarshal:", err)
		}

		songId := response.Data.Song.ID
		if relativeRating {
			currRating := response.Data.Song.UserRating
			if backwardRating {
				rating = currRating - rating
			} else {
				rating = currRating + rating
			}
		}

		if rating > 5 {
			rating = 5
		} else if rating < 0 {
			rating = 0
		}

		controls.Rate(songId, rating)
	},
}
