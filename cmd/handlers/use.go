package handlers

import (
	"fmt"
	"gomv/config"
	"gomv/utils"
	"os"
	"path/filepath"
)

func UseVersion(version string) {
	srcDir := filepath.Join(config.VersionsDir, version, "go", "bin")

	// Verificamos si el directorio de la versión existe
	if _, err := os.Stat(srcDir); os.IsNotExist(err) {
		fmt.Printf("Version %s is not installed.\n", version)
		return
	}

	// Necesitamos permisos de administrador para copiar los archivos
	fmt.Println("Copying Go binaries to /usr/local/go/bin...")
	err := utils.CopyDirectory(srcDir, config.GoBinDir)
	if err != nil {
		fmt.Println("Error copying binaries:", err)
		return
	}

	fmt.Printf("Go version %s is now active.\n", version)
}
