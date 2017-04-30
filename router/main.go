package main

import (
    "io"
    "log"
    "net"
)

func check(e error) {
    if e != nil {
        log.Println(e)
    }
}

func forward(conn net.Conn) {
    log.Printf("Before Dockerstuff %v\n", conn)
    //target := dockerStuff()
    target := "localhost:8082"
    log.Printf("After Dockerstuff %v\n", conn)
    client, err := net.Dial("tcp", target)
    if err != nil {
	check(err)
    }
    log.Printf("Connected to localhost %v\n", conn)
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
    listener, err := net.Listen("tcp", "0.0.0.0:8085")
    check(err)

    for {
        conn, err := listener.Accept()
        if err != nil {
            log.Fatalf("ERROR: failed to accept listener: %v", err)
        }
        log.Printf("Accepted connection %v\n", conn)
        go forward(conn)
    }
}

