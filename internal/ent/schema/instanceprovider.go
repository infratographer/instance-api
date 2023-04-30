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

// InstanceProvider holds the schema definition for the InstanceProvider entity.
type InstanceProvider struct {
	ent.Schema
}

// Mixin for InstanceProvider.
func (InstanceProvider) Mixin() []ent.Mixin {
	return []ent.Mixin{
		entx.TimestampsMixin{},
	}
}

// Fields of the InstanceProvider.
func (InstanceProvider) Fields() []ent.Field {
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
				// entgql.Type("ID")
				entgql.Skip(entgql.SkipWhereInput, entgql.SkipMutationUpdateInput),
			),
	}
}

// Indexed for InstanceMetadata
func (InstanceProvider) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("tenant_id"),
	}
}

// Edges of the InstanceProvider.
func (InstanceProvider) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("instances", Instance.Type).
			Ref("instance_provider").
			Annotations(
				entgql.RelayConnection(),
			),
	}
}

func (InstanceProvider) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entx.GraphKeyDirective("id"),
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}
