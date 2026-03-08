package cmd

import (
	"log"

	"magshare/internal/handlers"

	"github.com/spf13/cobra"
)

var (
	receiveSecure bool
)

var receiveCmd = &cobra.Command{
	Use:   "receive",
	Short: "Start a dropzone server to receive files",
	Run: func(cmd *cobra.Command, args []string) {
		opts := handlers.ReceiveOptions{
			Port:   portFlag,
			Secure: receiveSecure,
			PIN:    pinFlag,
			Demo:   demoMode,
		}

		// Precedence Logic: 
		// 1. Explicit Flag
		// 2. Config File
		// 3. Application Default (false)
		if !cmd.Flags().Changed("secure") && appConfig.SecureMode {
			opts.Secure = true
		}

		if pinFlag != "" {
			opts.Secure = true
		}

		path := appConfig.DownloadDir

		if err := handlers.StartReceiveServer(path, opts); err != nil {
			log.Fatalf("\n[Error] %v\n", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(receiveCmd)
	receiveCmd.Flags().BoolVarP(&receiveSecure, "secure", "s", false, "Require a PIN to upload files")
}
