package logic

import (
	"fmt"
	"net/http"
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
	data := make(map[string]interface{})

	filteredTables := []string{}
	for _, table := range tables {
		if table != "sqlite_sequence" {
			filteredTables = append(filteredTables, table)
		}
	}

	data["names"] = filteredTables

	if category != "" {
		content := GetDataFromTable(category)
		data[category] = content
		data["selectedCategory"] = content
		data["categoryName"] = category
	} else {
		for _, table := range filteredTables {
			content := GetDataFromTable(table)
			data[table] = content
		}
	}

	RenderTemplateGlobal(w, r, "templates/dashboard.html", data)
}
