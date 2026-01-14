package cmd

import (
	"fmt"
	"os"

	"golang.org/x/term"
)

const VaultFile = "vault.enc"

func promptHidden(label string) string {
	fmt.Print(label)
	b, _ := term.ReadPassword(int(os.Stdin.Fd()))
	fmt.Println()
	return string(b)
}

func Usage() {
	fmt.Println(`usage:
  	pm init
  	pm add -name github
  	pm get -name github`)
}
