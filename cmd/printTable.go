package cmd

import (
	"fmt"
	"os"
	"time"

	r "github.com/Luiggy102/ligapro-cli/internal/requests"
	lg "github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
)

func PrintTable() {

	const (
		titleColor   = lg.Color("230")
		mainColor    = lg.Color("86")
		footerColor  = lg.Color("255")
		evenRowColor = lg.Color("248")
		oddRowColor  = lg.Color("255")
		errColor     = lg.Color("196")
	)

	// first check for an err
	tableData, err := r.GetTable()
	if err != nil {

		errStyle := lg.NewStyle().
			Width(80).
			Align(lg.Left).
			Foreground(errColor).
			Italic(true)

		panic(errStyle.Render("Error al obtener tabla, Intentalo de nuevo"))
	}
	// continue

	// title
	titleStyle := lg.NewStyle().
		Width(80).
		Bold(true).
		Padding(1).
		Underline(true).
		Align(lg.Center).
		Foreground(titleColor)

	switch time.Now().Month() {
	case time.March, time.April, time.May, time.June, time.July:
		fmt.Println(titleStyle.Render("Tabla De Posiciones LigaPro Ecuador 1ra Etapa 🇪🇨"))
	case time.August, time.September, time.October, time.November, time.December:
		fmt.Println(titleStyle.Render("Tabla De Posiciones LigaPro Ecuador 2da Etapa 🇪🇨"))
	default:
		fmt.Println(titleStyle.Render("Tabla De Posiciones LigaPro Ecuador 🇪🇨"))
	}

	// table
	headers := []string{"POS", "EQUIPO", "PJ", "G", "E", "D", "GF", "GE", "GD", "PTS"}
	data := [][]string{}
	for _, v := range tableData {
		data = append(data, []string{
			v.Position, v.Name, v.Stats.GamesPlayed, v.Stats.Wins, v.Stats.Draws, v.Stats.Loses,
			v.Stats.GoalsFor, v.Stats.GoalsAgainst, v.Stats.GoalDifference, v.Stats.Points,
		})
	}

	// table content
	t := table.New().
		Headers(headers...).
		Rows(data...)

	re := lg.NewRenderer(os.Stdout) // determine terminal colors

	// table style
	hStyle := lg.NewStyle().Foreground(lg.Color(mainColor))
	evenRowStyle := lg.NewStyle().Foreground(evenRowColor)
	oddRowStyle := lg.NewStyle()

	t.
		Width(80).
		Border(lg.RoundedBorder()).
		StyleFunc(func(row, col int) lg.Style {
			switch {
			case row == 0:
				return hStyle
			case row%2 == 0:
				return evenRowStyle
			default:
				return oddRowStyle
			}
		}).
		BorderStyle(re.NewStyle().Foreground(mainColor))

	fmt.Println(t)

	// footer
	footerStyle := lg.NewStyle().
		Align(lg.Left).
		Foreground(footerColor).
		Width(80)

	fmt.Println(footerStyle.Render("Datos Obtenidos de: espn.com"))
}
