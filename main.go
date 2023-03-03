package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/pete-woods/lastpass2keychain/internal/keychain"
)

type opts struct {
	Filename string
	DryRun   bool
}

func main() {
	o := opts{}
	flag.StringVar(&o.Filename, "filename", "", "file to read")
	flag.BoolVar(&o.DryRun, "dry-run", true, "dry run mode")
	flag.Parse()
	err := keychain.Load(o.Filename, o.DryRun)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
