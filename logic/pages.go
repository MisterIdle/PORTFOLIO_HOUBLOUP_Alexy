package logic

import "net/http"

func ExperiencesHandler(w http.ResponseWriter, r *http.Request) {
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
	data.Entries.Columns, _ = GetColumnNames("experiences")
	data.Entries.Rows, _ = GetValuesFromTable("experiences")
	data.Entries.View = "experiences"

	for i, column := range data.Entries.Columns {
		data.Entries.Columns[i] = Capitalize(column)
	}

	RenderTemplateGlobal(w, r, "templates/experiences.html", data)
}

func FormationHandler(w http.ResponseWriter, r *http.Request) {
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
	data.Entries.Columns, _ = GetColumnNames("formations")
	data.Entries.Rows, _ = GetValuesFromTable("formations")
	data.Entries.View = "formations"

	for i, column := range data.Entries.Columns {
		data.Entries.Columns[i] = Capitalize(column)
	}

	RenderTemplateGlobal(w, r, "templates/formations.html", data)
}

func SkillsHandler(w http.ResponseWriter, r *http.Request) {
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
	data.Entries.Columns, _ = GetColumnNames("competences")
	data.Entries.Rows, _ = GetValuesFromTable("competences")
	data.Entries.View = "competences"

	for i, column := range data.Entries.Columns {
		data.Entries.Columns[i] = Capitalize(column)
	}

	RenderTemplateGlobal(w, r, "templates/skills.html", data)
}
