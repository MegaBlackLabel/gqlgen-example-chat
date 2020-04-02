package graph

import (
	"context"
	"sync"

	"github.com/99designs/gqlgen/graphql"
	"github.com/MegaBlackLabel/gqlgen-example-chat/graph/generated"
	"github.com/MegaBlackLabel/gqlgen-example-chat/graph/model"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Rooms map[string]*model.Chatroom
	mu    sync.Mutex // nolint: structcheck
}

func New() generated.Config {
	return generated.Config{
		Resolvers: &Resolver{
			Rooms: map[string]*model.Chatroom{},
		},
		Directives: generated.DirectiveRoot{
			User: func(ctx context.Context, obj interface{}, next graphql.Resolver, username string) (res interface{}, err error) {
				return next(context.WithValue(ctx, "username", username))
			},
		},
	}
}
