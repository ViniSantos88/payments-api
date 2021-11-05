package postgresql

import (
	"github.com/ViniSantos88/payments-api/framework/config"
	"github.com/jmoiron/sqlx"
)

var postgreSQLRepoImpl *PostgreSQLRepoImpl

// PostgreSQLRepoImpl is the implementation of Payments data interface
type PostgreSQLRepoImpl struct {
	//Database is the db driver
	Db *sqlx.DB
}

// GetPostgreSQLRepoImpl returns an instance of PostgreSQLRepoImpl
func GetPostgreSQLRepoImpl() *PostgreSQLRepoImpl {
	if postgreSQLRepoImpl == nil {
		postgreSQLRepoImpl = NewPostgreSQLRepoImpl(config.Database)

	}
	return postgreSQLRepoImpl
}

// NewPostgreSQLRepoImpl is a dependency injection an instance of NewPostgreSQLRepoImpl
func NewPostgreSQLRepoImpl(db *sqlx.DB) *PostgreSQLRepoImpl {
	return &PostgreSQLRepoImpl{
		Db: db,
	}
}
