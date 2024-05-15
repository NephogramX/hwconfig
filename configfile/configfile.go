package configfile

import (
	"fmt"
	"os"
)

type File struct {
	Name string
	Path string
}

func (f File) String() string {
	return f.Path + f.Name
}

type ConfigFile interface {
	Write() error
	ReadFrom(p string) error
	GetFile() string
	IsNil() bool
}

func ReadFile(path string) ([]byte, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		return nil, fmt.Errorf("error getting file information: %v", err)
	}
	fileSize := fileInfo.Size()

	data := make([]byte, fileSize)
	_, err = file.Read(data)
	if err != nil {
		return nil, fmt.Errorf("error reading file: %v", err)
	}

	return data, nil
}

func WriteFile(path string, data []byte) error {
	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("error creating file: %v", err)
	}
	defer file.Close()

	_, err = file.Write(data)
	if err != nil {
		return fmt.Errorf("error writing to file: %v", err)
	}

	return nil
}
