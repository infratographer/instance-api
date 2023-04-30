package graphapi

import (
	"context"
	"encoding/json"

	"go.infratographer.com/instance-api/internal/ent/generated"
	"go.infratographer.com/x/gidx"
)

type Location struct {
	ID        gidx.PrefixedID               `json:"id"`
	Instances *generated.InstanceConnection `json:"instances"`
}

func (Location) IsEntity() {}

type Tenant struct {
	ID                gidx.PrefixedID                       `json:"id"`
	Instances         *generated.InstanceConnection         `json:"instances"`
	InstanceProviders *generated.InstanceProviderConnection `json:"instanceProviders"`
}

func (Tenant) IsEntity() {}

type InstanceMetadataWriter interface {
	SetData(json.RawMessage) InstanceMetadataWriter
	Save(context.Context) (*generated.InstanceMetadata, error)
}
