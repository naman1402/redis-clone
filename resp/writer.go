package resp

import "io"

type Write struct {
	writer io.Writer
}

func NewWriter(writer io.Writer) *Write {
	return &Write{writer: writer}
}

func (w *Write) Write(v Value) error {
	return nil
}
