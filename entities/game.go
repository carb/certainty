package entities

type Game struct {
	TerritoryDeck     []TerritoryCard `json:"territory_deck"`
	EventDeck         []EventCard     `json:"event_deck"`
	MissionDeck       []MissionCard   `json:"mission_deck"`
	AvailableFactions []FactionName   `json:"available_factions"`
}
