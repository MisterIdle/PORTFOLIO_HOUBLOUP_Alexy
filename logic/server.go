package logic

import (
	"fmt"
	"net/http"
)

func LaunchApp() {
	HandleAll()
	HandleAPI()

	fmt.Println("")
	fmt.Println("######################")
	fmt.Println("Welcome to the forum!")
	fmt.Println("######################")
	fmt.Println("")

	fmt.Println("Server is running on port 3030 🌐")
	fmt.Println("Visit http://localhost:3030 to access the folio")

	err := http.ListenAndServe(":3030", nil)
	if err != nil {
		fmt.Println("Error starting server: ", err)
	}
}

func HandleAll() {
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))
	http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("img"))))
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("js"))))

	http.HandleFunc("/", IndexHandler)
	http.HandleFunc("/add", AddContactHandler)

	http.HandleFunc("/experiences", ExperiencesHandler)
	http.HandleFunc("/formations", FormationHandler)
	http.HandleFunc("/skills", SkillsHandler)

	http.HandleFunc("/dashboard", DashboardHandler)

	http.HandleFunc("/dashboard/add", AddContentHandler)
	http.HandleFunc("/dashboard/delete", DeleteContentHandler)
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	RenderTemplateWithoutData(w, "templates/index.html")
}

func AddContactHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		r.ParseForm()

		addContent := make(map[string]string)
		for key, value := range r.Form {
			addContent[key] = value[0]
		}

		addContentConverted := make(map[string][]string)
		for key, value := range addContent {
			addContentConverted[key] = []string{value}
		}

		fmt.Print(addContentConverted)

		InsertDataIntoTable("contact", addContentConverted)
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}
