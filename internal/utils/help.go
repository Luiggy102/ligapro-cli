package utils

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
)

func Help() {

	t := table.New().Border(lipgloss.HiddenBorder()).Width(60)
	ecu := lipgloss.NewStyle().SetString("Ecu").Background(DimYellow).Foreground(Black)
	ad := lipgloss.NewStyle().SetString("ad").Background(Skyblue).Foreground(Black)
	or := lipgloss.NewStyle().SetString("or").Background(Red)
	title := lipgloss.NewStyle().Bold(true).Margin(1, 0, 0, 2).Foreground(lipgloss.Color("#6C50FF"))
	argument := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#6C50FF"))
	flagStyle := lipgloss.NewStyle().Foreground(DimWhite).PaddingLeft(1)

	fmt.Println()
	fmt.Println("  CLI para estar al tanto de la ligapro de",
		lipgloss.JoinHorizontal(lipgloss.Left, ecu.Render(), ad.Render(), or.Render()), "🇪🇨")

	fmt.Println(title.Render(strings.ToUpper("Uso")))
	t.Row(flagStyle.Render("-h"), "Mostrar este mensaje de ayuda")
	t.Row(flagStyle.Render("-tabla"), "Tabla de posiciones")
	t.Row(flagStyle.Render("-resultados"), "Últimos resultados de la fecha")
	t.Row(flagStyle.Render("-goleadores"), "Tabla de goleadores")
	t.Row(lipgloss.JoinHorizontal(lipgloss.Left, flagStyle.Render("-fecha"), argument.Render(" <número>")), "Partidos de dicha fecha")
	fmt.Println(t)

}
