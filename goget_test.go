package goget_test

import (
	"goget"
	"testing"
)

type Object interface {}

type Stringer interface {
	String() string
}

func AssertEquals(t *testing.T, expected string, got string) {
	if expected != got {
		t.Error("Expected " + expected + " but got " + got)
	}
}

func AssertNil(t *testing.T, value Object) {
	if value != nil {
		t.Error("Expected nil")
	}
}

func TestGetOutfileSimple(t *testing.T) {
	outfile, err := goget.GetOutfile("http://example.com/file.mp3")
	AssertNil(t, err)
	AssertEquals(t, "file.mp3", outfile)
}

func TestGetOutfileCompound(t *testing.T) {
	outfile, err := goget.GetOutfile("http://example.com/path/to/file.mp3")
	AssertNil(t, err)
	AssertEquals(t, "file.mp3", outfile)
}

func TestGetOutfileDefault(t *testing.T) {
	outfile, err := goget.GetOutfile("http://example.com/")
	AssertNil(t, err)
	AssertEquals(t, "index.html", outfile)
}

func TestGetOutfileDefault2(t *testing.T) {
	outfile, err := goget.GetOutfile("http://example.com")
	AssertNil(t, err)
	AssertEquals(t, "index.html", outfile)
}

