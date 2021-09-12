//go:build unit
// +build unit

package main

import (
	"testing"
)

func TestBetterErrorOk(t *testing.T) {
	res, err := BetterError(true)
	if err != nil {
		t.Fatal("should not have failed")
	}

	if res != "ok" {
		t.Fatalf("expected %v got %v", "ok", res)
	}
}

func TestBetterErrorFail(t *testing.T) {
	res, err := BetterError(false)
	if err == nil {
		t.Fatal("should have failed")
	}

	if err == ErrFail {
		t.Logf("failed with known error")
	}

	if res != "fail" {
		t.Fatalf("expected %v got %v", "fail", res)
	}
}
