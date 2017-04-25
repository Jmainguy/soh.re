package main

import (
    "io"
    "log"
    "net"
    "strings"
    "os/exec"
    "fmt"
    "crypto/rand"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func forward(conn net.Conn) {
    target := dockerStuff()
    client, err := net.Dial("tcp", target)
    if err != nil {
        log.Fatalf("Dial failed: %v", err)
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

func dockerStuff() (target string) {
    // Random name for container
    n := 10
    b := make([]byte, n)
    if _, err := rand.Read(b); err != nil {
        panic(err)
    }
    randomname := fmt.Sprintf("%X", b)
    // Spin up docker container
    _, err := exec.Command("docker", "run", "-Pd", "--name", randomname, "soh.re/site").Output()
    check(err)
    // Get port
    port, err := exec.Command("docker", "inspect", "--format='{{(index (index .NetworkSettings.Ports \"8080/tcp\") 0).HostPort}}'", randomname).Output()
    check(err)
    // Send client to port
    sendurl := fmt.Sprintf("localhost:%v", string(port))
    target = strings.Replace(sendurl, "\n", "", -1)
    return
}


func main() {
    listener, err := net.Listen("tcp", "localhost:8085")
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

