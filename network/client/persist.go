package client

import (
	"bufio"
	"log"
	"os"
)

const FILE_PATH = "./w2e.store.data"

var (
	iPersister = persister{}
)

type persister struct {
	f  *os.File
	rw *bufio.ReadWriter
}

func init() {
	f, err := os.OpenFile(FILE_PATH, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0777)
	if err != nil {
		log.Fatalln(err.Error())
	}
	iPersister.f = f
	iPersister.rw = bufio.NewReadWriter(bufio.NewReader(f), bufio.NewWriter(f))

}

func (p *persister) Append(b []byte) {
	nn, err := p.rw.Write(b)
	if err != nil || nn != len(b) {
		log.Fatalln(err.Error())
	}

	err = p.rw.Flush()
	if err != nil {
		log.Fatalln(err.Error())
	}
}
