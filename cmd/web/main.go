package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"time"

	"git.01.alem.school/bbaktyke/forum.git/pkg/models/mysql"
	_ "github.com/mattn/go-sqlite3"
)

type application struct {
	errorLog         *log.Logger
	infoLog          *log.Logger
	posts            *mysql.PostModel
	users            *mysql.UserModel
	session          *mysql.SessionModel
	comments         *mysql.CommentModel
	reaction         *mysql.ReactionModel
	comment_reaction *mysql.CommentReactionModel
}

func main() {
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	db, err := openDB("forum.db")
	if err != nil {
		errorLog.Fatal(err)
	}
	defer db.Close()
	app := application{
		errorLog:         errorLog,
		infoLog:          infoLog,
		posts:            &mysql.PostModel{DB: db},
		users:            &mysql.UserModel{DB: db},
		session:          &mysql.SessionModel{DB: db},
		comments:         &mysql.CommentModel{DB: db},
		reaction:         &mysql.ReactionModel{DB: db},
		comment_reaction: &mysql.CommentReactionModel{DB: db},
	}
	srv := &http.Server{
		Addr:         ":8080",
		ErrorLog:     errorLog,
		Handler:      app.router(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	infoLog.Println("http://localhost:8080")
	err = srv.ListenAndServe()
	errorLog.Fatal(err)
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
