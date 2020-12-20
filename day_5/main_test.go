package main

import "testing"

func TestDecode(t *testing.T) {
	if id := getRow("FBFBBFF"); id != 44 {
		t.Fatal(id)
	}
	if id := getSeat("RLR"); id != 5 {
		t.Fatal(id)
	}
	if id := getSeatId("FBFBBFFRLR"); id != 357 {
		t.Fatal(id)
	}
}
