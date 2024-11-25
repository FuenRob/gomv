package handlers

import (
	"fmt"
	"gomv/config"
	"os"
)

func ListVersions() {
	versionDir := config.VersionsDir
	files, err := os.ReadDir(versionDir)
	if err != nil {
		fmt.Println("Error reading versions:", err)
		return
	}
	if len(files) == 0 {
		fmt.Printf("Empty folder %s\n", versionDir)
		return
	}
	fmt.Println("Listing installed Go versions...")
	for _, file := range files {
		fmt.Println(file.Name())
	}
}
