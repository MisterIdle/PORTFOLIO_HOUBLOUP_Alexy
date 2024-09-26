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

	http.HandleFunc("/dashboard/add", AddContentHandler)
	http.HandleFunc("/dashboard/delete", DeleteContentHandler)
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	RenderTemplateWithoutData(w, "templates/index.html")
}

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	RenderTemplateWithoutData(w, "templates/about.html")
}
