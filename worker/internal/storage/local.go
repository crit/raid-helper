package storage

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
)

type LocalStorage struct {
	dir string
	mux sync.RWMutex
}

func (l *LocalStorage) Get(fileName string) ([]byte, error) {
	l.mux.RLock()
	defer l.mux.RUnlock()

	data, err := os.ReadFile(filepath.Join(l.dir, fileName))

	if os.IsNotExist(err) {
		return nil, nil
	}

	return data, err
}

func (l *LocalStorage) Set(fileName string, data []byte) error {
	l.mux.Lock()
	defer l.mux.Unlock()

	return os.WriteFile(filepath.Join(l.dir, fileName), data, 0664)
}

func (l *LocalStorage) SetModel(fileName string, version int, data []byte) error {
	return l.Set(fullName(fileName, version, "json"), data)
}

func (l *LocalStorage) GetModel(fileName string, version int) ([]byte, error) {
	return l.Get(fullName(fileName, version, "json"))
}

func NewLocalStorage(dir string) (FileStorage, error) {
	_, err := os.Stat(dir)
	if os.IsNotExist(err) {
		return nil, fmt.Errorf("directory does not exist: '%s'", dir)
	}

	return &LocalStorage{dir: dir}, nil
}
