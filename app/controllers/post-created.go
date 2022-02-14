package controllers

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"

	"daily_comment/app/models"
	"daily_comment/config"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

func HandlerPostCreated(w http.ResponseWriter, req *http.Request) {
	tpl := template.Must(template.ParseFiles("app/views/post-created.html"))

	result, user_id := insertPost(req)

	values := map[string]interface{}{
		"result":  result,
		"user_id": user_id,
	}

	if err := tpl.ExecuteTemplate(w, "post-created.html", values); err != nil {
		fmt.Println(err)
	}
}

func insertPost(req *http.Request) (r, user_id string) {
	result := "つぶやきが成功しました。"
	var err error

	DbConnection, err = sql.Open(config.Config.SQLDriver, config.Config.DbName)
	if err != nil {
		result = "データベースへの接続に失敗しました。"
	}
	defer DbConnection.Close()

	userId, _ := strconv.Atoi(req.FormValue("user_id"))
	intUserId := int64(userId)

	err = models.CreatePost(req.FormValue("content"), intUserId, DbConnection)
	if err != nil {
		fmt.Println(err)
		result = "つぶやきが失敗しました"
	}

	return result, req.FormValue("user_id")
}
