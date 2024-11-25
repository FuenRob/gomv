package handlers

import (
	"fmt"
	"gomv/config"
	"os"
	"path/filepath"
)

func UninstallVersion(version string) {
	versionDir := filepath.Join(config.VersionsDir, version)

	// Verificar si la versión existe
	if _, err := os.Stat(versionDir); os.IsNotExist(err) {
		fmt.Printf("Version %s is not installed.\n", version)
		return
	}

	// Confirmar la eliminación y proceder
	fmt.Printf("Uninstalling Go version %s...\n", version)
	err := os.RemoveAll(versionDir)
	if err != nil {
		fmt.Println("Error uninstalling version:", err)
		return
	}

	fmt.Printf("Go version %s uninstalled successfully.\n", version)
}
