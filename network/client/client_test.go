package client

import (
	"context"
	"log"
	"testing"
)

func TestClient(t *testing.T) {
	err := ClientEnd.Commit(context.TODO(), "hello")
	if err != nil {
		log.Fatalln(err.Error())
	}
}
