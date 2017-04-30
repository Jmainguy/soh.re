package main

import (
    "io"
    "log"
    "net"
)

func check(e error) {
    if e != nil {
        log.Println(e)
        panic(e)
    }
}

func forward(sqldb string, conn net.Conn) {
    log.Printf("Before Dockerstuff %v\n", conn)
    //target := dockerStuff()
    target := pull_docker_from_pool(sqldb)
    log.Printf("After Dockerstuff %v\n", conn)
    client, err := net.Dial("tcp", target)
    if err != nil {
	check(err)
    }
    log.Printf("Connected to localhost %v\n", conn)
    // Add another host to pool
    go dockerStuff(sqldb)
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
    listener, err := net.Listen("tcp", "0.0.0.0:8085")
    keep_10_in_pool(sqldb)
    check(err)

    for {
        conn, err := listener.Accept()
        if err != nil {
            log.Fatalf("ERROR: failed to accept listener: %v", err)
        }
        log.Printf("Accepted connection %v\n", conn)
        go forward(sqldb, conn)
    }
}

