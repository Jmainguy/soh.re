package main

import (
    "io/ioutil"
    "github.com/ghodss/yaml"
)

type Config struct {
    Sqldb string `json:"sqldb"`
}

func config() (sqldb string){
    var v Config
    config_file, err := ioutil.ReadFile("/etc/soh_router/config.yaml")
    check(err)
    yaml.Unmarshal(config_file, &v)
    sqldb = v.Sqldb
    return
}
