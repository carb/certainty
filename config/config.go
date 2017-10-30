package config

import (
// e "github.com/carb/certainty/entities"
)

// Configuration holds the structure of yaml configuration.
type Configuration struct {
	GameName string `json:"game_name"`
	// Game     GameConfig  `json:"game"`
	Board BoardConfig `json:"board"`
}

// // GameConfig basics to deal with later.
// type GameConfig struct {
// 	TerritoryDeck     []e.TerritoryCard `json:"territory_deck"`
// 	EventDeck         []e.EventCard     `json:"event_deck"`
// 	MissionDeck       []e.MissionCard   `json:"mission_deck"`
// 	AvailableFactions []e.FactionName   `json:"available_factions"`
// }

type Border []string

// BoardConfig
type BoardConfig struct {
	Borders []Border `json:"borders"`
}
