package main

import (
    _ "github.com/mattn/go-sqlite3"
    "database/sql"
    "log"
    "fmt"
)

type TestItem struct {
    Url     string
    Name    string
}


func InitDB(filepath string) *sql.DB {
    file := fmt.Sprintf("file:%v?cache=shared&mode=rwc", filepath)
    db, err := sql.Open("sqlite3", file)
    check(err)
    if db == nil { panic("db nil") }
    return db
}

func CreateTable(db *sql.DB) {
    // create table if not exists
    sql_table := `
    CREATE TABLE IF NOT EXISTS docker_pool(
        url string primary key,
        name string
    );
    `
    _, err := db.Exec(sql_table)
    check(err)
}


func ReadItem(db *sql.DB) (url string) {
    sql_readall := `
    SELECT Url FROM docker_pool limit 1;
    `

    stmt, err := db.Prepare(sql_readall)
    check(err)
    err = stmt.QueryRow().Scan(&url)
    check(err)
    stmt.Close()
    return url
}

/*
func StoreItem(db *sql.DB, items []TestItem) {
    sql_additem := `
    INSERT OR REPLACE INTO docker_pool(
        Url,
        Name
    ) values(?, ?)
    `

    stmt, err := db.Prepare(sql_additem)
    check(err)

    for _, item := range items {
        _, err = stmt.Exec(item.Url, item.Name)
        check(err)
    }
    stmt.Close()
}
*/
func StoreItem(db *sql.DB, items []TestItem) {
    sql_additem := `
    INSERT OR REPLACE INTO docker_pool(
        Url,
        Name
    ) values(?, ?)
    `

    routeSQL, err := db.Prepare(sql_additem)
    check(err)

    for _, item := range items {
        tx, err := db.Begin()
        check(err)
        _, err = tx.Stmt(routeSQL).Exec(item.Url, item.Name)
        if err != nil {
            log.Println(err)
            log.Println("doing rollback")
            tx.Rollback()
        } else {
            err = tx.Commit()
            check(err)
        }
    }
}

func DelItem(db *sql.DB, url string) {
    sql_delitem := `
    DELETE FROM docker_pool where Url = ?
    `
    stmt, err := db.Prepare(sql_delitem)
    check(err)
    check(err)
    _, err = stmt.Exec(url)
    check(err)
    stmt.Close()
}

func DelName(db *sql.DB, name string) {
    routeSQL, err := db.Prepare("delete from docker_pool where name = ?;")
    check(err)
    tx, err := db.Begin()
    check(err)
    _, err = tx.Stmt(routeSQL).Exec(name)
    if err != nil {
       log.Println(err)
       log.Println("doing rollback")
       tx.Rollback()
    } else {
       err = tx.Commit()
       check(err)
    }
}
