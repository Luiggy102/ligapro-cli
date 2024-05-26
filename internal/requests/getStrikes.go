package requests

import (
	"strings"

	"github.com/Luiggy102/ligapro-cli/internal/utils"
	"github.com/Luiggy102/ligapro-cli/types"
	"github.com/gocolly/colly"
)

func GetStrikers() types.Striker {
	var nombre, equipo, goles []string
	c := colly.NewCollector(
		colly.AllowedDomains("www.tycsports.com", "tycsports.com"),
	)
	c.OnError(func(_ *colly.Response, err error) {
		if err != nil {
			utils.ErrMsg("Error al traer datos de goleadores.\nInténtalo de nuevo.")
		}
	})
	c.OnHTML("td.pos.jugador", func(h *colly.HTMLElement) {
		nombre = append(nombre, strings.TrimSpace(h.Text))
	})
	c.OnHTML("td.goles", func(h *colly.HTMLElement) {
		goles = append(goles, strings.TrimSpace(h.Text))
	})
	c.OnHTML("td.escudo", func(h *colly.HTMLElement) {
		equipo = append(equipo, strings.TrimSpace(h.ChildAttr("img", "alt")))
	})
	c.Visit("https://www.tycsports.com/estadisticas/futbol-de-ecuador/tabla-de-goleadores.html")
	return types.Striker{Name: nombre, Team: equipo, Goals: goles}
}
