package repo

import (
	"errors"
	"os"
)

type JSONRepo struct {
	Path string
}

func (s *JSONRepo) Load() ([]byte, error) {
	return os.ReadFile(s.Path)
}

func (s *JSONRepo) Save(data []byte) error {
	tmp := s.Path + ".tmp"
	if err := os.WriteFile(tmp, data, 0600); err != nil {
		return err
	}
	return os.Rename(tmp, s.Path)
}

func NewJSONRepo(path string) *JSONRepo {
	return &JSONRepo{Path: path}
}

func (s *JSONRepo) Init(data []byte) error {
	_, err := os.Stat(s.Path)
	if err == nil {
		return errors.New("valut already exist")
	}

	return os.WriteFile(s.Path, data, 0600)
}
