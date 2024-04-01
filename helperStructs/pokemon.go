package helperStructs

type Pokemon struct {
	Ability string `json:"ability"`
	Moves   []Move `json:"moves"`
	Name    string `json:"name"`
}

type Move struct {
	Name string `json:"name"`
	Type string `json:"type"`
}
