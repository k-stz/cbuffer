package circular

import (
	"errors"
	"fmt"
)

// Implement a circular buffer of bytes supporting both overflow-checked writes
// and unconditional, possibly overwriting, writes.
//
// We chose the provided API so that Buffer implements io.ByteReader
// and io.ByteWriter and can be used (size permitting) as a drop in
// replacement for anything using that interface.

// Define the Buffer type here.

type Buffer struct {
	size    int
	buffer  []byte
	reader  int
	writer  int
	content int
}

func NewBuffer(size int) *Buffer {
	// Make creates a slice with length and capacity of given size
	buffer := make([]byte, size)
	return &Buffer{
		size:   size,
		buffer: buffer,
		reader: 0,
		writer: 0,
		// tracks current size of buffer, used to differentiate
		// empty cbuffer from full buffer to track when
		// writer == reader constitutes an overwrite
		content: 0,
	}
}

func (b *Buffer) ReadByte() (byte, error) {
	if b.reader == b.writer && b.content == 0 {
		return 0, errors.New("attempted to read from empty buffer")
	}
	if b.reader == b.writer {
		b.writer = (b.writer + 1) % b.size
	}
	byte := b.buffer[b.reader]
	b.buffer[b.reader] = 0
	b.reader = (b.reader + 1) % b.size
	b.content--
	fmt.Println("ReadByte:(", byte, ")", b)
	return byte, nil
}

// WriteByte appends byte to the circular buffer until its full.
// When the buffer is full an error will be raised, alerting the
// client that further writes are blocked until a slot becomes free.
func (b *Buffer) WriteByte(c byte) error {
	if b.writer == b.reader && b.content != 0 {
		return errors.New("write blocked: circular buffer is full")
	}
	b.buffer[b.writer] = c
	b.writer = (b.writer + 1) % b.size
	b.content++
	fmt.Println("WriteByte:(", c, ")", b)
	return nil
}

func (b *Buffer) Overwrite(c byte) {
	// check whether an overwrite actually takes place
	if (b.writer == b.reader && b.content != 0) {
		b.buffer[b.writer] = c
		b.writer = (b.writer + 1) % b.size
		b.reader = (b.reader + 1) % b.size
	}
	b.WriteByte(c)	
	fmt.Println("Overwrite:(", c, ")", b)
}

func (b *Buffer) Reset() {

	// b = NewBuffer(b.size) // this won't work
	// as NewBuffer returns a *buffer and overwrites the address
	// stored in b. Upon returning from Reset() the callstack gets
	// popped and what b was pointing to stays unchanged.
	*b = *NewBuffer(b.size)
}
