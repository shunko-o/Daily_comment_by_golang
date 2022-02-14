package models

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"strconv"
)

type Post struct {
	Id      string `db:"id"`
	Content string `db:"content"`
	Created string `db:"created"`
	UserId  string `db:"user_id"`
}

type PostList []Post

func CreatePostTable(db *sql.DB) {
	cmd := fmt.Sprintf(`
					CREATE TABLE IF NOT EXISTS Post (
							id INTEGER PRIMARY KEY AUTOINCREMENT,
							content NVARCHAR(100) DEFAULT “”,
							user_id INTEGER NOT NULL,
							created DATETIME NOT NULL)`)
	db.Exec(cmd)
}

func CreatePost(content string, user_id int64, db *sql.DB) (err error) {
	stmt, err := db.Prepare("INSERT INTO Post(content, user_id, created) VALUES(?, ?, datetime('now', 'localtime'))")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(content, user_id)
	if err != nil {
		return err
	}

	return err
}

func SelectPostAll(userid string, db *sql.DB) (postlist PostList, err error) {
	var pl PostList

	userId, _ := strconv.Atoi(userid)
	intUserId := int64(userId)

	stmt, err := db.Prepare("SELECT id, content, strftime('%Y/%m/%d/%H:%M:%S', created), user_id FROM Post WHERE user_id = ?")
	if err != nil {
		return pl, err
	}

	rows, err := stmt.Query(intUserId)
	if err != nil {
		return pl, err
	}
	defer rows.Close()

	for rows.Next() {
		var post Post
		err = rows.Scan(&post.Id, &post.Content, &post.Created, &post.UserId)
		pl = append(pl, post)
	}

	return pl, nil
}
