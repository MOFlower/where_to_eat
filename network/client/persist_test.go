package client

import (
	"testing"
)

func Test_persister_Append(t *testing.T) {
	iPersister.Append([]byte("hello world\n"))

}

func Test_persister_Tail(t *testing.T) {
	print(iPersister.Tail())
}
