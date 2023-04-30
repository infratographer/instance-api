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

// Instance holds the schema definition for the Instance entity.
type Instance struct {
	ent.Schema
}

func (Instance) Mixin() []ent.Mixin {
	return []ent.Mixin{
		entx.TimestampsMixin{},
	}
}

// Fields of the Instance.
func (Instance) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			GoType(gidx.PrefixedID("")).
			DefaultFunc(func() gidx.PrefixedID { return gidx.MustNewID(idPrefixInstance) }).
			Unique().
			Immutable(),
		field.Text("name").
			NotEmpty().
			Annotations(
				entgql.OrderField("NAME"),
			),
		field.String("tenant_id").
			GoType(gidx.PrefixedID("")).
			Immutable().
			Annotations(
				entgql.Skip(entgql.SkipWhereInput, entgql.SkipMutationUpdateInput),
			),
		field.String("location_id").
			GoType(gidx.PrefixedID("")).
			Immutable().
			Annotations(
				entgql.Skip(entgql.SkipWhereInput, entgql.SkipMutationUpdateInput),
			),
		field.String("instance_provider_id").
			GoType(gidx.PrefixedID("")).
			Immutable().
			Annotations(
				entgql.Skip(entgql.SkipWhereInput, entgql.SkipMutationUpdateInput),
			),
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

func (Instance) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("instance_provider_id"),
		index.Fields("location_id"),
		index.Fields("tenant_id"),
	}
}

func (Instance) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entx.GraphKeyDirective("id"),
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}
