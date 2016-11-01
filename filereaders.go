package gzutils

import (
	"compress/gzip"
	"io"
	"os"
)

type FileReader struct {
	file *os.File
	*gzip.Reader
}

type FileReaders struct {
	files   []*os.File
	readers []*gzip.Reader
}

func NewFileReader(name string) *FileReader {

	file, err := os.Open(name)
	if err != nil {
		panic(err)
	}
	reader, err := gzip.NewReader(file)
	if err != nil {
		panic(err)
	}

	return &FileReader{file, reader}
}

func (gz *FileReader) Close() {
	gz.Reader.Close()
	gz.file.Close()
}

func NewFileReaders(names []string) *FileReaders {

	n := len(names)
	files := make([]*os.File, n)
	readers := make([]*gzip.Reader, n)
	for i := 0; i < n; i++ {
		var err error
		files[i], err = os.Open(names[i])
		if err != nil {
			panic(err)
		}
		readers[i], err = gzip.NewReader(files[i])
		if err != nil {
			panic(err)
		}
	}

	return &FileReaders{files, readers}
}

func (gz *FileReaders) Close() {

	for i := 0; i < len(gz.files); i++ {
		gz.readers[i].Close()
		gz.files[i].Close()
	}
}

func (gz *FileReaders) GetReaders() []io.Reader {

	n := len(gz.readers)
	rdr := make([]io.Reader, n)
	for i := 0; i < n; i++ {
		rdr[i] = gz.readers[i]
	}

	return rdr
}
