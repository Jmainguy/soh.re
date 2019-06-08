package main

import (
    "strings"
    "os/exec"
    "fmt"
    "crypto/rand"
    "log"
    "time"
    _ "github.com/mattn/go-sqlite3"
    "database/sql"
)

func dockerStuff(db *sql.DB) {
    // Random name for container
    n := 10
    b := make([]byte, n)
    if _, err := rand.Read(b); err != nil {
        check(err)
    }
    randomname := fmt.Sprintf("%X", b)
    // Spin up docker container
    _, err := exec.Command("docker", "run", "-Pd", "--name", randomname, "--pids-limit", "20", "hub.soh.re/soh.re/site").Output()
    check(err)
    // Get port
    portbyte, err := exec.Command("docker", "inspect", "--format='{{(index (index .NetworkSettings.Ports \"8080/tcp\") 0).HostPort}}'", randomname).Output()
    port := string(portbyte)
    port = strings.Replace(port, "'", "", -1)
    check(err)
    // Send client to port
    sendurl := fmt.Sprintf("localhost:%s", port)
    target := strings.Replace(sendurl, "\n", "", -1)
    // Add to pool
    add_docker_to_pool(db, target, randomname)
}

func add_docker_to_pool(db *sql.DB, url, name string) {
    // Store current, and average
    items := []TestItem{
        TestItem{url, name},
    }

    StoreItem(db, items)
}


func pull_docker_from_pool(db *sql.DB) (target string) {
    target = ReadItem(db)
    DelItem(db, target)
    return target
}

func keep_10_in_pool(db *sql.DB) {
    // add 10 to pool initially
    i := 1
    for i <= 10 {
        dockerStuff(db)
        i = i + 1
    }
    go pool_reaper(db)
}

func reap(db *sql.DB, name string) {
    // If container does not exist, remove from pool)
    running, err := exec.Command("docker", "inspect", "--format='{{.State.Running}}'", name).Output()
    check(err)
    is_running := string(running)
    if err != nil {
        DelName(db, name)
        log.Printf("Reaped %v", name)
    }
    if is_running == "false\n" {
        DelName(db, name)
        log.Printf("Reaped %v", name)
    }
}

func pool_reaper(db *sql.DB) {
    for {
        // Get all the rows
        var name string
        rows, err := db.Query("SELECT name FROM docker_pool;")
        var s []string
        for rows.Next() {
           err = rows.Scan(&name)
           check(err)
           s = append(s, name)
        }

        rows.Close()
        for _, v := range s {
           reap(db, v)
        }

        log.Println("Reaper is sleeping")
        time.Sleep(10 * time.Second)
    }
}
