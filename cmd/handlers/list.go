package handlers

import (
	"fmt"
	"gomv/colors"
	"gomv/config"
	"os"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// ListVersions lists all installed Go versions in the configured versions directory.
// It reads the directory specified in the configuration and prints the names of the files found.
// If the directory is empty or an error occurs while reading, it prints appropriate messages.
//
// Parameters:
//   - cmd: The Cobra command that triggered this function.
//   - _: A slice of strings representing additional arguments (not used).
func ListVersions(cmd *cobra.Command, _ []string) {
	versionDir := config.VersionsDir
	files, err := os.ReadDir(versionDir)
	if err != nil {
		colors.SetColor(color.FgHiRed, fmt.Sprintf("Error listing folder %s: %s\n", versionDir, err))
		return
	}
	if len(files) == 0 {
		colors.SetColor(color.FgHiYellow, fmt.Sprintf("Empty folder %s\n", versionDir))
		return
	}
	colors.SetColor(color.FgBlue, "Listing installed Go versions...\n")
	for _, file := range files {
		colors.SetColor(color.FgBlue, file.Name()+"\n")
	}
}
