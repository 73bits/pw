package cmd

import (
	"encoding/json"
	"log"

	"github.com/73bits/pw/internal/model"
	"github.com/73bits/pw/internal/repo"
	"github.com/73bits/pw/internal/service"
)

func loadVault(r *repo.JSONRepo, password string) (*model.Vault, []byte) {
	data, err := r.Load()
	if err != nil {
		log.Fatal(err)
	}

	if string(data[:4]) != magic {
		log.Fatal("invalid vault")
	}

	salt := data[4 : 4+saltSize]
	encrypted := data[4+saltSize:]

	key, _ := service.DeriveKey(password, salt)
	plain, err := service.Decrypt(key, encrypted)
	if err != nil {
		log.Fatal("wrong password")
	}

	var vault model.Vault
	json.Unmarshal(plain, &vault)

	return &vault, salt
}

func saveVault(r *repo.JSONRepo, password string, salt []byte, vault *model.Vault) {
	plain, _ := json.Marshal(vault)
	key, _ := service.DeriveKey(password, salt)
	encrypted, _ := service.Encrypt(key, plain)

	data := append([]byte(magic), salt...)
	data = append(data, encrypted...)

	if err := r.Save(data); err != nil {
		log.Fatal(err)
	}
}
