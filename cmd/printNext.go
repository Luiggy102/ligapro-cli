package cmd

import (
	"regexp"
	"time"

	"github.com/Luiggy102/ligapro-cli/internal/requests"
	"github.com/Luiggy102/ligapro-cli/internal/utils"
)

func PrintNext() {
	fixture := requests.GetFixture()
	filas := [][]string{}
	partidosAMostrar := 8
	for i := 0; i < len(fixture.Result)-1; i++ {
		found, err := regexp.MatchString(":", fixture.Result[i])
		if err != nil {
			utils.ErrMsg("Error al filtrar los datos para resultados.\nInténtalo de nuevo.")
		}
		if found {
			if partidosAMostrar <= 0 {
				break
			}
			parsedTime, err := time.Parse("02/01 15:04", fixture.Result[i])
			if err != nil {
				utils.ErrMsg("Error al cambiar el tiempo")
			}
			diff := parsedTime.Add(time.Hour * -7)
			filas = append(filas, []string{
				fixture.Home[i], diff.Format("02/Jan 15:04"), fixture.Visit[i],
			})
			partidosAMostrar--
		}
	}
	const width = 80
	utils.PrintTitle("Siguientes partidos LigaPro", width)
	utils.PrintTable(width, nil, filas...)
	utils.LastUpdate()
}
