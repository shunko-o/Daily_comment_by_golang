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

var DbConnection *sql.DB

func HandlerUserRegistered(w http.ResponseWriter, req *http.Request) {
	result := insertPostedUser(req)

	tpl := template.Must(template.ParseFiles("app/views/user-registered.html"))

	values := map[string]string{
		"result": result,
	}

	if err := tpl.ExecuteTemplate(w, "user-registered.html", values); err != nil {
		fmt.Println(err)
	}
}

func insertPostedUser(req *http.Request) string {
	result := "ユーザ情報の登録に成功しました。"
	var err error

	DbConnection, err = sql.Open(config.Config.SQLDriver, config.Config.DbName)
	if err != nil {
		result = "データベースへの接続に失敗しました。"
	}
	defer DbConnection.Close()

	// // 接続確認
	// err = DbConnection.Ping()
	// if err != nil {
	// 	fmt.Println("データベース接続失敗")
	// } else {
	// 	fmt.Println("データベース接続成功!")
	// }

	models.CreateUserTable(DbConnection)

	id, err := models.CreateUser(req.FormValue("name"), DbConnection)
	if err != nil {
		fmt.Println(err)
		result = "ユーザ情報の登録に失敗しました。"
	}
	fmt.Println(id)

	user, err := models.GetUser(id, DbConnection)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("[ID] %s\n", user.Id)
	fmt.Printf("[名前] %s\n", user.Name)
	fmt.Printf("[登録日] %s\n", user.Created)

	return result
}
