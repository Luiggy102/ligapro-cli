package requests

import (
	"strings"

	"github.com/Luiggy102/ligapro-cli/internal/utils"
	"github.com/Luiggy102/ligapro-cli/types"
	"github.com/gocolly/colly"
)

func GetFixture() types.Fixture {
	var local, visitante, resultados []string
	c := colly.NewCollector(colly.AllowedDomains("marca.com", "www.marca.com"))
	c.OnError(func(_ *colly.Response, err error) {
		if err != nil {
			utils.ErrMsg("Error al traer información de próximos partidos.\nInténtalo de nuevo")
		}
	})
	c.OnHTML("td.local", func(h *colly.HTMLElement) {
		local = append(local, strings.TrimSpace(h.Text))
	})
	c.OnHTML("td.visitante", func(h *colly.HTMLElement) {
		visitante = append(visitante, strings.TrimSpace(h.Text))
	})
	c.OnHTML("td.resultado", func(h *colly.HTMLElement) {
		resultados = append(resultados, strings.TrimSpace(h.Text))
	})
	c.Visit("https://www.marca.com/futbol/ecuador/primera-etapa/calendario.html")
	return types.Fixture{Home: local, Visit: visitante, Result: resultados}
}

// fecha := 1
// for i, v := range local {
// 	if i%8 == 0 {
// 		fmt.Println("\tFecha ", fecha)
// 		fecha++
// 	}
// 	fmt.Println(v)
// }
