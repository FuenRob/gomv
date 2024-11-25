package handlers

import (
	"fmt"
	"gomv/colors"
	"gomv/config"
	"os"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

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
