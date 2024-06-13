package cmd

import (
	"strconv"

	"github.com/Luiggy102/ligapro-cli/internal/requests"
	"github.com/Luiggy102/ligapro-cli/internal/utils"
)

func PrintTable(year int, season int) {
	// lógica
	d := requests.GetPositionsTable(year, season)
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
	utils.PrintTitle(d.TableTitle, width)
	utils.PrintTable(width, headers, data...)
	utils.LastUpdate()
}
