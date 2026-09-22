package sharedTypes

import "testing"

func TestST_OnOffStringUsesWordOneZero(t *testing.T) {
	on, err := ParseUnionST_OnOff("1")
	if err != nil {
		t.Fatal(err)
	}
	if got := on.String(); got != "1" {
		t.Fatalf("on from 1: got %q, want 1", got)
	}

	fromTrue, err := ParseUnionST_OnOff("true")
	if err != nil {
		t.Fatal(err)
	}
	if got := fromTrue.String(); got != "1" {
		t.Fatalf("on from true: got %q, want 1", got)
	}

	off, err := ParseUnionST_OnOff("0")
	if err != nil {
		t.Fatal(err)
	}
	if got := off.String(); got != "0" {
		t.Fatalf("off from 0: got %q, want 0", got)
	}
}
