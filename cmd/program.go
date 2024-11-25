package cmd

import (
	"fmt"
	"gomv/cmd/handlers"
	"gomv/colors"
	"gomv/config"
	"os"

	"github.com/fatih/color"
)

func Init() {
	if len(os.Args) < 2 {
		colors.SetColor(color.FgGreen, "Usage: govm <command> [version] or <help> to more info\n")
		return
	}

	command := os.Args[1]
	switch command {
	case "help":
		handlers.HelpUser()
	case "version":
		fmt.Printf("govm %s\n", config.Version)
		os.Exit(0)
	case "config":
		handlers.SetupConfig()
	case "list":
		handlers.ListVersions()
	case "use":
		if len(os.Args) < 3 {
			fmt.Println("Please specify a version")
			return
		}
		handlers.UseVersion(os.Args[2])
	case "install":
		if len(os.Args) < 3 {
			fmt.Println("Please specify a version")
			return
		}
		handlers.InstallVersion(os.Args[2])
		handlers.UseVersion(os.Args[2])
	case "uninstall":
		if len(os.Args) < 3 {
			fmt.Println("Please specify a version")
			return
		}
		handlers.UninstallVersion(os.Args[2])
	default:
		colors.SetColor(color.FgRed, "Unknown command: %s\n", command)
		os.Exit(1)
	}
}
