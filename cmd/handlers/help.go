package handlers

import (
	"gomv/colors"
	"os"

	"github.com/fatih/color"
)

func HelpUser() {
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
