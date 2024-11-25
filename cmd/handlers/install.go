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
