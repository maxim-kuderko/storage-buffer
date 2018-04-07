package buffer_drivers

import (
	"bytes"
	"compress/gzip"
	"io"
)

type MemBuffer struct {
	buff *bytes.Buffer
	gz   *gzip.Writer
}

func NewMemBuffer(compression int) io.ReadWriteCloser {
	b := &bytes.Buffer{}
	g, _ := gzip.NewWriterLevel(b, compression)
	return &MemBuffer{
		buff: b,
		gz:  g,
	}
}

func (mb *MemBuffer) Read(p []byte) (n int, err error) {
	return mb.buff.Read(p)
}
func (mb *MemBuffer) Write(p []byte) (n int, err error) {
	i, e := mb.gz.Write(p)
	i2, _ := mb.gz.Write([]byte("\n"))
	return i + i2, e
}

func (mb *MemBuffer) Close() error {
	return mb.gz.Close()
}
