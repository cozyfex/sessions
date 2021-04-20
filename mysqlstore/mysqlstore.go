package mysqlstore

import (
	"database/sql"
	"github.com/gin-contrib/sessions"
	"github.com/srinathgs/mysqlstore"
)

type Store interface {
	sessions.Store
}

func NewStore(endpoint string, table string, path string, maxAge int, keyPairs ...[]byte) (Store, error) {
	s, err := mysqlstore.NewMySQLStore(endpoint, table, path, maxAge, keyPairs...)
	if err != nil {
		return nil, err
	}
	return &store{s}, nil
}

func NewStoreWithDB(db *sql.DB, table string, path string, maxAge int, keyPairs ...[]byte) (Store, error) {
	s, err := mysqlstore.NewMySQLStoreFromConnection(db, table, path, maxAge, keyPairs...)
	if err != nil {
		return nil, err
	}
	return &store{s}, nil
}

type store struct {
	*mysqlstore.MySQLStore
}

func (c *store) Options(options sessions.Options) {
	c.MySQLStore.Options = options.ToGorillaOptions()
}
