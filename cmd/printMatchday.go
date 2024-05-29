package cmd

import (
	"fmt"
	"regexp"
	"time"

	"github.com/Luiggy102/ligapro-cli/internal/requests"
	"github.com/Luiggy102/ligapro-cli/internal/utils"
)

func PrintMatchday(matchday int) {
	if matchday < 1 || matchday > 30 {
		utils.WarningMsg("Número de fecha incorrecto.\nColoque una fecha válida.")
	}
	fixture := requests.GetFixture()
	fecha := 1
	filas := [][]string{}
	// lógica
	for i := 0; i < len(fixture.Result)-1; i++ {
		if i%8 == 0 {
			fecha++
		}
		if fecha == matchday+1 {
			// cambiar la fecha a horario ecuatoriano
			found, err := regexp.MatchString(":", fixture.Result[i])
			if err != nil {
				utils.ErrMsg("Error al cambiar el formato de la fecha")
			}
			if found {
				// 26/05 04:00
				parsedTime, err := time.Parse("02/01 15:04", fixture.Result[i])
				if err != nil {
					utils.ErrMsg("Error al cambiar el tiempo")
				}
				// 7 horas de diferencia
				diff := parsedTime.Add(time.Hour * -7)

				// agregar filas
				filas = append(filas, []string{
					fixture.Home[i], diff.Format("02/Jan 15:04"), fixture.Visit[i],
				})
			} else {
				// agregar filas
				filas = append(filas, []string{
					fixture.Home[i], fixture.Result[i], fixture.Visit[i],
				})
			}

		}
	}
	// imprimir
	const width = 75
	utils.PrintTitle(fmt.Sprintf("Partidos de la fecha %d LigaPro %d", matchday, time.Now().Year()), width)
	utils.PrintTable(width, nil, filas...)
	utils.LastUpdate()
}
