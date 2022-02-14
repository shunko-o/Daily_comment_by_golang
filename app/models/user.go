package models

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	Id      string `db:"id"`
	Name    string `db:"name"`
	Created string `db:"created"`
}

type UserList []User

func CreateUserTable(db *sql.DB) {
	cmd := fmt.Sprintf(`
					CREATE TABLE IF NOT EXISTS User (
							id INTEGER PRIMARY KEY AUTOINCREMENT,
							name NVARCHAR(20) NOT NULL DEFAULT “”,
							created DATETIME NOT NULL)`)
	db.Exec(cmd)
}

func CreateUser(name string, db *sql.DB) (id int64, err error) {
	stmt, err := db.Prepare("INSERT INTO User(name, created) VALUES(?, datetime('now', 'localtime'))")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(name)
	if err != nil {
		return 0, err
	}

	// オートインクリメントのIDを取得
	createdId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return createdId, nil
}

func GetUser(id int64, db *sql.DB) (userinfo User, err error) {
	var user User

	stmt, err := db.Prepare("SELECT id, name, created FROM User WHERE id = ?")
	if err != nil {
		return user, err
	}

	rows, err := stmt.Query(id)
	if err != nil {
		return user, err
	}
	defer rows.Close()

	rows.Next()
	err = rows.Scan(&user.Id, &user.Name, &user.Created)
	if err != nil {
		return user, err
	}

	return user, nil
}

func SelectUserAll(db *sql.DB) (userlist UserList, err error) {
	var ul UserList

	stmt, err := db.Prepare("SELECT id, name, strftime('%Y/%m/%d/%H:%M:%S', created) FROM User")
	if err != nil {
		return ul, err
	}

	rows, err := stmt.Query()
	if err != nil {
		return ul, err
	}
	defer rows.Close()

	for rows.Next() {
		var user User
		err = rows.Scan(&user.Id, &user.Name, &user.Created)
		ul = append(ul, user)
	}

	return ul, nil
}

func UpdateUser(name, id string, db *sql.DB) (err error) {
	stmt, err := db.Prepare("UPDATE User SET name = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(name, id)
	if err != nil {
		return err
	}

	return nil
}
