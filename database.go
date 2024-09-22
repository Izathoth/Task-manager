package main

import (
    "database/sql"
    "log"
    _ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func initDatabase() {
    var err error
    db, err = sql.Open("sqlite3", "tasks.db")
    if err != nil {
        log.Fatal(err)
    }
    createTable()
}

func createTable() {
    createUserTable := `
    CREATE TABLE IF NOT EXISTS users (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        username TEXT NOT NULL,
        password TEXT NOT NULL
    );`
    
    createTaskTable := `
    CREATE TABLE IF NOT EXISTS tasks (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        user_id INTEGER,
        title TEXT NOT NULL,
        completed BOOLEAN NOT NULL,
        FOREIGN KEY(user_id) REFERENCES users(id)
    );`

    if _, err := db.Exec(createUserTable); err != nil {
        log.Fatal(err)
    }
    if _, err := db.Exec(createTaskTable); err != nil {
        log.Fatal(err)
    }
}
