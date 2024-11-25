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

// UseVersion activates a specified Go version by copying its binaries to the Go binary directory.
// It requires administrator permissions to perform the copy operation.
//
// Parameters:
//   - cmd: The cobra command that triggered this function.
//   - args: A slice of arguments where the first element is the version to be activated.
//
// The function performs the following steps:
//  1. Constructs the source directory path for the specified version's binaries.
//  2. Checks if the source directory exists. If not, it prints an error message and returns.
//  3. Copies the binaries from the source directory to the Go binary directory.
//  4. Prints success or error messages based on the outcome of the copy operation.
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
