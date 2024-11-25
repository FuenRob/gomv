package handlers

import (
	"fmt"
	"gomv/config"
	"gomv/utils"
	"io"
	"net/http"
	"os"
)

func InstallVersion(version string) {
	url := fmt.Sprintf("https://golang.org/dl/go%s.linux-amd64.tar.gz", version)
	tarPath := fmt.Sprintf("/tmp/go%s.tar.gz", version)

	fmt.Printf("Downloading Go %s...\n", version)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error downloading file:", err)
		return
	}
	defer resp.Body.Close()

	out, err := os.Create(tarPath)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		fmt.Println("Error saving file:", err)
		return
	}

	fmt.Printf("Extracting Go %s...\n", version)
	err = utils.ExtractTarGz(tarPath, config.VersionsDir, version)
	if err != nil {
		fmt.Println("Error extracting file:", err)
		return
	}

	fmt.Printf("Go %s installed successfully.\n", version)
}
