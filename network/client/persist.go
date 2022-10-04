package client

import (
	"io"
	"log"
	"os"
)

const (
	FILE_PATH = "./w2e.store.data"
)

var (
	iPersister = persister{}
)

type persister struct {
	f *os.File
}

func init() {
	f, err := os.OpenFile(FILE_PATH, os.O_CREATE|os.O_RDWR, 0777)
	if err != nil {
		log.Fatalln(err.Error())
	}
	iPersister.f = f

}

func (p *persister) Tail() string {
	_, err := p.f.Seek(0, io.SeekEnd)
	if err != nil {
		log.Fatalln(err.Error())
	}
	fs, err := p.f.Stat()
	if err != nil {
		log.Fatalln(err.Error())
	}
	fileSize := fs.Size()
	back := int64(0)
	b := make([]byte, 1)
	for cnt := 0; cnt < 4; {
		back--
		if back*-1 >= fileSize {
			break
		}
		_, err = p.f.Seek(back, io.SeekEnd)
		if err != nil {
			log.Fatalln(err.Error())
		}

		_, err = p.f.Read(b)
		if err != nil {
			log.Fatalln(err.Error())
		}

		if b[0] == '\n' {
			cnt++
		}

	}
	if back*-1 < fileSize {
		back += 1
	}
	_, err = p.f.Seek(back, io.SeekEnd)
	if err != nil {
		log.Fatalln(err.Error())
	}
	b = make([]byte, back*-1)
	nn, err := p.f.Read(b)
	if err != nil {
		log.Fatalln(err.Error())
	}
	if int64(nn) != back*-1 {
		log.Fatalln("tail failed, read", nn)
	}
	ret := string(b)
	return ret
}

func (p *persister) Append(b []byte) {
	_, err := p.f.Seek(0, io.SeekEnd)
	if err != nil {
		return
	}
	nn, err := p.f.Write(b)
	if err != nil || nn != len(b) {
		log.Fatalln(err.Error())
	}
}

func (p *persister) Close() {
	err := p.f.Close()
	if err != nil {
		log.Fatalln(err.Error())
	}
}
