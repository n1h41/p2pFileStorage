package main

import (
	"bytes"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPathTransformFunc(t *testing.T) {
	key := "meowmeow"
	pathName := CASPathTransformFunc(key)
	log.Println(pathName)
	expectedPathName := "b6ccb/4ece5/454dc/ae517/78b3e/239eb"
	if pathName != expectedPathName {
		t.Errorf("have %s want %s", pathName, expectedPathName)
	}
}

func TestStore(t *testing.T) {
	storeOpts := StoreOpts{
		PathTransformFunc: CASPathTransformFunc,
	}
	store := NewStore(storeOpts)

	data := bytes.NewReader([]byte("Test data"))
	if err := store.writeStream("meow", data); err != nil {
		assert.Nil(t, err)
	}
}
