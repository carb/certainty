package entities

type FactionName string

const (
	NoFaction   FactionName = ""
	RedFaction  FactionName = "RED"
	BlueFaction FactionName = "BLUE"
)

type Faction struct {
	Name           FactionName     `json:"name"`
	Kingdom        []TerritoryName `json:"kingdom"`
	RedStarCount   int             `json:"red_star_count"`
	Coins          int             `json:"coins"`
	TerritoryCards []TerritoryCard `json:"territory_cards"`
}
