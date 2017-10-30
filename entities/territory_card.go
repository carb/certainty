package entities

type TerritoryCard struct {
	Territory TerritoryName `json:"territory_name"`
	Coins     int           `json:"coins"`
}
