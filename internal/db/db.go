package db

import (
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	NewMysql,
)

// Data .
type Data struct {
	// TODO wrapped database client
}
