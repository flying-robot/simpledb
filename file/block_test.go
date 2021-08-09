package file

import "testing"

func TestBlock_Filename(t *testing.T) {
	b := NewBlock("a.tbl", 1)
	if b.Filename() != "a.tbl" {
		t.Errorf("unexpected block filename %q", b.Filename())
	}
}

func TestBlock_Number(t *testing.T) {
	b := NewBlock("a.tbl", 1)
	if b.Number() != 1 {
		t.Errorf("unexpected block number %d", b.Number())
	}
}

func TestBlock_Equal(t *testing.T) {
	b1 := NewBlock("a.tbl", 1)
	b2 := NewBlock("a.tbl", 1)
	if !b1.Equal(b2) {
		t.Errorf("unexpected block inequality")
	}
}

func TestBlock_NotEqual(t *testing.T) {
	b1 := NewBlock("a.tbl", 1)
	b2 := NewBlock("b.tbl", 2)
	if b1.Equal(b2) {
		t.Errorf("unexpected block equality")
	}
}

func TestBlock_String(t *testing.T) {
	b := NewBlock("a.tbl", 1)
	if b.String() != `[file "a.tbl", block 1]` {
		t.Errorf("unexpected block string %q", b.String())
	}
}
