package gzutils

import (
	"compress/gzip"
	"io"
	"os"
)

type GZFileReader struct {
	file *os.File
	*gzip.Reader
}

type GZFileReaders struct {
	files     []*os.File
	gzreaders []*gzip.Reader
}

func NewGZFileReader(name string) *GZFileReader {

	file, err := os.Open(name)
	if err != nil {
		panic(err)
	}
	gzreader, err := gzip.NewReader(file)
	if err != nil {
		panic(err)
	}

	return &GZFileReader{file, gzreader}
}

func (gz *GZFileReader) Close() {
	gz.Reader.Close()
	gz.file.Close()
}

func NewGZFileReaders(names []string) *GZFileReaders {

	n := len(names)
	files := make([]*os.File, n)
	gzreaders := make([]*gzip.Reader, n)
	for i := 0; i < n; i++ {
		var err error
		files[i], err = os.Open(names[i])
		if err != nil {
			panic(err)
		}
		gzreaders[i], err = gzip.NewReader(files[i])
		if err != nil {
			panic(err)
		}
	}

	return &GZFileReaders{files, gzreaders}
}

func (gz *GZFileReaders) Close() {

	for i := 0; i < len(gz.files); i++ {
		gz.gzreaders[i].Close()
		gz.files[i].Close()
	}
}

func (gz *GZFileReaders) GetReaders() []io.Reader {

	n := len(gz.gzreaders)
	rdr := make([]io.Reader, n)
	for i := 0; i < n; i++ {
		rdr[i] = gz.gzreaders[i]
	}

	return rdr
}
