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

func HandlerPostList(w http.ResponseWriter, req *http.Request) {
	tpl := template.Must(template.ParseFiles("app/views/post-list.html"))

	var err error
	DbConnection, err = sql.Open(config.Config.SQLDriver, config.Config.DbName)
	if err != nil {
		fmt.Println("データベース接続に失敗しました。")
	}
	defer DbConnection.Close()

	postlist, err := models.SelectPostAll(req.URL.Query().Get("user_id"), DbConnection)
	if err != nil {
		fmt.Println(err)
	}

	reqId, _ := strconv.Atoi(req.URL.Query().Get("user_id"))
	trgId := int64(reqId)
	user, err := models.GetUser(trgId, DbConnection)

	values := map[string]interface{}{
		"postlist": postlist,
		"user":     user,
	}

	if err := tpl.ExecuteTemplate(w, "post-list.html", values); err != nil {
		fmt.Println(err)
	}
}
