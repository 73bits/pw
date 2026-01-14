package cmd

import (
	"flag"
	"log"
	"os"

	"github.com/73bits/pw/internal/repo"
)

func Get() {
	fs := flag.NewFlagSet("get", flag.ExitOnError)
	name := fs.String("name", "", "entry name")
	fs.Parse(os.Args[2:])

	if *name == "" {
		log.Fatal("missing -name")
	}

	password := promptHidden("master password: ")
	r := repo.NewJSONRepo(VaultFile)

	vault, _ := loadVault(r, password)

	for _, e := range vault.Entries {
		if e.Name == *name {
			println(e.Password)
			return
		}
	}

	log.Fatal("entry not found")
}
