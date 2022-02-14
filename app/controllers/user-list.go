package controllers

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"

	"daily_comment/app/models"
	"daily_comment/config"

	_ "github.com/mattn/go-sqlite3"
)

func HandlerUserList(w http.ResponseWriter, req *http.Request) {
	tpl := template.Must(template.ParseFiles("app/views/user-list.html"))

	var err error
	DbConnection, err = sql.Open(config.Config.SQLDriver, config.Config.DbName)
	if err != nil {
		fmt.Println("データベース接続に失敗しました。")
	}
	defer DbConnection.Close()

	userlist, err := models.SelectUserAll(DbConnection)

	if err := tpl.ExecuteTemplate(w, "user-list.html", userlist); err != nil {
		fmt.Println(err)
	}
}
