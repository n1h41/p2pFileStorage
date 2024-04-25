package main

import (
	"io"
	"log"
	"os"
)

type PathTransformFunc func(string) string

type StoreOpts struct {
	PathTransformFunc
}

var DefaultPathTransformFunc = func(key string) string {
	return key
}

type Store struct {
	StoreOpts
}

func NewStore(opts StoreOpts) *Store {
	return &Store{
		StoreOpts: opts,
	}
}

func (s *Store) writeStream(key string, r io.Reader) error {
	pathname := s.PathTransformFunc(key)

	if err := os.MkdirAll(pathname, os.ModePerm); err != nil {
		return err
	}

	filename := "something"
	pathAndFilename := pathname + "/" + filename

	f, err := os.Create(pathAndFilename)
	if err != nil {
		return nil
	}

	n, err := io.Copy(f, r)
	if err != nil {
		return nil
	}

	log.Printf("Written (%d) bytes to disk: %s", n, pathAndFilename)

	return nil
}
