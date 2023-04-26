package graph

// External models listed below. This allows us to mark the fields on the external
// models. We don't include the connections (Location -> Instances) as these being
// missing helps gqlgen understand that these are functions that we need to create.

// type Tenant struct {
// 	ID string `json:"id"`
// }

// func (Tenant) IsEntity() {}

// type Location struct {
// 	ID string `json:"id"`
// }

// func (Location) IsEntity() {}

// type Node interface {
// 	IsNode()
// }

// type Instance struct {
// 	ID        string            `json:"id"`
// 	Name      string            `json:"name"`
// 	Tenant    *Tenant           `json:"tenant"`
// 	Location  *Location         `json:"location"`
// 	Provider  *InstanceProvider `json:"provider"`
// 	CreatedAt time.Time         `json:"createdAt"`
// 	UpdatedAt time.Time         `json:"updatedAt"`
// 	DeletedAt *time.Time        `json:"deletedAt,omitempty"`
// }

// func (Instance) IsEntity() {}
// func (Instance) IsNode()   {}

// type InstanceProvider struct {
// 	ID        string              `json:"id"`
// 	Name      string              `json:"name"`
// 	Tenant    *Tenant             `json:"tenant"`
// 	CreatedAt time.Time           `json:"createdAt"`
// 	UpdatedAt time.Time           `json:"updatedAt"`
// 	DeletedAt *time.Time          `json:"deletedAt,omitempty"`
// 	Instances *InstanceConnection `json:"instances,omitempty"`
// }
