package entities

import (
	"github.com/carb/certainty/config"
)

type Board struct {
	TerritoryMap  map[TerritoryName]*Territory
	AdjacencyList map[TerritoryName][]TerritoryName
	Factions      map[FactionName]*Faction
}

func NewBoard(c *config.Configuration) Board {
	b := Board{
		TerritoryMap:  make(map[TerritoryName]*Territory),
		AdjacencyList: make(map[TerritoryName][]TerritoryName),
		Factions:      make(map[FactionName]*Faction),
	}

	for _, border := range c.Board.Borders {
		alpha := TerritoryName(border[0])
		omega := TerritoryName(border[1])

		// Update adjacency list
		if _, ok := b.AdjacencyList[alpha]; !ok {
			b.AdjacencyList[alpha] = make([]TerritoryName, 0)
		}
		b.AdjacencyList[alpha] = append(b.AdjacencyList[alpha], omega)

		if _, ok := b.AdjacencyList[omega]; !ok {
			b.AdjacencyList[omega] = make([]TerritoryName, 0)
		}
		b.AdjacencyList[omega] = append(b.AdjacencyList[omega], alpha)

		// Update the territory map
		if _, ok := b.TerritoryMap[alpha]; !ok {
			t := NewTerritory(alpha)
			b.TerritoryMap[alpha] = &t
		}

		if _, ok := b.TerritoryMap[omega]; !ok {
			t := NewTerritory(omega)
			b.TerritoryMap[omega] = &t
		}
	}

	return b
}
