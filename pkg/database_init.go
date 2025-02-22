package pkg

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/lib/pq"
)

func InitDatabase(user, pass, dbname, host, port string) (*sql.DB, error) {
	//Connect to PostgreSQL Server
	pgsServerConn := fmt.Sprintf("user=%s password=%s host=%s port=%s sslmode=disable", user, pass, host, port)
	pgsServer, err := sql.Open("postgres", pgsServerConn)
	if err != nil {
		return nil, errors.New("error connecting to PostgreSQL Server")
	}
	defer pgsServer.Close()

	err = pgsServer.Ping()
	if err != nil {
		return nil, errors.New("error pinging PostgreSQL Server")
	}
	//Check if database exists
	var dbExists bool

	err = pgsServer.QueryRow("SELECT EXISTS(SELECT datname FROM pg_catalog.pg_database WHERE datname = $1)", dbname).Scan(&dbExists)
	if err != nil {
		return nil, errors.New("error checking if database exists")
	}
	//Create database if not exists
	if !dbExists {
		_, err = pgsServer.Exec("CREATE DATABASE " + dbname)
		if err != nil {
			return nil, errors.New("error creating database")
		}
	}
	//Connect directly to PostgreSQL Database
	db, err := sql.Open("postgres", "user="+user+" password="+pass+" dbname="+dbname+" host="+host+" port="+port+" sslmode=disable")
	if err != nil {
		return nil, errors.New("error connecting to database")
	}
	//Create table ads if not exists (DSP)
	createAdsTable := `CREATE TABLE IF NOT EXISTS ads (
		id SERIAL PRIMARY KEY,
		dsp_id TEXT,
		ad_name TEXT,
		ad_condition TEXT
		);`
	_, err = db.Exec(createAdsTable)
	if err != nil {
		return nil, errors.New("error creating ads table")
	}
	//Create table deals if not exists (SSP)
	createDealsTable := `CREATE TABLE IF NOT EXISTS tracked (
	id SERIAL PRIMARY KEY,
	bid_Id TEXT,
	price TEXT,
	event_type TEXT
	);`
	_, err = db.Exec(createDealsTable)
	if err != nil {
		return nil, errors.New("error creating tracked table")
	}
	return db, nil
}
