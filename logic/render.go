package logic

import (
	"fmt"
	"net/http"
	"text/template"
)

// RenderTemplateGlobal is a function that renders a template with data
func RenderTemplateGlobal(w http.ResponseWriter, r *http.Request, tmpl string, data interface{}) {
	tmpt, err := template.ParseFiles(tmpl)
	if err != nil {
		fmt.Print("Error parsing template: ", err)
		return
	}

	err = tmpt.Execute(w, data)
	if err != nil {
		fmt.Print("Error executing template: ", err)
		return
	}
}

// RenderTemplateWithoutData is a function that renders a template without data
func RenderTemplateWithoutData(w http.ResponseWriter, tmpl string) {
	tmpt, err := template.ParseFiles(tmpl)
	if err != nil {
		fmt.Print("Error parsing template: ", err)
		return
	}

	err = tmpt.Execute(w, nil)
	if err != nil {
		fmt.Print("Error executing template: ", err)
		return
	}
}
