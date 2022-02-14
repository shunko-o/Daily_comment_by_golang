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

func HandlerUserUpdate(w http.ResponseWriter, req *http.Request) {
	result := updatePostedUser(req)

	tpl := template.Must(template.ParseFiles("app/views/user-registered.html"))

	values := map[string]string{
		"result": result,
	}

	if err := tpl.ExecuteTemplate(w, "user-registered.html", values); err != nil {
		fmt.Println(err)
	}
}

func updatePostedUser(req *http.Request) string {
	result := "ユーザ情報の更新に成功しました。"

	var err error
	DbConnection, err = sql.Open(config.Config.SQLDriver, config.Config.DbName)
	if err != nil {
		fmt.Println("データベース接続に失敗しました。")
	}
	defer DbConnection.Close()

	err = models.UpdateUser(req.FormValue("name"), req.FormValue("id"), DbConnection)
	if err != nil {
		result = "ユーザ情報の更新に失敗しました。"
	}

	return result
}
