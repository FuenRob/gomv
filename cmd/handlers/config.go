package handlers

import (
	"fmt"
	"gomv/config"
	"gomv/utils"
)

func SetupConfig() {
	versionsDir := config.VersionsDir
	if err := utils.EnsureDirExists(versionsDir); err != nil {
		fmt.Printf("Error creating config directory: %v\n", err)
	}
	utils.EnsureDirExists(versionsDir)
}
