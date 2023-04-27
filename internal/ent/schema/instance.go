package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"

	"go.infratographer.com/instance-api/xthings/entx"
	"go.infratographer.com/instance-api/xthings/idx"
)

// Instance holds the schema definition for the Instance entity.
type Instance struct {
	ent.Schema
}

func (Instance) Mixin() []ent.Mixin {
	return []ent.Mixin{
		idx.PrimaryKeyMixin(idPrefixInstance),
		entx.ExternalEdge(
			entx.WithExternalEdgeType("Location"),
			entx.ExternalEdgeImmutable(true),
		),
		entx.ExternalEdge(
			entx.WithExternalEdgeType("Tenant"),
			entx.ExternalEdgeImmutable(true),
		),
		entx.TimestampsMixin{},
	}
}

// Fields of the Instance.
func (Instance) Fields() []ent.Field {
	return []ent.Field{
		field.Text("name").
			NotEmpty().
			Annotations(
				entgql.OrderField("NAME"),
			),
		field.String("instance_provider_id").
			GoType(idx.PrefixedID("")).
			Immutable(),
	}
}

// Edges of the Instance.
func (Instance) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("instance_provider", InstanceProvider.Type).
			Unique().
			Required().
			Immutable().
			Field("instance_provider_id").
			Annotations(),
		edge.From("metadata", InstanceMetadata.Type).
			Ref("instance").
			Annotations(
				entgql.RelayConnection(),
			),
	}
}

func (Instance) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate()),
	}
}
