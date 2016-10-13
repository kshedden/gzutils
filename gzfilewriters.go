package gzutils

import (
	"compress/gzip"
	"io"
	"os"
)

type GZFileWriter struct {
	file     *os.File
	gzwriter *gzip.Writer
}

type GZFileWriters struct {
	files     []*os.File
	gzwriters []*gzip.Writer
}

func NewGZFileWriter(name string) *GZFileWriter {

	file, err := os.Create(name)
	if err != nil {
		panic(err)
	}
	gzwriter := gzip.NewWriter(file)

	return &GZFileWriter{file, gzwriter}
}

func (gz *GZFileWriter) Close() {
	gz.gzwriter.Close()
	gz.file.Close()
}

func (gz *GZFileWriter) GetWriter() io.Writer {
	var x io.Writer = gz.gzwriter
	return x
}

func NewGZFileWriters(names []string) *GZFileWriters {

	n := len(names)
	files := make([]*os.File, n)
	gzwriters := make([]*gzip.Writer, n)
	for i := 0; i < n; i++ {
		var err error
		files[i], err = os.Create(names[i])
		if err != nil {
			panic(err)
		}
		gzwriters[i] = gzip.NewWriter(files[i])
	}

	return &GZFileWriters{files, gzwriters}
}

func (gz *GZFileWriters) Close() {

	for i := 0; i < len(gz.files); i++ {
		gz.gzwriters[i].Close()
		gz.files[i].Close()
	}
}

func (gz *GZFileWriters) GetWriters() []io.Writer {

	n := len(gz.gzwriters)
	wtr := make([]io.Writer, n)
	for i := 0; i < n; i++ {
		wtr[i] = gz.gzwriters[i]
	}

	return wtr
}
