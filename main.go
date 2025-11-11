package main

import (
	"fmt"

	"github.com/milocs/server-cd/internal/cd"
)

func main() {
	cfg := cd.Config{}
	fmt.Println("Config:", cfg)
	changed, err := cd.repoChanged("github.com/go-git/go-git", "main", "96c9ec5f250777017fe34a61b342af0a613225b7")
	if err != nil {
		panic(err)
	}
	if changed {
		fmt.Println("changed")
	} else {
		fmt.Println("same")
	}
}
