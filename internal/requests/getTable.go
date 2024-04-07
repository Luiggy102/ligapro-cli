package requests

import (
	t "github.com/Luiggy102/ligapro-cli/types"
)

func GetTable() ([]t.Team, error) {

	var teams, errT = getTeams()
	if errT != nil {
		return nil, errT
	}
	var stats, errS = getStats()
	if errS != nil {
		return nil, errS
	}

	teams[0].Stats = stats[0]
	teams[1].Stats = stats[1]
	teams[2].Stats = stats[2]
	teams[3].Stats = stats[3]
	teams[4].Stats = stats[4]
	teams[5].Stats = stats[5]
	teams[6].Stats = stats[6]
	teams[7].Stats = stats[7]
	teams[8].Stats = stats[8]
	teams[9].Stats = stats[9]
	teams[10].Stats = stats[10]
	teams[11].Stats = stats[11]
	teams[12].Stats = stats[12]
	teams[13].Stats = stats[13]
	teams[14].Stats = stats[14]
	teams[15].Stats = stats[15]

	return teams, nil
}
