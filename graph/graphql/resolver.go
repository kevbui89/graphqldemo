package graphql

import (
	"com.example/graphql/graph/domain"
)

//go:generate go run github.com/99designs/gqlgen generate

// The Resolver needs to hold the Repositories you will use to query tables
type Resolver struct {
	Domain *domain.Domain
}
