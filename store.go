package main

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"log"
	"os"
	"strings"
)

func CASPathTransformFunc(key string) string {
	hash := md5.Sum([]byte(key))
	hashStr := hex.EncodeToString(hash[:])

	blockSize := 5
	sliceLen := len(hashStr) / blockSize

	paths := make([]string, sliceLen)
	for i := 0; i < sliceLen; i++ {
		from, to := i*blockSize, (i*blockSize)+blockSize
		paths[i] = hashStr[from:to]
	}

	return strings.Join(paths, "/")
}

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
