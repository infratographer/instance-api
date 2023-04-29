package schema

import (
	"time"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"go.infratographer.com/x/idx"
)

// InstanceProvider holds the schema definition for the InstanceProvider entity.
type InstanceProvider struct {
	ent.Schema
}

func (InstanceProvider) Mixin() []ent.Mixin {
	return []ent.Mixin{
		idx.PrimaryKeyMixin(idPrefixInstanceProvider),
	}
}

// Fields of the InstanceProvider.
func (InstanceProvider) Fields() []ent.Field {
	return []ent.Field{
		field.Text("name").
			NotEmpty().
			Annotations(
				entgql.OrderField("NAME"),
			),
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
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate()),
	}
}
