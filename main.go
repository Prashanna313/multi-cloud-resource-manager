package main

import (
    "multi-cloud-resource-manager/cmd"
)

func main() {
    if err := cmd.Execute(); err != nil {
        println(err)
    }
}
