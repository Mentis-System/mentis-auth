package main

import "testing"

func TestGetMessage(t *testing.T) {
	if GetMessage() != "Hello" {
		t.Errorf("Error Test")
	}
}
