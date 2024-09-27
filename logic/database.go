package logic

import (
	"database/sql"
	"flag"
	"fmt"
	"os"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

// InitData initializes the database
func InitData() {
	var err error
	db, err = sql.Open("sqlite3", "./database.db")
	if err != nil {
		return
	}

	reset := flag.Bool("reset", false, "Reset the database")
	force := flag.Bool("force", false, "Force the database reset")
	flag.Parse()

	if *force {
		os.Remove("./database.db")
		resetAll()
		createData()

		InsertDataIntoTable("contact", map[string][]string{"nom": {"Doe"}, "prenom": {"John"}, "email": {"email"}, "telephone": {"123456789"}})

		fmt.Println("Database has been deleted and reinitialized 🔄")
	}

	if *reset {
		resetAll()

		fmt.Println("Database has been reset 🔄")
	}

	fmt.Println("Database has been initialized ✔️")
}

// createData creates the tables in the database
func createData() {
	query := `
	CREATE TABLE IF NOT EXISTS contact (
		id INTEGER,
		nom TEXT,
		prenom TEXT,
		email TEXT,
		telephone TEXT
	);
	
	CREATE TABLE IF NOT EXISTS formations (
		id INTEGER,
		etablissement TEXT,
		diplome TEXT,
		date_debut TEXT,
		date_fin TEXT,
		description TEXT
	);

	CREATE TABLE IF NOT EXISTS experiences (
		id INTEGER,
		poste TEXT,
		societe TEXT,
		date_debut TEXT,
		date_fin TEXT,
		description TEXT
	);

	CREATE TABLE IF NOT EXISTS competences (
		id INTEGER,
		nom TEXT,
		description TEXT
	);`

	db.Exec(query)
}

// resetAll resets the database if -force flag is passed
func resetAll() {
	db.Exec("DROP TABLE IF EXISTS contact")
	db.Exec("DROP TABLE IF EXISTS formations")
	db.Exec("DROP TABLE IF EXISTS experiences")
	db.Exec("DROP TABLE IF EXISTS competences")
}

// GetAllTablesNames returns all the tables in the database
func GetAllTablesNames() []string {
	rows, err := db.Query("SELECT name FROM sqlite_master WHERE type='table';")
	if err != nil {
		return nil
	}

	var tables []string
	for rows.Next() {
		var table string
		rows.Scan(&table)
		tables = append(tables, table)
	}

	return tables
}

// GetColumnNames returns the columns of a table
func GetColumnNames(table string) ([]string, error) {
	rows, err := db.Query("SELECT * FROM " + table + " LIMIT 1;")
	if err != nil {
		return nil, fmt.Errorf("error querying table: %s, error: %w", table, err)
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		return nil, fmt.Errorf("error getting columns from table: %s, error: %w", table, err)
	}

	return columns, nil
}

// GetValuesFromTable returns the values from a table
func GetValuesFromTable(table string) ([][]string, error) {
	rows, err := db.Query("SELECT * FROM " + table + ";")
	if err != nil {
		return nil, fmt.Errorf("error querying table: %s, error: %w", table, err)
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		return nil, fmt.Errorf("error getting columns from table: %s, error: %w", table, err)
	}

	var allRows [][]string
	for rows.Next() {
		values := make([]interface{}, len(columns))
		valuePtrs := make([]interface{}, len(columns))
		for i := range columns {
			valuePtrs[i] = &values[i]
		}

		rows.Scan(valuePtrs...)

		rowData := make([]string, len(columns))
		for i, v := range values {
			if v != nil {
				rowData[i] = fmt.Sprintf("%v", v)
			} else {
				rowData[i] = ""
			}
		}

		allRows = append(allRows, rowData)
	}

	return allRows, nil
}

// InsertDataIntoTable inserts data into a table
func InsertDataIntoTable(table string, data map[string][]string) {
	var id int

	// Get the last ID in the table
	err := db.QueryRow("SELECT MAX(id) FROM " + table + ";").Scan(&id)
	if err != nil {
		id = 0
	}

	id++

	columns := getColumns(data)
	values := getValues(data)

	// Insert the data into the table
	query := fmt.Sprintf("INSERT INTO %s (id, %s) VALUES (%d, %s);", table, columns, id, values)
	_, err = db.Exec(query)
	if err != nil {
		fmt.Println("Error inserting data into table:", err)
		return
	}
}

// getColumns returns the columns of a table
func getColumns(data map[string][]string) string {
	var columns string
	for key := range data {
		columns += key + ","
	}

	return columns[:len(columns)-1]
}

// getValues returns the values of a table
func getValues(data map[string][]string) string {
	var values string
	for _, value := range data {
		values += "'" + value[0] + "',"
	}

	return values[:len(values)-1]
}

// DeleteRowFromTable deletes a row from a table
func DeleteRowFromTable(table, id string) {
	idInt, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("Error converting id to integer:", err)
		return
	}

	_, err = db.Exec("DELETE FROM "+table+" WHERE id = ?;", idInt)
	if err != nil {
		fmt.Println("Error deleting row from table:", err)
		return
	}

	// Update the IDs of the rows
	updateQuery := fmt.Sprintf(`
		UPDATE %s
		SET id = id - 1
		WHERE id > ?;`, table)

	_, err = db.Exec(updateQuery, idInt)
	if err != nil {
		fmt.Println("Error updating row IDs in table:", err)
		return
	}

	// Reassign the IDs of the rows
	reassignQuery := fmt.Sprintf(`
		UPDATE %s
		SET id = (
			SELECT COUNT(*) 
			FROM %s AS t 
			WHERE t.id <= %s
		)
		WHERE id = %s;`, table, table, id, id)

	_, err = db.Exec(reassignQuery, id)
	if err != nil {
		fmt.Println("Error reassigning IDs in table:", err)
		return
	}
}
