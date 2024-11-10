// Copyright © 2024 GOMV
// Licensed under the MIT License. See LICENSE file for details.

package main

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"gomv/colors"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/fatih/color"
)

var versionsDir = filepath.Join(os.Getenv("HOME"), ".govm", "versions") // Cambia a tu ruta de versiones
const goBinDir = "/usr/local/go/bin"                                    // Ruta donde deben copiarse los archivos binarios
const Version string = "0.0.2"

// Función para asegurar que un directorio existe, creándolo si es necesario
func ensureDirExists(dir string) error {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err := os.MkdirAll(dir, 0755)
		if err != nil {
			return fmt.Errorf("failed to create directory %s: %w", dir, err)
		}
	}
	return nil
}

// Función para listar versiones instaladas
func listVersions() {
	files, err := os.ReadDir(versionsDir)
	if err != nil {
		fmt.Println("Error reading versions:", err)
		return
	}
	if len(files) == 0 {
		fmt.Printf("Empty folder %s\n", versionsDir)
		return
	}
	fmt.Println("Listing installed Go versions...")
	for _, file := range files {
		fmt.Println(file.Name())
	}
}

// Función para seleccionar una versión específica
// Función para copiar archivos binarios a /usr/local/go/bin
func useVersion(version string) {
	srcDir := filepath.Join(versionsDir, version, "go", "bin")
	destDir := goBinDir

	// Verificamos si el directorio de la versión existe
	if _, err := os.Stat(srcDir); os.IsNotExist(err) {
		fmt.Printf("Version %s is not installed.\n", version)
		return
	}

	// Necesitamos permisos de administrador para copiar los archivos
	fmt.Println("Copying Go binaries to /usr/local/go/bin...")
	err := copyDirectory(srcDir, destDir)
	if err != nil {
		fmt.Println("Error copying binaries:", err)
		return
	}

	fmt.Printf("Go version %s is now active.\n", version)
}

// Función para copiar el contenido de un directorio a otro
func copyDirectory(srcDir, destDir string) error {
	// Asegurarse de que el directorio destino existe
	err := ensureDirExists(destDir)
	if err != nil {
		return err
	}

	entries, err := os.ReadDir(srcDir)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		srcPath := filepath.Join(srcDir, entry.Name())
		destPath := filepath.Join(destDir, entry.Name())

		// Copiamos archivos regulares
		if entry.Type().IsRegular() {
			srcFile, err := os.Open(srcPath)
			if err != nil {
				return err
			}
			defer srcFile.Close()

			destFile, err := os.Create(destPath)
			if err != nil {
				return err
			}
			defer destFile.Close()

			_, err = io.Copy(destFile, srcFile)
			if err != nil {
				return err
			}
			// Asegúrate de dar permisos ejecutables a los binarios
			err = os.Chmod(destPath, 0755)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// Función para descargar e instalar una versión de Go
func installVersion(version string) {
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
	err = extractTarGz(tarPath, versionsDir, version)
	if err != nil {
		fmt.Println("Error extracting file:", err)
		return
	}

	fmt.Printf("Go %s installed successfully.\n", version)
}

// Función para extraer el archivo tar.gz en el directorio de versiones
func extractTarGz(tarGzPath, targetDir, version string) error {
	file, err := os.Open(tarGzPath)
	if err != nil {
		return err
	}
	defer file.Close()

	gzipReader, err := gzip.NewReader(file)
	if err != nil {
		return err
	}
	defer gzipReader.Close()

	tarReader := tar.NewReader(gzipReader)
	versionDir := filepath.Join(targetDir, version)

	os.MkdirAll(versionDir, 0755)
	for {
		header, err := tarReader.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("could not read tar entry: %w", err)
		}

		path := filepath.Join(versionDir, header.Name)

		switch header.Typeflag {
		case tar.TypeDir:
			if err := ensureDirExists(path); err != nil {
				return err
			}
		case tar.TypeReg:
			outFile, err := os.Create(path)
			if err != nil {
				return err
			}

			if _, err := io.Copy(outFile, tarReader); err != nil {
				outFile.Close()
				return err
			}

			outFile.Close()
		}
	}
	return nil
}

// Función para desinstalar una versión de Go
func uninstallVersion(version string) {
	versionDir := filepath.Join(versionsDir, version)

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

func helpUser() {
	colors.SetColor(color.FgGreen, "Use: govm <command> [version]\n")
	colors.SetColor(color.FgGreen, "help \n\t Get help ready and end with a successful exit\n")
	colors.SetColor(color.FgGreen, "config \n\t Create the folder in the home path\n")
	colors.SetColor(color.FgGreen, "list \n\t List installed versions of go\n")
	colors.SetColor(color.FgGreen, "use \n\t use the specific version of go, if installed\n")
	colors.SetColor(color.FgGreen, "install \n\t Install the specified version of go\n")
	colors.SetColor(color.FgGreen, "uninstall \n\t uninstall the specific version of go\n")
	colors.SetColor(color.FgGreen, "version \n\t Displays the version and exits successfully\n")
	os.Exit(0)
}

func main() {
	if len(os.Args) < 2 {
		colors.SetColor(color.FgGreen, "Usage: govm <command> [version] or <help> to more info\n")
		return
	}

	command := os.Args[1]
	switch command {
	case "help":
		helpUser()
	case "version":
		fmt.Printf("govm %s\n", Version)
		os.Exit(0)
	case "config":
		if err := ensureDirExists(versionsDir); err != nil {
			fmt.Printf("Error creating config directory: %v\n", err)
		}
		ensureDirExists(versionsDir)
	case "list":
		listVersions()
	case "use":
		if len(os.Args) < 3 {
			fmt.Println("Please specify a version")
			return
		}
		useVersion(os.Args[2])
	case "install":
		if len(os.Args) < 3 {
			fmt.Println("Please specify a version")
			return
		}
		installVersion(os.Args[2])
	case "uninstall":
		if len(os.Args) < 3 {
			fmt.Println("Please specify a version")
			return
		}
		uninstallVersion(os.Args[2])
	default:
		colors.SetColor(color.FgRed, "Unknown command: %s\n", command)
		os.Exit(1)
	}
}
