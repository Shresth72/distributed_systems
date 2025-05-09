package main

import (
	"bytes"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPathTransformFunc(t *testing.T) {
	key := "thisisalongpath"
	pathname := CASPathTransformFunc(key)
	expectedFilename := "6d104575efa4e4fd3c7307bd99ba93534cfcbeb9"
	expectedPathname := "6d104/575ef/a4e4f/d3c73/07bd9/9ba93/534cf/cbeb9"

	assert.Equal(t, expectedPathname, pathname.Pathname)
	assert.Equal(t, expectedFilename, pathname.Filename)
}

func TestStoreDeleteKey(t *testing.T) {
	opts := StoreOpts{
		PathTransformFunc: CASPathTransformFunc,
	}
	s := NewStore(opts)
	key := "specials"
	data := []byte("some bytes")

	assert.NoError(t, s.writeStream(key, bytes.NewReader(data)))
	assert.NoError(t, s.Delete(key))
}

func TestStore(t *testing.T) {
	opts := StoreOpts{
		PathTransformFunc: CASPathTransformFunc,
	}
	s := NewStore(opts)
	key := "test_folder"

	data := []byte("some data to be written")
	assert.NoError(t, s.writeStream(key, bytes.NewReader(data)))

	r, err := s.Read(key)
	assert.NoError(t, err)

	b, err := io.ReadAll(r)
	assert.Equal(t, data, b)

	assert.True(t, s.Has(key))
	assert.NoError(t, s.Delete(key))
}
