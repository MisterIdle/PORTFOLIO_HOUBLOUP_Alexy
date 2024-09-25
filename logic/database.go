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

		MakeFakeData()

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

func GetDataFromTable(table string) []map[string]interface{} {
	rows, err := db.Query("SELECT * FROM " + table)
	if err != nil {
		fmt.Println("Error querying data:", err)
		return nil
	}
	defer rows.Close()

	// Récupérer les noms des colonnes de la table
	columns, err := rows.Columns()
	if err != nil {
		fmt.Println("Error getting columns:", err)
		return nil
	}

	fmt.Println("Columns:", columns)

	var data []map[string]interface{}

	for rows.Next() {
		values := make([]interface{}, len(columns))
		for i := range values {
			values[i] = new(interface{})
		}

		err = rows.Scan(values...)
		if err != nil {
			fmt.Println("Error scanning rows:", err)
			return nil
		}

		entry := make(map[string]interface{})
		for i, col := range columns {
			val := values[i].(*interface{})
			entry[col] = *val
		}

		data = append(data, entry)
	}

	return data
}

func MakeFakeData() {
	db.Exec("INSERT INTO contact (nom, prenom, email, telephone) VALUES ('Doe', 'John', 'ex', '1234567890');")

	db.Exec("INSERT INTO formations (etablissement, diplome, date_debut, date_fin, description) VALUES ('Université de Paris', 'Licence', '2018-09-01', '2021-06-30', 'Informatique');")
	db.Exec("INSERT INTO formations (etablissement, diplome, date_debut, date_fin, description) VALUES ('Université de Paris', 'Master', '2021-09-01', '2023-06-30', 'Informatique');")
	db.Exec("INSERT INTO formations (etablissement, diplome, date_debut, date_fin, description) VALUES ('Université de Paris', 'Doctorat', '2023-09-01', '2027-06-30', 'Informatique');")

	db.Exec("INSERT INTO experiences (poste, societe, date_debut, date_fin, description) VALUES ('Développeur', 'Google', '2021-07-01', '2021-12-31', 'Développement web');")

	db.Exec("INSERT INTO competences (nom, description) VALUES ('Go', 'Langage de programmation');")
	db.Exec("INSERT INTO competences (nom, description) VALUES ('Python', 'Langage de programmation');")
}
