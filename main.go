package main

import (
	"flag"

	"github.com/Luiggy102/ligapro-cli/cmd"
)

func main() {

	printTable := flag.Bool("tabla", false, "Muesta la tabla de posiciones actual")
	flag.Parse()

	if *printTable {
		cmd.PrintTable()
		return
	}

	cmd.PrintTable()
}
