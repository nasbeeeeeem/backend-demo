package cloudsql

import (
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx/v4/stdlib"
)

func ConnectWithConnector() (*sql.DB, error) {
	var (
		dbUser         = "user1"                                                   // e.g. 'my-db-user'
		dbPwd          = "Ke~`73)[Bp+/1I,}"                                        // e.g. 'my-db-password'
		unixSocketPath = "/cloudsql/api-gateway-demo-404010:us-central1:instance1" // e.g. 'project:region:instance'
		dbName         = "demo"                                                    // e.g. 'my-database'
	)
	dbURI := fmt.Sprintf("user=%s password=%s database=%s host=%s", dbUser, dbPwd, dbName, unixSocketPath)

	// dbPool is the pool of database connections.
	dbPool, err := sql.Open("pgx", dbURI)
	if err != nil {
		return nil, fmt.Errorf("sql.Open: %w", err)
	}

	return dbPool, nil

}
