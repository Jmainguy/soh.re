package main

import (
    "fmt"
    "net/http"
    "github.com/gin-gonic/gin"
    "os/exec"
    "crypto/rand"
    "log"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func dockerStuff() (sendurl string) {
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
    log.Println(string(port))
    sendurl = fmt.Sprintf("http://localhost:%v", string(port))
    log.Println(sendurl)
    return sendurl
}

func main() {
    // Read Config, load values
    listenport := 8083

    r := gin.Default()
    r.GET("/", func(c *gin.Context) {
        sendurl := dockerStuff()
        c.Redirect(http.StatusTemporaryRedirect, sendurl)
    })

    listen := fmt.Sprintf(":%v", listenport)
    r.Run(listen) // listen and serve content on 0.0.0.0:$listenport
}
