package cmd

import (
	"flag"
	"log"
	"os"

	"github.com/73bits/pw/internal/model"
	"github.com/73bits/pw/internal/repo"
)

func Add() {
	fs := flag.NewFlagSet("add", flag.ExitOnError)
	name := fs.String("name", "", "entry name")
	fs.Parse(os.Args[2:])

	if *name == "" {
		log.Fatal("missing -name")
	}

	password := promptHidden("master password: ")
	r := repo.NewJSONRepo(VaultFile)

	vault, salt := loadVault(r, password)

	vault.Entries = append(vault.Entries, model.Entry{
		Name:     *name,
		Password: promptHidden("entry password: "),
	})

	saveVault(r, password, salt, vault)
}
