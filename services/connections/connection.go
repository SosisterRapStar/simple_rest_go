package connections

import (
	"context"
	"first-proj/appconfig"
)

type Connection interface {
	Open(config appconfig.Config) error
	Close(ctx context.Context) error
}
