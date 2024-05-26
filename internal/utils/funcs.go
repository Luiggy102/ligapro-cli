package utils

import (
	"fmt"
	"os"

	"github.com/charmbracelet/lipgloss"
)

func ErrMsg(str string) {
	errBox := lipgloss.NewStyle().Background(Red).Margin(1)
	msj := lipgloss.NewStyle().Margin(1).Italic(true).Foreground(Gray)
	fmt.Println(lipgloss.JoinHorizontal(lipgloss.Left,
		errBox.Render("ERROR"),
		msj.Render(str),
	))
	os.Exit(1)
}

func Title() {
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
