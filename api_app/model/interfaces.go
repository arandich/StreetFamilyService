package model

import "golang.org/x/net/context"

type UserService interface {
	Get(ctx context.Context, id int32) (*User, error)
}
