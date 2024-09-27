package logic

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// LaunchApp is a function that launches the application
func LaunchApp() {
	HandleAll()

	r := gin.Default()

	fmt.Println("")
	fmt.Println("######################")
	fmt.Println("Welcome to my website 🚀")
	fmt.Println("######################")
	fmt.Println("")

	// Go func is used to run the server in a goroutine (bug if not used for gin server)
	go func() {
		// Start the server on port 3030
		fmt.Println("Server is running on port 3030 🌐")
		fmt.Println("Visit http://localhost:3030 to see the website")
		if err := http.ListenAndServe(":3030", nil); err != nil {
			fmt.Println("Error starting server: ", err)
		}
	}()

	// API routes
	r.GET("/api/:category", GetCategoryData)

	// Start the server on port 8080
	if err := r.Run(":8080"); err != nil {
		fmt.Println("Error starting server:", err)
	}

	fmt.Println("Server is running on port 8080 🌐")
}

// HandleAll is a function that handles all the routes
func HandleAll() {
	// Handle all the routes
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

// IndexHandler is a function that renders the index page
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	RenderTemplateWithoutData(w, "templates/index.html")
}

// AddContactHandler is a function that adds a contact
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
