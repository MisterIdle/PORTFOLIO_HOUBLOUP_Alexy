package logic

import (
	"database/sql"
	"flag"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

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

		MakeDataFromTable("contact", []interface{}{"1", "Doe", "John", "", ""})

		fmt.Println("Database has been deleted and reinitialized 🔄")
	}

	if *reset {
		resetAll()

		fmt.Println("Database has been reset 🔄")
	}

	fmt.Println("Database has been initialized ✔️")
}

func createData() {
	query := `
	CREATE TABLE IF NOT EXISTS contact (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		nom TEXT,
		prenom TEXT,
		email TEXT,
		telephone TEXT
	);
	
	CREATE TABLE IF NOT EXISTS formations (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		etablissement TEXT,
		diplome TEXT,
		date_debut TEXT,
		date_fin TEXT,
		description TEXT
	);

	CREATE TABLE IF NOT EXISTS experiences (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		poste TEXT,
		societe TEXT,
		date_debut TEXT,
		date_fin TEXT,
		description TEXT
	);

	CREATE TABLE IF NOT EXISTS competences (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		nom TEXT,
		description TEXT
	);`

	db.Exec(query)
}

func resetAll() {
	db.Exec("DROP TABLE IF EXISTS contact")
	db.Exec("DROP TABLE IF EXISTS formations")
	db.Exec("DROP TABLE IF EXISTS experiences")
	db.Exec("DROP TABLE IF EXISTS competences")
}

// Recuperer les noms de toutes les tables dynamiquement
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

func GetDataFromTable(table string) []Entry {
	rows, err := db.Query("SELECT * FROM " + table + ";")
	if err != nil {
		fmt.Println("Error querying table:", table, "Error:", err)
		return nil
	}
	defer rows.Close()

	columns, _ := rows.Columns()

	data := []Entry{}
	for rows.Next() {
		values := make([]interface{}, len(columns))
		pointers := make([]interface{}, len(columns))

		for i := range values {
			pointers[i] = &values[i]
		}

		rows.Scan(pointers...)

		entry := Entry{
			Columns: columns,
			Values:  values,
		}

		data = append(data, entry)
	}

	return data
}

func MakeDataFromTable(table string, data []interface{}) {
	query := "INSERT INTO " + table + " VALUES ("
	for i := 0; i < len(data); i++ {
		query += "?,"
	}
	query = query[:len(query)-1] + ");"

	stmt, err := db.Prepare(query)
	if err != nil {
		fmt.Println("Error preparing query:", query, "Error:", err)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(data...)
	if err != nil {
		fmt.Println("Error executing query:", query, "Error:", err)
		return
	}
}
