package cmd

import (
	"fmt"

	"github.com/Luiggy102/ligapro-cli/internal/requests"
	"github.com/Luiggy102/ligapro-cli/internal/utils"
)

func PrintStrikers() {
	goleadores := requests.GetStrikers()
	filas := [][]string{}
	for i := 0; i < len(goleadores.Goals)-1; i++ {
		filas = append(filas, []string{
			goleadores.Name[i], goleadores.Goals[i], goleadores.Team[i],
		})
	}
	// Imprimir
	const width = 65
	utils.PrintTitle(fmt.Sprintf("Tabla de goleadores"), width)
	utils.PrintTable(width, []string{"Jugador", "Goles", "Equipo"}, filas...)
	utils.LastUpdate()
}
