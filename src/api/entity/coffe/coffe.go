package entityCoffe

import (
	"context"

	coffeDomain "github.com/rayzalzero/go-sukha/src/domain/coffe"
	"golang.org/x/sync/errgroup"
)

func (a *coffeEntity) GetListCoffe(c context.Context, offset, limit int, search, sort string) (res []coffeDomain.ListCoffe, count int, err error) {
	_, ctx := errgroup.WithContext(c)
	res, count, err = a.repo.GetListCoffe(ctx, offset, limit, search, sort)
	if err != nil {
		return nil, count, err
	}
	return
}
