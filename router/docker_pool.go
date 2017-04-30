package main

import (
    "strings"
    "os/exec"
    "fmt"
    "crypto/rand"
)

func dockerStuff(sqldb string) {
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
    add_docker_to_pool(sqldb, target)
}

func add_docker_to_pool(sqldb, url string) {
    db := InitDB(sqldb)
    CreateTable(db)
    // Store current, and average
    items := []TestItem{
        TestItem{url},
    }

    StoreItem(db, items)
}


func pull_docker_from_pool(sqldb string) (target string) {
    db := InitDB(sqldb)
    CreateTable(db)
    target = ReadItem(db)
    DelItem(db, target)
    return target
}

func keep_10_in_pool(sqldb string) {
    // add 10 to pool initially
    i := 0
    for i <= 10 {
        dockerStuff(sqldb)
        i = i + 1
    }
}
