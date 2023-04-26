package schema

import (
	"time"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/vektah/gqlparser/v2/ast"
)

// Instance holds the schema definition for the Instance entity.
type Instance struct {
	ent.Schema
}

// Fields of the Instance.
func (Instance) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New),
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
		field.UUID("instance_provider_id", uuid.UUID{}).
			Immutable(),
		field.UUID("tenant_id", uuid.UUID{}).
			Immutable().
			Annotations(
				entgql.Type("ID"),
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

func (Instance) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate()),
		entgql.Directives(KeyDirective("id")),
	}
}

// entgql.Directives(entgql.Deprecated("Use `description` instead.")),
//
//	),
//
// and the GraphQL type will be generated with the directive.
//
//	type Todo {
//		id: ID
//		text: String @deprecated(reason: "Use `description` instead.")
//
// Deprecated create `@deprecated` directive to apply on the field/type
func KeyDirective(fields string) entgql.Directive {
	var args []*ast.Argument
	if fields != "" {
		args = append(args, &ast.Argument{
			Name: "fields",
			Value: &ast.Value{
				Raw:  fields,
				Kind: ast.StringValue,
			},
		})
	}
	return entgql.NewDirective("key", args...)
}
