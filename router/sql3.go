package main

import (
    _ "github.com/mattn/go-sqlite3"
    "database/sql"
)

type TestItem struct {
    Url     string
}


func InitDB(filepath string) *sql.DB {
    db, err := sql.Open("sqlite3", filepath)
    check(err)
    if db == nil { panic("db nil") }
    return db
}

func CreateTable(db *sql.DB) {
    // create table if not exists
    sql_table := `
    CREATE TABLE IF NOT EXISTS docker_pool(
        url string primary key
    );
    `
    _, err := db.Exec(sql_table)
    check(err)
}

func StoreItem(db *sql.DB, items []TestItem) {
    sql_additem := `
    INSERT OR REPLACE INTO docker_pool(
        Url
    ) values(?)
    `

    stmt, err := db.Prepare(sql_additem)
    check(err)
    defer stmt.Close()

    for _, item := range items {
        _, err = stmt.Exec(item.Url)
        check(err)
    }
}

func ReadItem(db *sql.DB) (url string) {
    sql_readall := `
    SELECT Url FROM docker_pool limit 1;
    `

    stmt, err := db.Prepare(sql_readall)
    check(err)
    defer stmt.Close()
    err = stmt.QueryRow().Scan(&url)
    check(err)
    return url
}

func DelItem(db *sql.DB, url string) {
    sql_delitem := `
    DELETE FROM docker_pool where Url = ?
    `
    stmt, err := db.Prepare(sql_delitem)
    check(err)
    defer stmt.Close()
    check(err)
    _, err = stmt.Exec(url)
    check(err)
}

