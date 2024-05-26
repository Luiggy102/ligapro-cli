package requests

import (
	"github.com/Luiggy102/ligapro-cli/internal/utils"
	t "github.com/Luiggy102/ligapro-cli/types"
	"github.com/gocolly/colly"
)

// function that returns a slice of teams (position table)
func getTeams() (teams []t.Team) {

	c := colly.NewCollector(
		colly.AllowedDomains("www.espn.com", "espn.com", "espndeportes.espn.com"),
	)

	c.OnError(func(_ *colly.Response, err error) {
		if err != nil {
			utils.ErrMsg("Error al traer información de Tabla.\nInténtalo de nuevo.")
		}
	})

	c.OnHTML(".Table__TD", func(h *colly.HTMLElement) {

		t := t.Team{
			Position: h.ChildText(".team-position.ml2.pr3"),
			Name:     h.ChildText(".hide-mobile"),
			Link:     h.ChildAttr("a", "href"),
		}
		teams = append(teams, t)
	})

	c.Visit("https://www.espn.com/soccer/standings/_/league/ecu.1")

	// only show 16 result, 16 teams

	var rTeams = []t.Team{}

	for i, v := range teams {
		if i == 16 {
			break
		}
		rTeams = append(rTeams, v)
		continue
	}

	return rTeams
}
