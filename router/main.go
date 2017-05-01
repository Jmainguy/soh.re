package main

import (
    "io"
    "log"
    "net"
    _ "github.com/mattn/go-sqlite3"
    "database/sql"
)

func check(e error) {
    if e != nil {
        log.Println(e)
        //panic(e)
    }
}

func forward(db *sql.DB, conn net.Conn) {
    target := pull_docker_from_pool(db)
    client, err := net.Dial("tcp", target)
    if err != nil {
	check(err)
    }
    log.Printf("Connected to localhost %v\n", conn)
    // Add another host to pool
    go dockerStuff(db)
    go func() {
        defer client.Close()
        defer conn.Close()
        io.Copy(client, conn)
    }()
    go func() {
        defer client.Close()
        defer conn.Close()
        io.Copy(conn, client)
    }()
}

func main() {
    sqldb := config()
    db := InitDB(sqldb)
    CreateTable(db)
    listener, err := net.Listen("tcp", "0.0.0.0:8085")
    go keep_10_in_pool(db)
    check(err)
    for {
        conn, err := listener.Accept()
        if err != nil {
            log.Fatalf("ERROR: failed to accept listener: %v", err)
        }
        log.Printf("Accepted connection %v\n", conn)
        go forward(db, conn)
    }
}

