package utils

import (
	"fmt"
	"os"
	"time"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
)

func PrintTitle(title string, width int) {
	tStyle := lipgloss.NewStyle().
		Width(width).
		Align(lipgloss.Center).
		Bold(true).
		Foreground(lipgloss.Color(Yellow)).
		Padding(1, 0)
	fmt.Println(tStyle.Render(title))
}

func PrintVisitingMsg(url string) {
	msgStyle := lipgloss.NewStyle().
		Foreground(LightBlue).
		Underline(true).
		Italic(true)
	fmt.Println("Visitando:", msgStyle.Render(url))
}

func PrintTable(width int, header []string, rows ...[]string) {
	re := lipgloss.NewRenderer(os.Stdout)
	evenRowStyle := lipgloss.NewStyle().Foreground(DimWhite)
	oddRowStyle := lipgloss.NewStyle()
	t := table.New().
		Width(width).
		Headers(header...).
		Rows(rows...).
		Border(lipgloss.RoundedBorder()).
		BorderStyle(re.NewStyle().Foreground(Skyblue)).
		StyleFunc(func(row, _ int) lipgloss.Style {
			switch {
			case row == 0:
				return oddRowStyle
			case row%2 == 0:
				return evenRowStyle
			default:
				return oddRowStyle
			}
		})
	fmt.Println(t)
}

func LastUpdate() {
	uStyle := lipgloss.NewStyle().
		Italic(true).
		Foreground(DimWhite)
	fmt.Println(uStyle.Render(fmt.Sprintf("\nActualizado: %v", time.Now().Format("02/01/2006"))))
}

func ErrMsg(str string) {
	errBox := lipgloss.NewStyle().Background(Red).Margin(1)
	msj := lipgloss.NewStyle().Margin(1).Italic(true).Foreground(Gray)
	fmt.Println(lipgloss.JoinHorizontal(lipgloss.Left,
		errBox.Render("ERROR"),
		msj.Render(str),
	))
	os.Exit(1)
}

func WarningMsg(str string) {
	warBox := lipgloss.NewStyle().Background(DimYellow).Margin(1).Foreground(Black)
	msj := lipgloss.NewStyle().Margin(1).Italic(true).Foreground(Gray)
	fmt.Println(lipgloss.JoinHorizontal(lipgloss.Left,
		warBox.Render("ADVERTENCIA"),
		msj.Render(str),
	))
	os.Exit(1)
}

func Welcome() {
	t := lipgloss.NewStyle().
		Background(DimYellow).
		Bold(true).
		Foreground(Black).
		SetString("Ligapro-cli").
		Margin(1)
	b := lipgloss.NewStyle().
		Margin(1).
		Bold(true).
		SetString("CLI para estar al tanto de la liga pro de Ecuador 🇪🇨")
	f := lipgloss.NewStyle().
		Italic(true).
		Foreground(Gray).
		SetString("Escriba \"ligapro -h\" para obtener ayuda")
	fmt.Println(lipgloss.JoinVertical(lipgloss.Center,
		lipgloss.JoinHorizontal(lipgloss.Center,
			t.Render(),
			b.Render(),
		),
		f.Render(),
	))
}
