package file

import (
	"encoding/binary"
	"unsafe"
)

var (
	intsize = uint64(unsafe.Sizeof(1))
)

// A Page holds the contents of a block in memory.
type Page struct {
	buf []byte
}

// GetInt returns the value stored at the given offset.
func (p *Page) GetInt(offset uint64) uint64 {
	return binary.BigEndian.Uint64(p.buf[offset : offset+intsize])
}

// SetInt sets the value n at the given offset.
func (p *Page) SetInt(offset, n uint64) {
	tmp := make([]byte, intsize)
	binary.BigEndian.PutUint64(tmp, n)

	for i := uint64(0); i < uint64(len(tmp)); i++ {
		p.buf[i+offset] = tmp[i]
	}
}

// GetBytes returns the bytes at a given offset. The assumption
// is that the length of the following data is provided first,
// which dictates how many bytes to return.
func (p *Page) GetBytes(offset uint64) []byte {
	length := p.GetInt(offset)
	tmp := make([]byte, length)

	for i := uint64(0); i < length; i++ {
		tmp[i] = p.buf[i+offset+intsize]
	}
	return tmp
}

// SetBytes writes bytes at the given offset. The length of
// the byte slice is written first, and then the subsequent
// bytes are offset accordingly (and given consideration of
// the length value written first).
func (p *Page) SetBytes(offset uint64, b []byte) {
	length := uint64(len(b))
	p.SetInt(offset, length)

	for i := uint64(0); i < length; i++ {
		p.buf[i+offset+intsize] = b[i]
	}
}

// GetString returns the string at the given offset. Strings
// are stored as binary blobs for convenience.
func (p *Page) GetString(offset uint64) string {
	tmp := p.GetBytes(offset)
	return string(tmp)
}

// SetString writes the string at the given offset.
func (p *Page) SetString(offset uint64, s string) {
	p.SetBytes(offset, []byte(s))
}

// NewPage returns a pointer to an initialized page.
func NewPage(blocksize uint64) *Page {
	p := &Page{
		buf: make([]byte, blocksize-1),
	}
	return p
}

// NewLogPage returns a pointer to an initialized page.
/*
func NewLogPage(b []byte) *Page {
	return &Page{
		buf: b,
	}
}
*/
