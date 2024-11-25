package cmd

import (
	"fmt"
	"gomv/cmd/handlers"

	"github.com/spf13/cobra"
)

func Init() {
	var rootCmd = &cobra.Command{
		Use:   "govm",
		Short: "Descripción corta principal",
		Long:  `Descripción corta larga principal`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Hello, world!")
		},
	}

	var installCmd = &cobra.Command{
		Use:   "install",
		Short: "Install the specified version of go",
		Long:  `Use this command to install a specific version of Go`,
		Args:  cobra.ExactArgs(1),
		Run:   handlers.InstallVersion,
	}

	var listCmd = &cobra.Command{
		Use:   "list",
		Short: "List all installed versions of go",
		Long:  `Use this command to list all installed versions of Go`,
		Run:   handlers.ListVersions,
	}

	var setupConfigCmd = &cobra.Command{
		Use:   "setup",
		Short: "Setup the configuration for govm",
		Long:  `Use this command to setup the configuration for govm`,
		Run:   handlers.SetupConfig,
	}

	var uninstallCmd = &cobra.Command{
		Use:   "uninstall",
		Short: "Uninstall the specified version of go",
		Long:  `Use this command to uninstall a specific version of Go`,
		Args:  cobra.ExactArgs(1),
		Run:   handlers.UninstallVersion,
	}

	var useCmd = &cobra.Command{
		Use:   "use",
		Short: "Use the specific version of go, if installed",
		Long:  `Use this command to switch to a specific version of Go`,
		Args:  cobra.ExactArgs(1),
		Run:   handlers.UseVersion,
	}

	rootCmd.AddCommand(installCmd, listCmd, setupConfigCmd, uninstallCmd, useCmd)

	rootCmd.Execute()
}
