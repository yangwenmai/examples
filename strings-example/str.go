package main

import (
	"io"
	"unsafe"
)

const prealloc_size = 20

type String struct {
	data [prealloc_size]byte
	buf  []byte
}

func (s *String) fill(r io.Reader, n uint) error {
	if n > prealloc_size {
		s.buf = make([]byte, n)
	} else {
		s.buf = s.data[:n]
	}
	_, err := io.ReadFull(r, s.buf)
	return err
}

func (s *String) String() string {
	return *(*string)(unsafe.Pointer(&s.buf))
}
