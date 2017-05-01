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
    _, err := exec.Command("docker", "run", "-Pd", "--name", randomname, "--pids-limit", "20", "soh.re/site").Output()
    check(err)
    // Get port
    port, err := exec.Command("docker", "inspect", "--format='{{(index (index .NetworkSettings.Ports \"8080/tcp\") 0).HostPort}}'", randomname).Output()
    check(err)
    // Send client to port
    sendurl := fmt.Sprintf("localhost:%v", string(port))
    target := strings.Replace(sendurl, "\n", "", -1)
    // Add to pool
    log.Println(target)
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
    for i <= 2 {
        log.Println(i)
        dockerStuff(db)
        i = i + 1
    }
    log.Println("Done with pool, try reaping now")
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
        log.Printf("Going to try and reap %v\n", name)
        DelName(db, name)
        log.Printf("Reaped %v", name)
    }
}

func pool_reaper(db *sql.DB) {
    for {
        // Get all the rows
        var name string
        log.Println("Select name")
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
