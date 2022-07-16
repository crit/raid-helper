package storage

import "fmt"

type FileStorage interface {
	Set(fileName string, data []byte) error
	Get(fileName string) ([]byte, error)
	SetModel(fileName string, version int, data []byte) error
	GetModel(fileName string, version int) ([]byte, error)
}

func fullName(name string, version int, ext string) string {
	return fmt.Sprintf("%s-%d.%s", name, version, ext)
}
