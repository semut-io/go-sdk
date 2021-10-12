package relationaldatabase

import (
	"os"
	"strconv"
)

// GetDSN is used to get a MySQL-compatible DSN URL to the
// application database available for use by the application manager,
// DSN string is of the form `mysql://<DB_USERNAME>:<DB_PASSWORD@<DB_HOST>:<DB_PORT>/<DB_NAME>`,
// NOTE: only application managers will be able to access this database!
func GetDSN() string {
	// MySQL DSN (data source name) for ADB
	return os.Getenv("ADB_DSN")
}

// GetDatabaseDetails is used to get Username, Password, Database name,
// MySQL Host and Port that will be used to connect to relational MySQL ADB
// from the application manager,
// NOTE: only application managers will be able to access this database!
func GetDatabaseDetails() (
	username string, password string, databaseName string,
	host string, port int) {

	username = os.Getenv("ADB_USERNAME")
	password = os.Getenv("ADB_PASSWORD")
	databaseName = os.Getenv("ADB_DB_NAME")
	host = os.Getenv("ADB_HOST")
	port, err := strconv.Atoi(os.Getenv("ADB_PORT"))
	if err != nil {
		port = 3306
	}

	return
}
