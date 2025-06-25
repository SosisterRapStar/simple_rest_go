// This package contains connections to third parties
// Abstraction on connection is necessary for Close/Open interface wich is used for gracefull shutdown
// It helps to close connections and connection pools to permanent stores, cache and elastic search stores

package connections

// import (
// 	"context"
// 	"first-proj/appconfig"
// )

// type Connection interface {
// 	Open(config appconfig.Config) error
// 	Close(ctx context.Context) error
// }
