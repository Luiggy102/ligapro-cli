package main

import (
	"flag"

	"github.com/Luiggy102/ligapro-cli/cmd"
	"github.com/Luiggy102/ligapro-cli/internal/utils"
)

func main() {
	printTable := flag.Bool("tabla", false, "Muesta la tabla de posiciones actual")
	printResults := flag.Bool("resultados", false, "Mostrar los últimos resultados")
	printNext := flag.Bool("siguientes", false, "mostrar los próximos partidos")
	printStrikers := flag.Bool("goleadores", false, "Mostrar tabla de goleadores")
	printMatchday := flag.Int("fecha", 0, "Mostrar los partidos de dicha fecha")
	flag.Usage = utils.Help
	flag.Parse()

	switch {
	case *printTable:
		cmd.PrintTable()
		return
	case *printResults:
		cmd.PrintResults()
		return
	case *printNext:
		cmd.PrintNext()
		return
	case *printStrikers:
		cmd.PrintStrikers()
		return
	case *printMatchday != 0:
		cmd.PrintMatchday(*printMatchday)
		return
	default:
		utils.Welcome()
	}
}
