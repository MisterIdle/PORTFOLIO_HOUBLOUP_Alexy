package logic

import (
	"fmt"
	"net/http"
	"strings"
)

func LaunchApp() {
	HandleAll()
	fmt.Println("")
	fmt.Println("######################")
	fmt.Println("Welcome to the forum!")
	fmt.Println("######################")
	fmt.Println("")

	fmt.Println("Server is running on port 3030 🌐")
	fmt.Println("Visit http://localhost:3030 to access the forum")

	err := http.ListenAndServe(":3030", nil)
	if err != nil {
		fmt.Println("Error starting server: ", err)
	}
}

func HandleAll() {
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))
	http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("img"))))

	http.HandleFunc("/", IndexHandler)
	http.HandleFunc("/about", AboutHandler)
	http.HandleFunc("/dashboard", DashboardHandler)
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	RenderTemplateWithoutData(w, "templates/index.html")
}

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	RenderTemplateWithoutData(w, "templates/about.html")
}

func DashboardHandler(w http.ResponseWriter, r *http.Request) {
	category := r.URL.Query().Get("view")
	tables := GetAllTablesNames()

	// Si le nom est sqlite_sequence, on le retire
	for i, table := range tables {
		if table == "sqlite_sequence" {
			tables = append(tables[:i], tables[i+1:]...)
		}
	}

	for i, table := range tables {
		tables[i] = strings.ToTitle(table)
	}

	var data struct {
		Names   []string
		Entries struct {
			Columns []string
			Values  [][]interface{}
		}
	}

	data.Names = tables

	if category != "" {
		entries := GetDataFromTable(category)
		if len(entries) > 0 {
			data.Entries.Columns = entries[0].Columns

			for _, entry := range entries {
				data.Entries.Values = append(data.Entries.Values, entry.Values)
			}
		}
	}

	RenderTemplateGlobal(w, r, "templates/dashboard.html", data)
}
