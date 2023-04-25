package graph

type Tenant struct {
	ID string `json:"id"`
}

func (Tenant) IsEntity() {}

type Location struct {
	ID string `json:"id"`
}

func (Location) IsEntity() {}
