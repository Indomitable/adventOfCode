package main

import "testing"

func TestSearchVertical_Fail(t *testing.T) {
	lines := []string{
		"TEST",
		"XARS",
		"XMAS",
		"1XAS",
	}
	if SearchVertical(lines, Position{Row: 3, Col: 1}) {
		t.Fail()
	}
}

func TestSearchVertical_Pass0(t *testing.T) {
	lines := []string{
		"TSST",
		"XARS",
		"XMAS",
		"1XAS",
	}
	if !SearchVertical(lines, Position{Row: 3, Col: 1}) {
		t.Fail()
	}
}

func TestSearchVertical_Pass1(t *testing.T) {
	lines := []string{
		"TSXT",
		"XAMS",
		"XMAS",
		"1XSS",
	}
	if !SearchVertical(lines, Position{Row: 0, Col: 2}) {
		t.Fail()
	}
}

func TestSearchVertical_Fail1(t *testing.T) {
	lines := []string{
		"TSXT",
		"XAMS",
		"XMAS",
		"1XSS",
	}
	if SearchVertical(lines, Position{Row: 1, Col: 0}) {
		t.Fail()
	}
}
