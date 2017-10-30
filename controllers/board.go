package controllers

import (
	"fmt"

	e "github.com/carb/certainty/entities"
)

type BoardController interface {
	PlaceHQ(faction e.FactionName, territory e.TerritoryName) error
	PlaceTroops(faction e.FactionName, territory e.TerritoryName, count int) error
	MoveTroops(faction e.FactionName, from e.TerritoryName, to e.TerritoryName, count int) error
}

type boardControllerImpl struct {
	board e.Board
}

func NewBoardController(b e.Board) BoardController {
	return boardControllerImpl{board: b}
}

func (bc boardControllerImpl) PlaceHQ(faction e.FactionName, territory e.TerritoryName) error {
	if _, ok := bc.board.TerritoryMap[territory]; !ok {
		return fmt.Errorf("Unable to place HQ, territory %v not on this board.", territory)
	}

	if bc.board.TerritoryMap[territory].HQ != e.NoFaction {
		return fmt.Errorf("Unable to place HQ, territory %v occupied.", territory)
	}

	for _, neighbor := range bc.board.AdjacencyList[territory] {
		if bc.board.TerritoryMap[neighbor].HQ != e.NoFaction {
			return fmt.Errorf("Unable to place HQ at %v, neighboring territory %v occupied.", territory, neighbor)
		}
	}

	bc.board.TerritoryMap[territory].HQ = faction
	return nil
}

func (bc boardControllerImpl) PlaceTroops(faction e.FactionName, territory e.TerritoryName, count int) error {
	if bc.board.TerritoryMap[territory].ControlledBy != faction {
		return fmt.Errorf("Unable to place troops, %v does not control %v.", faction, territory)
	}

	bc.board.TerritoryMap[territory].Strength += count
	return nil
}

func (bc boardControllerImpl) MoveTroops(faction e.FactionName, from e.TerritoryName, to e.TerritoryName, count int) error {
	if bc.board.TerritoryMap[from].ControlledBy != faction {
		return fmt.Errorf("Unable to move troops, %v does not control starting territory %v.", faction, from)
	}

	if cb := bc.board.TerritoryMap[to].ControlledBy; cb != e.NoFaction {
		return fmt.Errorf("Unable to move troops, destination territory %v occupied by %v.", to, cb)
	}

	bc.board.TerritoryMap[from].Strength -= count
	bc.board.TerritoryMap[to].ControlledBy = faction
	bc.board.TerritoryMap[to].Strength = count
	bc.board.Factions[faction].Kingdom = append(bc.board.Factions[faction].Kingdom, to)
	return nil
}
