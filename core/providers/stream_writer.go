package providers

import (
	"bytes"
	"io"
)

type StreamWriter struct {
	logFunction func(string, ...interface{})
	prefix      string
	buffer      *bytes.Buffer
}

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
