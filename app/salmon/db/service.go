package db

import (
	"fmt"
	"os"
	"path/filepath"

	"database/sql"

	_ "database/sql"

	_ "modernc.org/sqlite"
)

func CreateLocalDB(profile_name, schema string, args []string) (string, error) {
	db_name := profile_name + "___salmon.sqlite"
	app_path, err := os.Executable()
	if err != nil {
		panic(fmt.Errorf("ERROR FINDING EXECUTABLE PATH: [ ERROR: %s ]", err))
	}
	db_dir := filepath.Dir(app_path)
	db_path := filepath.Join(db_dir, db_name)

	_, err = os.Stat(db_path)
	if err != nil {
		db, err := sql.Open("sqlite", db_path)

		_, err = db.Exec(schema, args)
		if err != nil {
			panic(fmt.Errorf("ERROR IN EXECUTING SCHEMA\n[ ERROR: %s ]", err))
		}
		defer db.Close()
		return "CREATED", nil
	}
	return "EXISTS", nil
}
