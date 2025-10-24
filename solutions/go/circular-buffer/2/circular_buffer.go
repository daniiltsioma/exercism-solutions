package circular

import (
	"errors"
)

// Implement a circular buffer of bytes supporting both overflow-checked writes
// and unconditional, possibly overwriting, writes.
//
// We chose the provided API so that Buffer implements io.ByteReader
// and io.ByteWriter and can be used (size permitting) as a drop in
// replacement for anything using that interface.

// Define the Buffer type here.
type Buffer struct {
	size int
	head int
	tail int
	count int
	elements []byte
}

var ErrEmpty = errors.New("empty buffer")
var ErrFull = errors.New("buffer full")

func NewBuffer(size int) *Buffer {
	buf := Buffer{
		size: size,
		head: 0,
		tail: 0,
		count: 0,
	}
	buf.elements = make([]byte, size)
	return &buf
}

func (b *Buffer) ReadByte() (byte, error) {
	if b.count == 0 {
		return 0, ErrEmpty
	}	
	val := b.elements[b.head]
	b.head = (b.head + 1) % b.size
	b.count--
	return val, nil
}

func (b *Buffer) WriteByte(c byte) error {
	if b.count == b.size {
		return ErrFull
	}

	b.elements[b.tail] = c
	b.tail = (b.tail + 1) % b.size
	b.count++
	return nil
}

func (b *Buffer) Overwrite(c byte) {
	if b.count == b.size {
		b.elements[b.head] = c
		b.head = (b.head + 1) % b.size
	} else {
		b.elements[b.tail] = c
		b.tail = (b.tail + 1) % b.size
		b.count++
	}
}

func (b *Buffer) Reset() {
	b.head = 0
	b.tail = 0
	b.count = 0
	for i := range b.elements {
		b.elements[i] = 0
	}
}
