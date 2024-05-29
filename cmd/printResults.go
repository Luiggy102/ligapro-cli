package cmd

import (
	"fmt"
	"regexp"
	"time"

	"github.com/Luiggy102/ligapro-cli/internal/requests"
	"github.com/Luiggy102/ligapro-cli/internal/utils"
)

func PrintResults() {
	fixture := requests.GetFixture()
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
	// Imprimir
	const width = 65
	utils.PrintTitle(fmt.Sprintf("Últimos partidos jugados LigaPro %v", time.Now().Year()), width)
	utils.PrintTable(width, nil, filas...)
	utils.LastUpdate()
}
