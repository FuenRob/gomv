package handlers

import (
	"fmt"
	"gomv/colors"
	"gomv/config"
	"os"
	"path/filepath"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// UninstallVersion uninstalls the specified Go version.
// It takes a cobra.Command and a slice of arguments as input.
// The first argument in the slice should be the version to uninstall.
//
// The function performs the following steps:
//  1. Constructs the directory path for the specified version.
//  2. Checks if the version exists in the directory.
//  3. If the version does not exist, it prints a warning message and returns.
//  4. If the version exists, it confirms the uninstallation and proceeds.
//  5. Attempts to remove the version directory.
//  6. If an error occurs during removal, it prints an error message and returns.
//  7. If successful, it prints a success message.
//
// Example usage:
//
//	UninstallVersion(cmd, []string{"1.16.3"})
//
// Note: This function assumes that config.VersionsDir is a valid directory path
// and that colors.SetColor is a valid function for setting terminal text colors.
func UninstallVersion(cmd *cobra.Command, args []string) {
	version := args[0]
	versionDir := filepath.Join(config.VersionsDir, version)

	// Verificar si la versión existe
	if _, err := os.Stat(versionDir); os.IsNotExist(err) {
		colors.SetColor(color.FgHiYellow, fmt.Sprintf("Version %s is not installed.\n", version))
		return
	}

	// Confirmar la eliminación y proceder
	fmt.Printf("Uninstalling Go version %s...\n", version)
	err := os.RemoveAll(versionDir)
	if err != nil {
		colors.SetColor(color.FgHiRed, fmt.Sprintf("Error uninstalling version: %s\n", err))
		return
	}

	colors.SetColor(color.FgHiGreen, fmt.Sprintf("Go version %s uninstalled successfully.\n", version))
}
