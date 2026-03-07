package cmd

import (
	"magshare/internal/handlers"
	"magshare/internal/ui"

	"github.com/spf13/cobra"
)

var (
	demoMode bool
	portFlag int
	pinFlag  string
)

var rootCmd = &cobra.Command{
	Use:   "MagShare",
	Short: "MagShare is an instant local network file sharing tool",
	Long:  `MagShare allows you to instantly share and receive files across your local network. It spawns ephemeral web servers and provides QR codes for easy access.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// If no subcommand is provided, run interactive mode
		cfg, err := ui.RunInteractivePrompts(demoMode)
		if err != nil {
			if err.Error() == "user cancelled" {
				return nil
			}
			return err
		}

		// Override interactive values if flags are provided
		if portFlag > 0 {
			cfg.Port = portFlag
		}
		if pinFlag != "" {
			cfg.PIN = pinFlag
			cfg.Secure = true
		}

		switch cfg.Action {
		case "send":
			opts := handlers.SendOptions{
				Port:   cfg.Port,
				Secure: cfg.Secure,
				PIN:    cfg.PIN,
				Demo:   demoMode,
			}
			return handlers.StartSendServer(cfg.Path, opts)
		case "receive":
			opts := handlers.ReceiveOptions{
				Port:   cfg.Port,
				Secure: cfg.Secure,
				PIN:    cfg.PIN,
				Demo:   demoMode,
			}
			return handlers.StartReceiveServer(cfg.Path, opts)
		default:
			return cmd.Help()
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	// Global flags can go here
	rootCmd.PersistentFlags().BoolVarP(&demoMode, "demo", "", false, "Enable demo mode with fake connection information")
	rootCmd.PersistentFlags().IntVarP(&portFlag, "port", "p", 0, "Custom port for the server")
	rootCmd.PersistentFlags().StringVarP(&pinFlag, "pin", "", "", "Custom 4-digit numeric PIN")
}
