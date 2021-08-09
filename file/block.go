package file

import "fmt"

// A Block identifies a logical block within a given file.
type Block struct {
	filename string
	blknum   uint64
}

// Filename returns the name of the file containing the block.
func (b *Block) Filename() string {
	return b.filename
}

// Number returns the block's logical number.
func (b *Block) Number() uint64 {
	return b.blknum
}

// Equal returns true if the supplied block is identical to
// the current one.
func (b *Block) Equal(b2 *Block) bool {
	return b.Filename() == b2.Filename() &&
		b.Number() == b2.Number()
}

// String returns a human-readable representation of the block.
func (b *Block) String() string {
	return fmt.Sprintf(
		"[file %q, block %d]",
		b.Filename(),
		b.Number(),
	)
}

// NewBlock returns a pouint64er to an initialized block.
func NewBlock(filename string, blknum uint64) *Block {
	return &Block{
		filename: filename,
		blknum:   blknum,
	}
}
