package ports

import (
	"context"
	"github.com/kayceeDev/caspa-events/ent-go/ent"
)

type IDataBase interface {
	Client() *ent.Client
	Close() error
	WithTx(ctx context.Context, fn func(tx *ent.Tx) error) error
}
