package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"go.infratographer.com/x/entx"
	"go.infratographer.com/x/idx"
)

// InstanceMetadata holds the schema definition for the InstanceMetadata entity.
type InstanceMetadata struct {
	ent.Schema
}

func (InstanceMetadata) Mixin() []ent.Mixin {
	return []ent.Mixin{
		idx.PrimaryKeyMixin(idPrefixInstanceMetadata),
		entx.NamespacedDataMixin{},
		entx.TimestampsMixin{},
	}
}

// Fields of the InstanceMetadata.
func (InstanceMetadata) Fields() []ent.Field {
	return []ent.Field{
		field.String("instance_id").
			GoType(idx.PrefixedID("")).
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
	}
}
