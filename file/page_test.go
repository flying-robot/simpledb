package file

import (
	"reflect"
	"testing"
)

func TestPage_GetSetInt(t *testing.T) {
	page := NewPage(256)
	page.SetInt(0, 1)
	page.SetInt(64, 2)
	page.SetInt(128, 3)

	if got := page.GetInt(0); got != 1 {
		t.Errorf("unexpected int %d", got)
	}
	if got := page.GetInt(64); got != 2 {
		t.Errorf("unexpected int %d", got)
	}
	if got := page.GetInt(128); got != 3 {
		t.Errorf("unexpected int %d", got)
	}
}

func TestPage_GetSetBytes(t *testing.T) {
	page := NewPage(256)
	page.SetBytes(0, []byte{1})
	page.SetBytes(64, []byte{2})
	page.SetBytes(128, []byte{3})

	if got := page.GetBytes(0); !reflect.DeepEqual(got, []byte{1}) {
		t.Errorf("unexpected bytes %x", got)
	}
	if got := page.GetBytes(64); !reflect.DeepEqual(got, []byte{2}) {
		t.Errorf("unexpected bytes %x", got)
	}
	if got := page.GetBytes(128); !reflect.DeepEqual(got, []byte{3}) {
		t.Errorf("unexpected bytes %x", got)
	}
}

func TestPage_GetSetString(t *testing.T) {
	page := NewPage(256)
	page.SetString(0, "hi")
	page.SetString(64, "hello")
	page.SetString(128, "howdy")

	if got := page.GetString(0); got != "hi" {
		t.Errorf("unexpected string %q", got)
	}
	if got := page.GetString(64); got != "hello" {
		t.Errorf("unexpected string %q", got)
	}
	if got := page.GetString(128); got != "howdy" {
		t.Errorf("unexpected string %q", got)
	}
}
