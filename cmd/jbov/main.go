package main

import (
	"log"
	"os"
	"github.com/kuking/jbov/cmd"
)

func main() {
	log.SetPrefix("jbov: ")
	cmd.RegisterCommands()
	if err := cmd.RootCmd.Execute(); err != nil {
		log.Fatal(err)
		os.Exit(-1)
	}

}
