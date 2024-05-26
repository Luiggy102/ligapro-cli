package cmd

import (
	"fmt"
	"time"

	"github.com/Luiggy102/ligapro-cli/internal/requests"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
)

func PrintStrikers() {

	goleadores := requests.GetStrikers()
	t := table.New().Headers("Jugador", "Goles", "Equipo").Width(55)
	filas := [][]string{}

	for i := 0; i < len(goleadores.Goals)-1; i++ {
		filas = append(filas, []string{
			goleadores.Name[i], goleadores.Goals[i], goleadores.Team[i],
		})
	}
	t.Rows(filas...)

	// Estilos
	hStyle := lipgloss.NewStyle().Width(55).Align(lipgloss.Center).Bold(true).Foreground(lipgloss.Color("230")).Padding(1, 0)
	t.BorderStyle(lipgloss.NewStyle().Foreground(lipgloss.Color("46")))

	// Imprimir
	h := fmt.Sprintf("Tabla de goleadores LigaPro %d", time.Now().Year())
	fmt.Println(hStyle.Render(h))
	fmt.Println(t)
	fmt.Println(fmt.Sprintf("\nActualizado: %v", time.Now().Format("02/01/2006")))
	fmt.Println("Datos obtenidos de: tycsports.com")
}
