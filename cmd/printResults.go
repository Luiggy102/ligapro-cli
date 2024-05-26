package cmd

import (
	"fmt"
	"regexp"
	"time"

	"github.com/Luiggy102/ligapro-cli/internal/requests"
	"github.com/Luiggy102/ligapro-cli/internal/utils"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
)

func PrintResults() {

	fixture := requests.GetFixture()
	t := table.New()
	filas := [][]string{}
	partidosAMostrar := 8

	// impimir 8 últimos partidos
	for i := len(fixture.Home) - 1; i >= 0; i-- {
		found, err := regexp.MatchString("-", fixture.Result[i])
		if err != nil {
			utils.ErrMsg("Error al filtrar los datos para resultados.\nInténtalo de nuevo.")
		}
		if found {
			if partidosAMostrar <= 0 {
				break
			}
			filas = append(filas, []string{
				fixture.Home[i],
				fixture.Result[i],
				fixture.Visit[i],
			})
			partidosAMostrar--
		}
	}
	t.Rows(filas...)

	// estilos
	hStyle := lipgloss.NewStyle().
		Foreground(utils.Yellow).
		Bold(true).
		Width(65).
		Align(lipgloss.Center).
		Padding(1, 0)

	t.Width(65).
		BorderStyle(lipgloss.NewStyle().
			Foreground(utils.DimYellow)).
		BorderColumn(false).BorderRow(true)

	h := fmt.Sprintf("Últimos partidos jugados LigaPro %v", time.Now().Year())
	u := fmt.Sprintf("Actualizado: %v", time.Now().Format("02/01/2006"))

	// Imprimir
	fmt.Println(hStyle.Render(h))
	fmt.Println(t)
	fmt.Println(u)
	fmt.Println("Datos obtenidos de: marca.com")
}
