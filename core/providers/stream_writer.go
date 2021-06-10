package providers

import (
	"bytes"
	"io"
)

// StreamWriter writes a stream to a logging function and prefixes each line
// with a prefix string. StreamWriter implements the io.Writer interface.
type StreamWriter struct {
	logFunction func(string, ...interface{})
	prefix      string
	buffer      *bytes.Buffer
}

// NewStreamWriter returns a newly initialised StreamWriter
func NewStreamWriter(logFunction func(string, ...interface{}), prefix string) *StreamWriter {
	return &StreamWriter{
		logFunction: logFunction,
		prefix:      prefix,
		buffer:      bytes.NewBuffer([]byte("")),
	}
}

func (this *StreamWriter) Write(p []byte) (n int, err error) {
	if n, err = this.buffer.Write(p); err != nil {
		return
	}

	err = this.OutputLines()
	return
}

func (this *StreamWriter) Close() error {
	this.Flush()
	this.buffer = bytes.NewBuffer([]byte(""))
	return nil
}

func (this *StreamWriter) Flush() error {
	var p []byte
	if _, err := this.buffer.Read(p); err != nil {
		return err
	}

	this.out(string(p))
	return nil
}

func (this *StreamWriter) OutputLines() (err error) {
	for {
		line, err := this.buffer.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		this.out(line)
	}

	return nil
}

func (this *StreamWriter) out(str string) (err error) {
	this.logFunction("%v: %v", this.prefix, str)

	return nil
}
