package cmd

import (
	"fmt"
	"os"
	"time"

	r "github.com/Luiggy102/ligapro-cli/internal/requests"
	"github.com/Luiggy102/ligapro-cli/internal/utils"
	lg "github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
)

func PrintTable() {

	// lógica
	tableData := r.GetTable()
	headers := []string{"POS", "EQUIPO", "PJ", "G", "E", "D", "GF", "GE", "GD", "PTS"}
	data := [][]string{}
	for _, v := range tableData {
		data = append(data, []string{
			v.Position, v.Name, v.Stats.GamesPlayed, v.Stats.Wins, v.Stats.Draws, v.Stats.Loses,
			v.Stats.GoalsFor, v.Stats.GoalsAgainst, v.Stats.GoalDifference, v.Stats.Points,
		})
	}
	t := table.New().
		Headers(headers...).
		Rows(data...)

	// estílos
	hStyle := lg.NewStyle().
		Foreground(utils.Yellow)

	evenRowStyle := lg.NewStyle().
		Foreground(utils.DimWhite)

	oddRowStyle := lg.NewStyle()

	re := lg.NewRenderer(os.Stdout)

	titleStyle := lg.NewStyle().
		Width(80).
		Bold(true).
		Padding(1).
		Align(lg.Center).
		Foreground(utils.Yellow)

	t.
		Width(80).
		Border(lg.RoundedBorder()).
		StyleFunc(func(row, _ int) lg.Style {
			switch {
			case row == 0:
				return hStyle
			case row%2 == 0:
				return evenRowStyle
			default:
				return oddRowStyle
			}
		}).
		BorderStyle(re.NewStyle().Foreground(utils.Skyblue))

	// imprimir
	switch time.Now().Month() {
	case 3, 4, 5, 6, 7:
		fmt.Println(titleStyle.Render("Tabla De Posiciones LigaPro 1ra Etapa"))
	case 8, 9, 10, 11, 12:
		fmt.Println(titleStyle.Render("Tabla De Posiciones LigaPro 2da Etapa"))
	default:
		fmt.Println(titleStyle.Render("Tabla De Posiciones LigaPro"))
	}
	fmt.Println(t)
	fmt.Printf("\nActualizado: %v", time.Now().Format("02/01/2006"))
	fmt.Println("\nDatos Obtenidos de: espn.com")
}
