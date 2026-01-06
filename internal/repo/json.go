package repo

import "os"

type JSONRepo struct {
	Path string
}

func (s *JSONRepo) Load() ([]byte, error) {
	return os.ReadFile(s.Path)
}

func (s *JSONRepo) Save(data []byte) error {
	return os.WriteFile(s.Path, data, 0600)
}
