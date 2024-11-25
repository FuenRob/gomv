package handlers

import (
	"fmt"
	"gomv/colors"
	"gomv/config"
	"gomv/utils"
	"os"
	"path/filepath"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func UseVersion(cmd *cobra.Command, args []string) {
	version := args[0]
	srcDir := filepath.Join(config.VersionsDir, version, "go", "bin")

	// Verificamos si el directorio de la versión existe
	if _, err := os.Stat(srcDir); os.IsNotExist(err) {
		colors.SetColor(color.FgHiRed, fmt.Sprintf("Version %s is not installed.\n", version))
		return
	}

	// Necesitamos permisos de administrador para copiar los archivos
	colors.SetColor(color.FgHiBlue, "Copying Go binaries to /usr/local/go/bin...\n")
	err := utils.CopyDirectory(srcDir, config.GoBinDir)
	if err != nil {
		colors.SetColor(color.FgHiRed, fmt.Sprintf("Error copying binaries: %v\n", err))
		return
	}

	colors.SetColor(color.FgHiGreen, fmt.Sprintf("Go version %s is now active.\n", version))
}
