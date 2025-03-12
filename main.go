package main

import (
	"log"
	"github.com/wilmoore/pro/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
