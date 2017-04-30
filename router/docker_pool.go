package main

import (
    "strings"
    "os/exec"
    "fmt"
    "crypto/rand"
)

func dockerStuff() (target string) {
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
    target = strings.Replace(sendurl, "\n", "", -1)
    return
}
