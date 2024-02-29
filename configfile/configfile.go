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
	Marshal() ([]byte, error)
	GetFile() string
	IsNil() bool
}

func CreateConfigFile(c ConfigFile) error {
	if c.IsNil() {
		return nil
	}
	b, err := c.Marshal()
	if err != nil {
		return err
	}
	fmt.Println("Create : ", c.GetFile())
	fmt.Println("Write  : ", string(b))
	file, err := os.Create(c.GetFile())
	if err != nil {
		return err
	}
	_, err = file.Write(b)
	return err
}
