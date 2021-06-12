package coffeDomain

import "context"

// abstact class / kontrak function
type Entity interface {
	GetListCoffe(ctx context.Context, offset, limit int, search, sort string) (res []ListCoffe, count int, err error)
}

type Repo interface {
	GetListCoffe(ctx context.Context, offset, limit int, search, sort string) (res []ListCoffe, count int, err error)
}
