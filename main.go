package main

import (
	"fmt"
	"github.com/zivmitrani/github-status/github"

)

func main() {
    fmt.Println("Hello, World!")

	status, _ := github.GetStatus("https://www.githubstatus.com/api/v2/status.json")
	fmt.Println(status)
}