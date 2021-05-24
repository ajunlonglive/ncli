package main

import (
	"log"
	"netgrif.com/ncli/cmd"
)

func main() {
	root := cmd.NewRootCmd("")
	err := root.Execute()
	if err != nil {
		log.Panic(err)
	}
}
