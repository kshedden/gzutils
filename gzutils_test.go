package gzutils

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestWrite(t *testing.T) {
	g := NewFileWriter("0.gz")
	msg := "This is a single file\n"
	g.Write([]byte(msg))
	g.Close()
}

func TestRead(t *testing.T) {
	g := NewFileReader("0.gz")
	b, err := ioutil.ReadAll(g)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v\n", string(b))
	g.Close()
}

func TestWriteMulti(t *testing.T) {

	names := []string{"0.gz", "1.gz", "2.gz", "3.gz"}
	g := NewFileWriters(names)
	w := g.GetWriters()

	for i, _ := range names {
		msg := fmt.Sprintf("This is file %d\n", i)
		w[i].Write([]byte(msg))
	}

	g.Close()
}

func TestReadMulti(t *testing.T) {

	names := []string{"0.gz", "1.gz", "2.gz", "3.gz"}
	g := NewFileReaders(names)
	r := g.GetReaders()

	for i, _ := range names {
		b, err := ioutil.ReadAll(r[i])
		if err != nil {
			t.Fail()
		}
		fmt.Printf("%v\n", string(b))
	}

	g.Close()
}
