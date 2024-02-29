package configfile

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

func (c *BasicsStation) Marshal() ([]byte, error) {
	return c.Content, nil
}

func (c *BasicsStation) GetFile() string {
	return c.File.String()
}

func (c *BasicsStation) IsNil() bool {
	return c == nil
}
