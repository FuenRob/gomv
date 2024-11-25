package handlers

import (
	"fmt"
	"gomv/colors"
	"gomv/config"
	"gomv/utils"
	"io"
	"net/http"
	"os"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// InstallVersion downloads and installs the specified version of Go.
// It takes a Cobra command and a list of arguments, where the first argument is the Go version to install.
// The function downloads the Go tarball from the official Go website, saves it to a temporary location,
// extracts it to the configured versions directory, and sets up the Go environment.
//
// Parameters:
//   - cmd: The Cobra command that triggered this function.
//   - args: A list of arguments where the first argument is the Go version to install.
//
// Example usage:
//
//	InstallVersion(cmd, []string{"1.16.3"})
func InstallVersion(cmd *cobra.Command, args []string) {
	version := args[0]
	url := fmt.Sprintf("https://golang.org/dl/go%s.linux-amd64.tar.gz", version)
	tarPath := fmt.Sprintf("/tmp/go%s.tar.gz", version)

	colors.SetColor(color.FgHiBlue, fmt.Sprintf("Downloading Go %s...\n", version))
	resp, err := http.Get(url)
	if err != nil {
		colors.SetColor(color.FgHiRed, fmt.Sprintf("Error downloading file: %v\n", err))
		return
	}
	defer resp.Body.Close()

	out, err := os.Create(tarPath)
	if err != nil {
		colors.SetColor(color.FgHiRed, fmt.Sprintf("Error creating file: %v\n", err))
		return
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		colors.SetColor(color.FgHiRed, fmt.Sprintf("Error downloading file: %v\n", err))
		return
	}

	colors.SetColor(color.FgHiBlue, fmt.Sprintf("Downloaded Go %s\n", version))
	err = utils.ExtractTarGz(tarPath, config.VersionsDir, version)
	if err != nil {
		colors.SetColor(color.FgHiRed, fmt.Sprintf("Error extracting file: %v\n", err))
		return
	}

	colors.SetColor(color.FgHiGreen, fmt.Sprintf("Setting up Go %s...\n", version))
}
