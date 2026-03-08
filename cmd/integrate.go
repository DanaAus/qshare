package cmd

import (
	"fmt"
	"magshare/internal/workspace"
	"runtime"

	"github.com/spf13/cobra"
)

var (
	installFlag   bool
	uninstallFlag bool
	statusFlag    bool
)

var integrateCmd = &cobra.Command{
	Use:   "integrate",
	Short: "Manage Windows Explorer context menu integration",
	Long:  `Manage the 'Share via Magshare' option in the Windows right-click context menu.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if runtime.GOOS != "windows" {
			return fmt.Errorf("explorer integration is only supported on Windows")
		}

		if installFlag {
			fmt.Println("Installing Windows context menu integration...")
			if err := workspace.RegisterContextMenu(); err != nil {
				return fmt.Errorf("failed to install integration: %w", err)
			}
			fmt.Println("Successfully installed 'Share via Magshare' to context menu.")
			return nil
		}

		if uninstallFlag {
			fmt.Println("Removing Windows context menu integration...")
			if err := workspace.UnregisterContextMenu(); err != nil {
				return fmt.Errorf("failed to remove integration: %w", err)
			}
			fmt.Println("Successfully removed 'Share via Magshare' from context menu.")
			return nil
		}

		if statusFlag {
			fmt.Println("Checking integration status...")
			isRegistered := workspace.IsContextMenuItemRegistered()
			if isRegistered {
				fmt.Println("Status: ENABLED (Magshare is in the context menu)")
			} else {
				fmt.Println("Status: DISABLED (Magshare is NOT in the context menu)")
			}
			fmt.Printf("Configured preference: %v\n", appConfig.ShellIntegration)
			return nil
		}

		return cmd.Help()
	},
}

func init() {
	rootCmd.AddCommand(integrateCmd)

	integrateCmd.Flags().BoolVar(&installFlag, "install", false, "Add 'Share via Magshare' to Windows right-click menu")
	integrateCmd.Flags().BoolVar(&uninstallFlag, "uninstall", false, "Remove 'Share via Magshare' from Windows right-click menu")
	integrateCmd.Flags().BoolVar(&statusFlag, "status", false, "Check if Windows context menu is integrated")
}
