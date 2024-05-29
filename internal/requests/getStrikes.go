package requests

import (
	"fmt"
	"strings"

	"github.com/Luiggy102/ligapro-cli/internal/utils"
	"github.com/Luiggy102/ligapro-cli/types"
	"github.com/gocolly/colly"
)

func GetStrikers() types.Striker {
	var nombre, equipo, goles []string
	var intentos uint8 = 4
	c := colly.NewCollector(colly.AllowedDomains("www.tycsports.com", "tycsports.com"))
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
	c.OnHTML("td.pos.jugador", func(h *colly.HTMLElement) { nombre = append(nombre, strings.TrimSpace(h.Text)) })
	c.OnHTML("td.goles", func(h *colly.HTMLElement) { goles = append(goles, strings.TrimSpace(h.Text)) })
	c.OnHTML("td.escudo", func(h *colly.HTMLElement) { equipo = append(equipo, strings.TrimSpace(h.ChildAttr("img", "alt"))) })
	c.Visit("https://www.tycsports.com/estadisticas/futbol-de-ecuador/tabla-de-goleadores.html")
	return types.Striker{Name: nombre, Team: equipo, Goals: goles}
}
