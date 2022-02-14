package controllers

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"daily_comment/app/models"
	"daily_comment/config"

	_ "github.com/mattn/go-sqlite3"
)

func HandlerUserEdit(w http.ResponseWriter, req *http.Request) {
	tpl := template.Must(template.ParseFiles("app/views/user-form.html"))

	var err error
	DbConnection, err = sql.Open(config.Config.SQLDriver, config.Config.DbName)
	if err != nil {
		fmt.Println("データベース接続に失敗しました。")
	}
	defer DbConnection.Close()

	reqId, _ := strconv.Atoi(req.FormValue("id"))
	trgId := int64(reqId)

	user, err := models.GetUser(trgId, DbConnection)

	if err := tpl.ExecuteTemplate(w, "user-form.html", user); err != nil {
		fmt.Println(err)
	}
}
