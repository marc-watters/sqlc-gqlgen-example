package gqlgen

import "github.com/marc-watters/sqlc-gqlgen-example/v2/pgx"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Repository pgx.Repository
}
