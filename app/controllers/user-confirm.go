package controllers

import (
	"fmt"
	"html/template"
	"net/http"
)

func HandlerUserConfirm(w http.ResponseWriter, req *http.Request) {
	tpl := template.Must(template.ParseFiles("app/views/user-confirm.html"))

	values := map[string]string{
		"name":     req.FormValue("name"),
		"hid_name": req.FormValue("name"),
	}

	if err := tpl.ExecuteTemplate(w, "user-confirm.html", values); err != nil {
		fmt.Println(err)
	}
}
