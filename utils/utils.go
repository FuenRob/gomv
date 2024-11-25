package utils

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// EnsureDirExists checks if a directory exists, and if not, creates it with the specified permissions.
// It returns an error if the directory cannot be created.
//
// Parameters:
//   - dir: The path of the directory to check or create.
//
// Returns:
//   - error: An error if the directory cannot be created, or nil if the directory exists or is successfully created.
//
// Example:
//
//	err := EnsureDirExists("/path/to/dir")
//	if err != nil {
//	    log.Fatalf("Error ensuring directory exists: %v", err)
//	}
func EnsureDirExists(dir string) error {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err := os.MkdirAll(dir, 0755)
		if err != nil {
			return fmt.Errorf("failed to create directory %s: %w", dir, err)
		}
	}
	return nil
}

// CopyDirectory copies all regular files from the source directory (srcDir) to the destination directory (destDir).
// It ensures that the destination directory exists before copying files.
// If a file is copied successfully, it sets the file permissions to 0755 (executable).
//
// Parameters:
//   - srcDir: The path to the source directory.
//   - destDir: The path to the destination directory.
//
// Returns:
//   - error: An error if any occurs during the process, otherwise nil.
func CopyDirectory(srcDir, destDir string) error {
	// Asegurarse de que el directorio destino existe
	err := EnsureDirExists(destDir)
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

// ExtractTarGz extracts a .tar.gz archive to a specified directory with a version subdirectory.
// It takes the path to the .tar.gz file, the target directory, and the version string as arguments.
// The function creates the necessary directories and files as specified in the archive.
//
// Parameters:
//   - tarGzPath: The file path to the .tar.gz archive.
//   - targetDir: The directory where the contents of the archive will be extracted.
//   - version: The version subdirectory where the contents will be placed.
//
// Returns:
//   - error: An error if any occurs during the extraction process, otherwise nil.
//
// Example:
//
//	err := ExtractTarGz("/path/to/archive.tar.gz", "/target/directory", "v1.0.0")
//	if err != nil {
//	    log.Fatal(err)
//	}
func ExtractTarGz(tarGzPath, targetDir, version string) error {
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
			if err := EnsureDirExists(path); err != nil {
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
