package main

import (
	"os"

	"github.com/xuender/fairy/cmd"
)

func main() {
	cmd.Execute()
	os.Unsetenv("FYNE_FONT")
}
