package requests

import (
	t "github.com/Luiggy102/ligapro-cli/types"
	"github.com/gocolly/colly"
)

// function that returns a slice of teams stats (position table)
func getStats() (stats []t.Stats, err error) {

	c := colly.NewCollector(
		colly.AllowedDomains("www.espn.com", "espn.com", "espndeportes.espn.com"),
	)

	c.OnError(func(r *colly.Response, errRequest error) {
		errRequest = err
	})
	if err != nil {
		return nil, err
	}

	// 1. get the raw numbers
	nums := []string{}

	c.OnHTML(".stat-cell", func(h *colly.HTMLElement) {
		nums = append(nums, h.Text)
	})

	c.Visit("https://www.espn.com/soccer/standings/_/league/ecu.1")

	// 2. orgnanize the data of every team
	var data1, data2, data3, data4, data5, data6, data7, data8, data9, data10, data11, data12, data13, data14, data15, data16 []string

	for i, v := range nums {
		i = i + 1
		switch {
		case i <= 8: // team 1
			data1 = append(data1, v)
		case i > 8 && i <= 8*2: // team 2
			data2 = append(data2, v)
		case i > 8*2 && i <= 8*3: // etc
			data3 = append(data3, v)
		case i > 8*3 && i <= 8*4:
			data4 = append(data4, v)
		case i > 8*4 && i <= 8*5:
			data5 = append(data5, v)
		case i > 8*5 && i <= 8*6:
			data6 = append(data6, v)
		case i > 8*6 && i <= 8*7:
			data7 = append(data7, v)
		case i > 8*7 && i <= 8*8:
			data8 = append(data8, v)
		case i > 8*8 && i <= 8*9:
			data9 = append(data9, v)
		case i > 8*9 && i <= 8*10:
			data10 = append(data10, v)
		case i > 8*10 && i <= 8*11:
			data11 = append(data11, v)
		case i > 8*11 && i <= 8*12:
			data12 = append(data12, v)
		case i > 8*12 && i <= 8*13:
			data13 = append(data13, v)
		case i > 8*13 && i <= 8*14:
			data14 = append(data14, v)
		case i > 8*14 && i <= 8*15:
			data15 = append(data15, v)
		case i > 8*15 && i <= 8*16:
			data16 = append(data16, v)
		}
	}

	// 3. make stat struc for all the teams
	var stat1, stat2, stat3, stat4, stat5, stat6, stat7, stat8, stat9, stat10, stat11, stat12, stat13, stat14, stat15, stat16 t.Stats

	for i, v := range data1 {
		switch i {
		case 0:
			stat1.GamesPlayed = v
		case 1:
			stat1.Wins = v
		case 2:
			stat1.Draws = v
		case 3:
			stat1.Loses = v
		case 4:
			stat1.GoalsFor = v
		case 5:
			stat1.GoalsAgainst = v
		case 6:
			stat1.GoalDifference = v
		case 7:
			stat1.Points = v
		}
	}
	stats = append(stats, stat1)

	for i, v := range data2 {
		switch i {
		case 0:
			stat2.GamesPlayed = v
		case 1:
			stat2.Wins = v
		case 2:
			stat2.Draws = v
		case 3:
			stat2.Loses = v
		case 4:
			stat2.GoalsFor = v
		case 5:
			stat2.GoalsAgainst = v
		case 6:
			stat2.GoalDifference = v
		case 7:
			stat2.Points = v
		}
	}
	stats = append(stats, stat2)

	for i, v := range data3 {
		switch i {
		case 0:
			stat3.GamesPlayed = v
		case 1:
			stat3.Wins = v
		case 2:
			stat3.Draws = v
		case 3:
			stat3.Loses = v
		case 4:
			stat3.GoalsFor = v
		case 5:
			stat3.GoalsAgainst = v
		case 6:
			stat3.GoalDifference = v
		case 7:
			stat3.Points = v
		}
	}
	stats = append(stats, stat3)

	for i, v := range data4 {
		switch i {
		case 0:
			stat4.GamesPlayed = v
		case 1:
			stat4.Wins = v
		case 2:
			stat4.Draws = v
		case 3:
			stat4.Loses = v
		case 4:
			stat4.GoalsFor = v
		case 5:
			stat4.GoalsAgainst = v
		case 6:
			stat4.GoalDifference = v
		case 7:
			stat4.Points = v
		}
	}
	stats = append(stats, stat4)

	for i, v := range data5 {
		switch i {
		case 0:
			stat5.GamesPlayed = v
		case 1:
			stat5.Wins = v
		case 2:
			stat5.Draws = v
		case 3:
			stat5.Loses = v
		case 4:
			stat5.GoalsFor = v
		case 5:
			stat5.GoalsAgainst = v
		case 6:
			stat5.GoalDifference = v
		case 7:
			stat5.Points = v
		}
	}
	stats = append(stats, stat5)

	for i, v := range data6 {
		switch i {
		case 0:
			stat6.GamesPlayed = v
		case 1:
			stat6.Wins = v
		case 2:
			stat6.Draws = v
		case 3:
			stat6.Loses = v
		case 4:
			stat6.GoalsFor = v
		case 5:
			stat6.GoalsAgainst = v
		case 6:
			stat6.GoalDifference = v
		case 7:
			stat6.Points = v
		}
	}
	stats = append(stats, stat6)

	for i, v := range data7 {
		switch i {
		case 0:
			stat7.GamesPlayed = v
		case 1:
			stat7.Wins = v
		case 2:
			stat7.Draws = v
		case 3:
			stat7.Loses = v
		case 4:
			stat7.GoalsFor = v
		case 5:
			stat7.GoalsAgainst = v
		case 6:
			stat7.GoalDifference = v
		case 7:
			stat7.Points = v
		}
	}
	stats = append(stats, stat7)

	for i, v := range data8 {
		switch i {
		case 0:
			stat8.GamesPlayed = v
		case 1:
			stat8.Wins = v
		case 2:
			stat8.Draws = v
		case 3:
			stat8.Loses = v
		case 4:
			stat8.GoalsFor = v
		case 5:
			stat8.GoalsAgainst = v
		case 6:
			stat8.GoalDifference = v
		case 7:
			stat8.Points = v
		}
	}
	stats = append(stats, stat8)

	for i, v := range data9 {
		switch i {
		case 0:
			stat9.GamesPlayed = v
		case 1:
			stat9.Wins = v
		case 2:
			stat9.Draws = v
		case 3:
			stat9.Loses = v
		case 4:
			stat9.GoalsFor = v
		case 5:
			stat9.GoalsAgainst = v
		case 6:
			stat9.GoalDifference = v
		case 7:
			stat9.Points = v
		}
	}
	stats = append(stats, stat9)

	for i, v := range data10 {
		switch i {
		case 0:
			stat10.GamesPlayed = v
		case 1:
			stat10.Wins = v
		case 2:
			stat10.Draws = v
		case 3:
			stat10.Loses = v
		case 4:
			stat10.GoalsFor = v
		case 5:
			stat10.GoalsAgainst = v
		case 6:
			stat10.GoalDifference = v
		case 7:
			stat10.Points = v
		}
	}
	stats = append(stats, stat10)

	for i, v := range data11 {
		switch i {
		case 0:
			stat11.GamesPlayed = v
		case 1:
			stat11.Wins = v
		case 2:
			stat11.Draws = v
		case 3:
			stat11.Loses = v
		case 4:
			stat11.GoalsFor = v
		case 5:
			stat11.GoalsAgainst = v
		case 6:
			stat11.GoalDifference = v
		case 7:
			stat11.Points = v
		}
	}
	stats = append(stats, stat11)

	for i, v := range data12 {
		switch i {
		case 0:
			stat12.GamesPlayed = v
		case 1:
			stat12.Wins = v
		case 2:
			stat12.Draws = v
		case 3:
			stat12.Loses = v
		case 4:
			stat12.GoalsFor = v
		case 5:
			stat12.GoalsAgainst = v
		case 6:
			stat12.GoalDifference = v
		case 7:
			stat12.Points = v
		}
	}
	stats = append(stats, stat12)

	for i, v := range data13 {
		switch i {
		case 0:
			stat13.GamesPlayed = v
		case 1:
			stat13.Wins = v
		case 2:
			stat13.Draws = v
		case 3:
			stat13.Loses = v
		case 4:
			stat13.GoalsFor = v
		case 5:
			stat13.GoalsAgainst = v
		case 6:
			stat13.GoalDifference = v
		case 7:
			stat13.Points = v
		}
	}
	stats = append(stats, stat13)

	for i, v := range data14 {
		switch i {
		case 0:
			stat14.GamesPlayed = v
		case 1:
			stat14.Wins = v
		case 2:
			stat14.Draws = v
		case 3:
			stat14.Loses = v
		case 4:
			stat14.GoalsFor = v
		case 5:
			stat14.GoalsAgainst = v
		case 6:
			stat14.GoalDifference = v
		case 7:
			stat14.Points = v
		}
	}
	stats = append(stats, stat14)

	for i, v := range data15 {
		switch i {
		case 0:
			stat15.GamesPlayed = v
		case 1:
			stat15.Wins = v
		case 2:
			stat15.Draws = v
		case 3:
			stat15.Loses = v
		case 4:
			stat15.GoalsFor = v
		case 5:
			stat15.GoalsAgainst = v
		case 6:
			stat15.GoalDifference = v
		case 7:
			stat15.Points = v
		}
	}
	stats = append(stats, stat15)

	for i, v := range data16 {
		switch i {
		case 0:
			stat16.GamesPlayed = v
		case 1:
			stat16.Wins = v
		case 2:
			stat16.Draws = v
		case 3:
			stat16.Loses = v
		case 4:
			stat16.GoalsFor = v
		case 5:
			stat16.GoalsAgainst = v
		case 6:
			stat16.GoalDifference = v
		case 7:
			stat16.Points = v
		}
	}
	stats = append(stats, stat16)

	return stats, nil
}
