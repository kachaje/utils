package utils

import (
	_ "embed"

	"bytes"
	"io"
)

//go:embed _utils.zip
var data []byte

type Fauxton struct{}

// FolderName folder used when registering the ui
func (Fauxton) FolderName() string {
	return "_utils"
}

// Reader returns a reader for the zip file
func (Fauxton) Reader() io.Reader {
	return bytes.NewReader(data)
}
