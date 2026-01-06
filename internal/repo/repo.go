package repo

type Repository interface {
	Load() ([]byte, error)
	Save(data []byte) error
}
