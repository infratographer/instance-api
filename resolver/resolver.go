package resolver

//
// This file will not be regenerated automatically.
//

import (
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/labstack/echo/v4"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"go.infratographer.com/instance-api/graph"
)

type Resolver struct {
	db boil.ContextExecutor
}

type Handler struct {
	r                 *Resolver
	graphqlHandler    http.Handler
	playgroundHandler http.Handler
}

func NewHandler(db boil.ContextExecutor) *Handler {
	h := &Handler{
		r: &Resolver{
			db: db,
		},
	}

	h.graphqlHandler = handler.NewDefaultServer(
		graph.NewExecutableSchema(
			graph.Config{
				Resolvers: h.r,
			},
		),
	)
	h.playgroundHandler = playground.Handler("GraphQL", "/query")

	return h
}

func (h *Handler) Routes(e *echo.Group) {
	e.POST("/query", func(c echo.Context) error {
		h.graphqlHandler.ServeHTTP(c.Response(), c.Request())
		return nil
	})

	e.GET("/playground", func(c echo.Context) error {
		h.playgroundHandler.ServeHTTP(c.Response(), c.Request())
		return nil
	})
}
