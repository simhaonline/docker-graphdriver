package cmd

import (
	"os"
	"os/signal"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/cvmfs/docker-graphdriver/daemon/lib"
)

func init() {
	loopCmd.Flags().BoolVarP(&overwriteLayer, "overwrite-layers", "f", false, "overwrite the layer if they are already inside the CVMFS repository")
	loopCmd.Flags().BoolVarP(&convertAgain, "convert-again", "g", false, "convert again images that are already successfull converted")
	rootCmd.AddCommand(loopCmd)
}

var loopCmd = &cobra.Command{
	Use:   "loop",
	Short: "An infinite loop that keep converting all the images",
	Run: func(cmd *cobra.Command, args []string) {
		AliveMessage()
		showWeReceivedSignal := make(chan os.Signal, 1)
		signal.Notify(showWeReceivedSignal, os.Interrupt)

		stopWishLoopSignal := make(chan os.Signal, 1)
		signal.Notify(stopWishLoopSignal, os.Interrupt)

		go func() {
			<-showWeReceivedSignal
			lib.Log().Info("Received SIGINT (Ctrl-C) waiting the last layer to upload then exiting.")
		}()

		for {
			wish, err := lib.GetAllWishes()
			if err != nil {
				lib.LogE(err).Error("Error in getting the desiderata")
			}
			for _, wish := range wish {

				select {
				case <-stopWishLoopSignal:
					lib.Log().Info("Exiting because of SIGINT")
					os.Exit(1)
				default:
					{
					}
				}

				fields := log.Fields{
					"input image":  wish.InputName,
					"CVMFS repo":   wish.CvmfsRepo,
					"output image": wish.OutputName,
				}
				lib.Log().WithFields(fields).Info("Working on desiderata")
				err = lib.ConvertWish(wish, convertAgain, overwriteLayer)
				if err != nil {
					lib.LogE(err).Error("Error in converting the desiderata")
				}
			}
		}
	},
}
