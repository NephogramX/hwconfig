package configfile

import (
	"errors"
	"path/filepath"
)

type BasicsStation struct {
	File
	Content []byte
}

func NewAuthenticationFile(file File, content *[]byte) *BasicsStation {
	return &BasicsStation{
		File:    file,
		Content: *content,
	}
}

func (c *BasicsStation) Write() error {
	if c == nil {
		return nil
	}
	return writeFile(c.File.String(), c.Content)
}

func (c *BasicsStation) ReadFrom(p string) error {
	if c == nil {
		return errors.New("nil interface")
	}
	c.File = File{
		Name: filepath.Base(p),
		Path: filepath.Dir(p),
	}
	var err error
	c.Content, err = readFile(p)
	if err != nil {
		return err
	}

	return nil
}
