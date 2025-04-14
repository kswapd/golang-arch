package sql

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

const (
	username = "kxw"
	password = "000000"
	hostname = "47.100.97.226:33010"
	dbName   = "auth-test"
)

func dsn() string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, hostname, dbName)
}
func SqlTest() {

	s := dsn()
	//sql.Register("mysql", &mysql.MySQLDriver{})
	db, err := sql.Open("mysql", s)
	if err != nil {
		log.Printf("Error %s when opening DB\n", err)
		return
	}

	//ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	//defer cancelfunc()
	//res, err := db.ExecContext(ctx, "select * from sys_menu.;")
	//db.Exec("use auth-test;")
	rows, err := db.Query("SELECT * FROM sys_menu")
	if err != nil {
		log.Printf("Error: %s\n", err)
		return
	}

	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	// Get column names
	columns, err := rows.Columns()
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	// Make a slice for the values
	values := make([]sql.RawBytes, len(columns))

	// rows.Scan wants '[]interface{}' as an argument, so we must copy the
	// references into such a slice
	// See http://code.google.com/p/go-wiki/wiki/InterfaceSlice for details
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	// Fetch rows
	for rows.Next() {
		// get RawBytes from data
		err = rows.Scan(scanArgs...)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}

		// Now do something with the data.
		// Here we just print each column as a string.
		var value string
		for i, col := range values {
			// Here we can check if the value is nil (NULL value)
			if col == nil {
				value = "NULL"
			} else {
				value = string(col)
			}
			fmt.Println(columns[i], ": ", value)
		}
		fmt.Println("-----------------------------------")
	}
	if err = rows.Err(); err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	log.Printf("res: %+v, %s\n", rows)
}
