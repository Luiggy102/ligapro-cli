package cmd

import (
	"strconv"
	"time"

	"github.com/Luiggy102/ligapro-cli/internal/requests"
	"github.com/Luiggy102/ligapro-cli/internal/utils"
)

func PrintTable() {
	// lógica
	d := requests.GetPositionsTable()
	headers := []string{"POS", "EQUIPO", "PJ", "G", "E", "D", "GF", "GE", "GD", "PTS"}
	data := [][]string{}
	for i := 0; i < len(d.Names); i++ {
		data = append(data, []string{
			strconv.Itoa(i + 1), d.Names[i],
			d.GamesPlayed[i], d.Wins[i], d.Draws[i], d.Loses[i],
			d.GoalsFor[i], d.GoalsAgainst[i], d.GoalDifference[i],
			d.Points[i],
		})
	}
	// imprimir
	const width = 80
	switch time.Now().Month() {
	case 3, 4, 5, 6, 7:
		utils.PrintTitle("Tabla De Posiciones 1ra Etapa", width)
	case 8, 9, 10, 11, 12:
		utils.PrintTitle("Tabla De Posiciones 2da Etapa", width)
	default:
		utils.PrintTitle("Tabla De Posiciones", width)
	}
	utils.PrintTable(width, headers, data...)
	utils.LastUpdate()
}
