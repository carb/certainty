package entities

type Territory struct {
	City      City          `json:"city"`
	Continent Continent     `json:"continent"`
	Name      TerritoryName `json:"name"`
	Scar      Scar          `json:"scar"`

	ControlledBy FactionName
	Strength     int
	HQ           FactionName
}

type City struct {
	Name              string `json:"name"`
	Population        int    `json:"population"`
	NamedBy           Player `json:"named_by"`
	FortifiedStrength int    `json:"fortified_strength"`
}

type Continent struct {
	Name         string `json:"name"`
	NamedBy      Player `json:"named_by"`
	InformalName string `json:"informal_name"`
}

type Scar string

const (
	AMMO_SHORTAGE Scar = "AMMO_SHORTAGE"
	BIOHAZARD     Scar = "BIOHAZARD"
	BUNKER        Scar = "BUNKER"
	FALLOUT       Scar = "FALLOUT"
	MERCENARY     Scar = "MERCENARY"
)

func NewTerritory(name TerritoryName) Territory {
	return Territory{Name: name}
}
