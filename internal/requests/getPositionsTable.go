package requests

import (
	"fmt"
	"strings"
	"time"

	"github.com/Luiggy102/ligapro-cli/internal/utils"
	"github.com/Luiggy102/ligapro-cli/types"
	"github.com/gocolly/colly"
)

func GetPositionsTable(year int, season int) types.PostionsTable {
	// datos
	var nombres, pj, pg, pe, pp, gf, gc, gdif, pts []string
	var titulo string
	var data []string
	var intentos uint8 = 4
	var url string
	// tratar url con año
	if year != 0 && season != 0 {
		if (year >= 2014 && year <= int(time.Now().Year())) && (season >= 1 && season <= 2) {
			url = fmt.Sprintf("https://espndeportes.espn.com/futbol/posiciones/_/liga/ECU.1/tipodetemporada/%d/temporada/%d",
				season, year)
		} else {
			utils.WarningMsg("Coloque año mayor a 2014 y/o una etapa válida")
		}
	} else {
		// default, sin argumentos
		switch time.Now().Month() {
		case 1, 2, 3, 4, 5:
			season = 1
			url = fmt.Sprintf("https://espndeportes.espn.com/futbol/posiciones/_/liga/ECU.1/tipodetemporada/1/temporada/%d",
				time.Now().Year())
		case 6, 7, 8, 9, 10, 11, 12:
			season = 2
			url = fmt.Sprintf("https://espndeportes.espn.com/futbol/posiciones/_/liga/ECU.1/tipodetemporada/2/temporada/%d",
				time.Now().Year())
		}
	}
	// traer datos
	c := colly.NewCollector(colly.AllowedDomains("www.espn.com", "espn.com", "espndeportes.espn.com", "www.espndeportes.espn.com"))
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
			utils.ErrMsg("Error al traer datos de la tabla.\nInténtalo de nuevo.")
		}
	})
	c.OnHTML("span.hide-mobile", func(h *colly.HTMLElement) { nombres = append(nombres, strings.TrimSpace(h.Text)) })
	c.OnHTML(".stat-cell", func(h *colly.HTMLElement) { data = append(data, strings.TrimSpace(h.Text)) })
	c.OnHTML(".flex.justify-between.mt3.mb5.items-center", func(h *colly.HTMLElement) { titulo = strings.TrimSpace(h.Text) })
	c.Visit(url)
	titulo += fmt.Sprintf(" - Etapa %d", season)
	for i := 0; i < len(data); i++ {
		if i%8 == 0 {
			pj = append(pj, data[i])
			pg = append(pg, data[i+1])
			pe = append(pe, data[i+2])
			pp = append(pp, data[i+3])
			gf = append(gf, data[i+4])
			gc = append(gc, data[i+5])
			gdif = append(gdif, data[i+6])
			pts = append(pts, data[i+7])
		}
	}
	return types.PostionsTable{
		TableTitle:     titulo,
		Names:          nombres,
		GamesPlayed:    pj,
		Wins:           pg,
		Draws:          pe,
		Loses:          pp,
		GoalsFor:       gf,
		GoalsAgainst:   gc,
		GoalDifference: gdif,
		Points:         pts,
	}
}
