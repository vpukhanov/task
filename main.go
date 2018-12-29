package main

import (
	"fmt"
	"os"
	"path/filepath"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/vpukhanov/task/cmd"
	"github.com/vpukhanov/task/db"
)

func main() {
	home, err := homedir.Dir()
	if err != nil {
		panic(err)
	}

	dbPath := filepath.Join(home, ".tasks.db")
	must(db.Init(dbPath))
	must(cmd.RootCmd.Execute())
}

func must(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
