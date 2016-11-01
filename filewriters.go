package gzutils

import (
	"compress/gzip"
	"io"
	"os"
)

type FileWriter struct {
	file *os.File
	*gzip.Writer
}

type FileWriters struct {
	files   []*os.File
	writers []*gzip.Writer
}

func NewFileWriter(name string) *FileWriter {

	file, err := os.Create(name)
	if err != nil {
		panic(err)
	}
	writer := gzip.NewWriter(file)

	return &FileWriter{file, writer}
}

func (gz *FileWriter) Close() {
	gz.Writer.Close()
	gz.file.Close()
}

func NewFileWriters(names []string) *FileWriters {

	n := len(names)
	files := make([]*os.File, n)
	writers := make([]*gzip.Writer, n)
	for i := 0; i < n; i++ {
		var err error
		files[i], err = os.Create(names[i])
		if err != nil {
			panic(err)
		}
		writers[i] = gzip.NewWriter(files[i])
	}

	return &FileWriters{files, writers}
}

func (gz *FileWriters) Close() {

	for i := 0; i < len(gz.files); i++ {
		gz.writers[i].Close()
		gz.files[i].Close()
	}
}

func (gz *FileWriters) GetWriters() []io.Writer {

	n := len(gz.writers)
	wtr := make([]io.Writer, n)
	for i := 0; i < n; i++ {
		wtr[i] = gz.writers[i]
	}

	return wtr
}
