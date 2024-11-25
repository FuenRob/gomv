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
