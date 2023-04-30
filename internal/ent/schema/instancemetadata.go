package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"go.infratographer.com/x/entx"
	"go.infratographer.com/x/gidx"
)

// InstanceMetadata holds the schema definition for the InstanceMetadata entity.
type InstanceMetadata struct {
	ent.Schema
}

func (InstanceMetadata) Mixin() []ent.Mixin {
	return []ent.Mixin{
		entx.NamespacedDataMixin{},
		entx.TimestampsMixin{},
	}
}

// Fields of the InstanceMetadata.
func (InstanceMetadata) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			GoType(gidx.PrefixedID("")).
			DefaultFunc(func() gidx.PrefixedID { return gidx.MustNewID(idPrefixInstance) }).
			Unique().
			Immutable(),
		field.String("instance_id").
			GoType(gidx.PrefixedID("")).
			Immutable(),
	}
}

// Indexed for InstanceMetadata
func (InstanceMetadata) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("instance_id"),
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
		entx.GraphKeyDirective("id"),
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationUpdate().Description("The fields used to set an annotation on an Instance")),
	}
}
