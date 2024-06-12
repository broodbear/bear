package main

import (
	_ "embed"
	"log"
	"os"

	"github.com/broodbear/tools/cmd"
)

func main() {
	if err := cmd.App.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
