package schema

import (
	"encoding/json"
	"time"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// InstanceMetadata holds the schema definition for the InstanceMetadata entity.
type InstanceMetadata struct {
	ent.Schema
}

// Fields of the InstanceMetadata.
func (InstanceMetadata) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New),
		field.Text("namespace").
			NotEmpty().
			Annotations(
				entgql.OrderField("NAME"),
			),
		field.JSON("data", json.RawMessage{}).
			Annotations(),
		field.Time("created_at").
			Default(time.Now).
			Immutable().
			Annotations(
				entgql.OrderField("CREATED_AT"),
			),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now).
			Immutable().
			Annotations(
				entgql.OrderField("UPDATED_AT"),
			),
		field.UUID("instance_id", uuid.UUID{}).
			Immutable(),
	}
}

// Edges of the InstanceMetadata.
func (InstanceMetadata) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("instance", Instance.Type).
			Unique().
			Required().
			Immutable().
			Field("instance_id").
			Annotations(),
	}
}

func (InstanceMetadata) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate()),
		entgql.Directives(KeyDirective("id")),
	}
}
