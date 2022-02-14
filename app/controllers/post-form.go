package controllers

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
	_ "strconv"

	"daily_comment/app/models"
	"daily_comment/config"

	_ "github.com/mattn/go-sqlite3"
)

func HandlerPostForm(w http.ResponseWriter, req *http.Request) {
	tpl := template.Must(template.ParseFiles("app/views/post-form.html"))

	var err error
	DbConnection, err = sql.Open(config.Config.SQLDriver, config.Config.DbName)
	if err != nil {
		fmt.Println("データベース接続に失敗しました。")
	}
	defer DbConnection.Close()

	models.CreatePostTable(DbConnection)

	values := map[string]string{
		"UserId": req.FormValue("user_id"),
	}

	if err := tpl.ExecuteTemplate(w, "post-form.html", values); err != nil {
		fmt.Println(err)
	}
}
