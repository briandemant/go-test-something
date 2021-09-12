//go:build unit
// +build unit

package main

import (
	"io/ioutil"
	"testing"
)

func TestReadFile(t *testing.T) {
	data, err := ioutil.ReadFile("testdata/data.txt")
	if err != nil {
		t.Fatal("can't read test file")
	}

	if string(data) != "Hello World" {
		t.Fatalf("content not as expected (expected %v got %v)", "Hello World", string(data))
	}
}
