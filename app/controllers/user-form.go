package controllers

import (
	"fmt"
	"html/template"
	"net/http"
)

func HandlerUserForm(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.ParseFiles("app/views/user-form.html"))

	values := map[string]string{}

	if err := tpl.ExecuteTemplate(w, "user-form.html", values); err != nil {
		fmt.Println(err)
	}
}
