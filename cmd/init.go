package cmd

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"log"

	"github.com/73bits/pw/internal/model"
	"github.com/73bits/pw/internal/repo"
	"github.com/73bits/pw/internal/service"
)

const (
	magic    = "PMV1"
	saltSize = 16
)

func Init() {
	password := promptHidden("master password: ")

	salt := make([]byte, saltSize)
	rand.Read(salt)

	vault := model.Vault{Entries: []model.Entry{}}
	plain, _ := json.Marshal(vault)

	key, _ := service.DeriveKey(password, salt)
	encrypted, _ := service.Encrypt(key, plain)

	data := append([]byte(magic), salt...)
	data = append(data, encrypted...)

	r := repo.NewJSONRepo(VaultFile)
	if err := r.Init(data); err != nil {
		log.Fatal(err)
	}

	fmt.Println("vault initialized")
}
