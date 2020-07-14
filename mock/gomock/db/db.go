package db
//go:generate mockgen -destination mock_genenrate_test.go -package db database/sql/driver Conn,Driver

type Repository interface {
	Create(key string, value []byte) error
	Retrieve(key string) ([]byte, error)
	Update(key string, value []byte) error
	Delete(key string) error
}