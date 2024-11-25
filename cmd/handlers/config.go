package handlers

import (
	"fmt"
	"gomv/colors"
	"gomv/config"
	"gomv/utils"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func SetupConfig(cmd *cobra.Command, _ []string) {
	versionsDir := config.VersionsDir
	if err := utils.EnsureDirExists(versionsDir); err != nil {
		colors.SetColor(color.FgHiRed, fmt.Sprintf("Error creating directory %s: %s", versionsDir, err))
	}
	utils.EnsureDirExists(versionsDir)
}
