package handlers

import (
	"fmt"
	"gomv/colors"
	"gomv/config"
	"gomv/utils"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// SetupConfig initializes the configuration for the given Cobra command.
// It ensures that the versions directory exists, creating it if necessary.
//
// Parameters:
//   - cmd: The Cobra command to set up the configuration for.
//   - _ : A slice of strings (unused).
//
// The function will log an error message if it fails to create the versions directory.
func SetupConfig(cmd *cobra.Command, _ []string) {
	versionsDir := config.VersionsDir
	if err := utils.EnsureDirExists(versionsDir); err != nil {
		colors.SetColor(color.FgHiRed, fmt.Sprintf("Error creating directory %s: %s", versionsDir, err))
	}
	utils.EnsureDirExists(versionsDir)
}
