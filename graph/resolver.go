package graph

import "github.com/go-pg/pg/v10"

// go run github.com/99designs/gqlgen generate
// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{
	DB  *pg.DB

}
