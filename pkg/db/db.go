package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

var Error = log.New(os.Stdout, "\n\u001b[31mERROR: \u001b[0m", log.LstdFlags|log.Lshortfile)

type Database struct {
	DB *sql.DB
}

func (database *Database) Open() (err error) {
	file := os.Getenv("DB_PATH_FILE")
	fmt.Printf("Using DB_PATH_FILE: %v\n", file)
	database.DB, err = sql.Open("sqlite3", file)
	return err
}

func (database *Database) Close() {
	database.DB.Close()
}

func (db *Database) GetComments(gh_owner string, gh_repo string, gh_issuenr string) []Comment {
	comments := []Comment{}

	query := `
	SELECT COMMENT.gh_owner, COMMENT.gh_repo, COMMENT.gh_issuenr, COMMENT.comment_text 
	FROM COMMENT
	WHERE COMMENT.gh_owner = "` + gh_owner + `"
	AND COMMENT.gh_repo = "` + gh_repo + `"
	AND COMMENT.gh_issuenr = "` + gh_issuenr + `"
	`
	fmt.Printf("Query: %v\n", query)

	rows, err := db.DB.Query(query)
	if err != nil {
		Error.Fatal(err)
	}
	for rows.Next() {
		comment := Comment{}
		err = rows.Scan(&comment.Owner, &comment.Repo, &comment.IssueNr, &comment.Text)
		if err != nil {
			Error.Fatal(err)
		}
		comments = append(comments, comment)
	}

	fmt.Printf("Comments: %v\n", comments)

	return comments
}

func (db *Database) PutComment(gh_owner string, gh_repo string, gh_issuenr string, text string) {

}
