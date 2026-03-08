package cmd

import (
	"log"

	"magshare/internal/handlers"

	"github.com/spf13/cobra"
)

var (
	sendSecure bool
)

var sendCmd = &cobra.Command{
	Use:   "send [file or directory]",
	Short: "Send a file or directory over the local network",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		path := args[0]

		opts := handlers.SendOptions{
			Port:   portFlag,
			Secure: sendSecure,
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

		if err := handlers.StartSendServer(path, opts); err != nil {
			log.Fatalf("\n[Error] %v\n", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(sendCmd)
	sendCmd.Flags().BoolVarP(&sendSecure, "secure", "s", false, "Require a PIN to download the file")
}
