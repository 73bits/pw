package repo

type Repository interface {
	Init(data []byte) error
	Load() ([]byte, error)
	Save(data []byte) error
}
