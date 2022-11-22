package model

import "golang.org/x/net/context"

type UserService interface {
	Post(ctx context.Context, id int32) (*User, error)
}

type CompanyService interface {
	Post(ctx context.Context, id int32) (*CompanySearch, error)
}

type CatalogService interface {
	Post(ctx context.Context, id int32) (*Catalog, error)
}
