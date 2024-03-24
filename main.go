package main

import (
	"fmt"

	"github.com/Luiggy102/ligapro-cli/internal/requests"
)

func main() {
	fmt.Println(requests.GetTeams())
	fmt.Println(requests.GetStats())
}
