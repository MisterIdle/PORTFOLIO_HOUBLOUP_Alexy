package logic

import (
	"net/http"
)

type Entry struct {
	Columns []string
	Rows    [][]string
	View    string
}

func DashboardHandler(w http.ResponseWriter, r *http.Request) {
	category := r.URL.Query().Get("view")
	tables := GetAllTablesNames()

	for i, table := range tables {
		if table == "sqlite_sequence" {
			tables = append(tables[:i], tables[i+1:]...)
		}
	}

	var data struct {
		Names   []string
		Entries struct {
			Columns []string
			Rows    [][]string
			View    string
		}
	}

	data.Names = tables
	data.Entries.Columns, _ = GetColumnNames(category)
	data.Entries.Rows, _ = GetValuesFromTable(category)
	data.Entries.View = category

	for i, column := range data.Entries.Columns {
		data.Entries.Columns[i] = Capitalize(column)
	}

	RenderTemplateGlobal(w, r, "templates/dashboard.html", data)
}

func AddContentHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {

		category := r.FormValue("category")

		addContent := make(map[string]string)
		for key, value := range r.Form {
			if key != "category" {
				addContent[key] = value[0]
			}
		}

		addContentConverted := make(map[string][]string)
		for key, value := range addContent {
			addContentConverted[key] = []string{value}
		}

		InsertDataIntoTable(category, addContentConverted)
		http.Redirect(w, r, "/dashboard?view="+category, http.StatusSeeOther)
	}
}

func Capitalize(s string) string {
	if len(s) == 0 {
		return s
	}
	return string(s[0]-32) + s[1:]
}

func DeleteContentHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		category := r.FormValue("category")
		id := r.FormValue("id")

		DeleteRowFromTable(category, id)
		http.Redirect(w, r, "/dashboard?view="+category, http.StatusSeeOther)
	}
}
