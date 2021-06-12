package entityCoffe

import (
	"time"

	coffeDomain "github.com/rayzalzero/go-sukha/src/domain/coffe"
)

type coffeEntity struct {
	repo    coffeDomain.Repo
	timeout time.Duration
}

func InitCoffeEntity(a coffeDomain.Repo, t time.Duration) coffeDomain.Entity {
	return &coffeEntity{
		repo:    a,
		timeout: t,
	}
}
