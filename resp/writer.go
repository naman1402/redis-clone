package resp

import "io"

type Write struct {
	writer io.Writer
}

func NewWriter(writer io.Writer) *Write {
	return &Write{writer: writer}
}

func (w *Write) Write(v Value) error {
	var bytes = v.Marshal()

	_, err := w.writer.Write(bytes)
	if err != nil {
		return err
	}

	return nil
}
