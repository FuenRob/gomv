package config

import (
	"os"
	"path/filepath"
)

var VersionsDir = filepath.Join(os.Getenv("HOME"), ".govm", "versions") // Cambia a tu ruta de versiones
const GoBinDir = "/usr/local/go/bin"                                    // Ruta donde deben copiarse los archivos binarios
const Version string = "0.0.2"
