package requests

import (
	"fmt"
	"strings"

	"github.com/Luiggy102/ligapro-cli/internal/utils"
	"github.com/Luiggy102/ligapro-cli/types"
	"github.com/gocolly/colly"
)

func GetFixture() types.Fixture {
	var local, visitante, resultados []string
	var intentos uint8 = 4
	c := colly.NewCollector(colly.AllowedDomains("marca.com", "www.marca.com"))
	c.OnRequest(func(r *colly.Request) {
		utils.PrintVisitingMsg(r.URL.Hostname())
	})
	c.OnError(func(r *colly.Response, err error) {
		utils.WarningMsg(fmt.Sprintf("Intentanto de nuevo...\nStatus Code %d", r.StatusCode))
		if r.StatusCode == 0 {
			intentos--
			r.Request.Retry()
		}
		if err != nil || intentos == 0 {
			utils.ErrMsg("Error al traer datos de goleadores.\nInténtalo de nuevo.")
		}
	})
	c.OnHTML("td.local", func(h *colly.HTMLElement) { local = append(local, strings.TrimSpace(h.Text)) })
	c.OnHTML("td.visitante", func(h *colly.HTMLElement) { visitante = append(visitante, strings.TrimSpace(h.Text)) })
	c.OnHTML("td.resultado", func(h *colly.HTMLElement) { resultados = append(resultados, strings.TrimSpace(h.Text)) })
	c.Visit("https://www.marca.com/futbol/ecuador/primera-etapa/calendario.html")
	return types.Fixture{Home: local, Visit: visitante, Result: resultados}
}
