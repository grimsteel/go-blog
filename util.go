package main

import (
	"fmt"
	"html/template"
	"net/http"
)

// jetbrains
func check(e error) {
	if e != nil {
		panic(e)
	}
}

// jetbrains
func renderTemplate(data any, templateFile string, w http.ResponseWriter) {
	t, err := template.ParseFiles(
		// base template
		"templates/base.html",
		// page template
		fmt.Sprintf("templates/%s.html", templateFile),
	)

	check(err)
	
	// multiple separate templates 
	check(t.ExecuteTemplate(w, "base", data))
}
