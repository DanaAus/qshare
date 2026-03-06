package cmd

import (
	"log"

	"magshare/internal/handlers"

	"github.com/spf13/cobra"
)

var receiveSecure bool

var receiveCmd = &cobra.Command{
	Use:   "receive",
	Short: "Start a dropzone server to receive files",
	Run: func(cmd *cobra.Command, args []string) {
		opts := handlers.ReceiveOptions{
			Secure: receiveSecure,
			Demo:   demoMode,
		}
		if err := handlers.StartReceiveServer("", opts); err != nil {
			log.Fatalf("\n[Error] %v\n", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(receiveCmd)
	receiveCmd.Flags().BoolVarP(&receiveSecure, "secure", "s", false, "Require a PIN to upload files")
}
