package main

import (
	"net/http"
	"daily_comment/app/controllers"
)

func main() {
	http.HandleFunc("/user-list", controllers.HandlerUserList)
	http.HandleFunc("/user-form", controllers.HandlerUserForm)
	http.HandleFunc("/user-confirm", controllers.HandlerUserConfirm)
	http.HandleFunc("/user-registered", controllers.HandlerUserRegistered)
	http.HandleFunc("/user-edit", controllers.HandlerUserEdit)
	http.HandleFunc("/user-update", controllers.HandlerUserUpdate)

	http.HandleFunc("/post-list", controllers.HandlerPostList)
	http.HandleFunc("/post-form", controllers.HandlerPostForm)
	http.HandleFunc("/post-created", controllers.HandlerPostCreated)

	// css・js・イメージファイル等の静的ファイル格納パス
	http.Handle("/asset/", http.StripPrefix("/asset/", http.FileServer(http.Dir("app/asset/"))))

	http.ListenAndServe(":8080", nil)
}
