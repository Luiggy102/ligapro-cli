package utils

import (
	"os"
	"os/exec"
)

func Check(err error) {
	if err != nil {
		panic(err)
	}
}

func ClearTerminal() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
