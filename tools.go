//go:build tools
// +build tools

package tools

import (
	_ "github.com/99designs/gqlgen"
	_ "github.com/99designs/gqlgen/graphql/introspection"
	_ "github.com/infratographer/sqlboiler-crdb/v4"
	_ "github.com/volatiletech/sqlboiler/v4"
)
