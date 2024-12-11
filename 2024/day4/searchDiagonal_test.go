package main

import "testing"

func TestSearchDiagonal_Pass0(t *testing.T) {
	lines := []string{
		"XEST",
		"XMRS",
		"XMAS",
		"1XAS",
	}
	if !SearchDiagonal(lines, Position{Row: 0, Col: 0}) {
		t.Fail()
	}
}

func TestSearchDiagonal_Pass1(t *testing.T) {
	lines := []string{
		"SEST",
		"XARS",
		"XMMS",
		"1XAX",
	}
	if !SearchDiagonal(lines, Position{Row: 3, Col: 3}) {
		t.Fail()
	}
}

func TestSearchDiagonal_Pass2(t *testing.T) {
	lines := []string{
		"SESX",
		"XAMS",
		"XAMS",
		"SXAX",
	}
	if !SearchDiagonal(lines, Position{Row: 0, Col: 3}) {
		t.Fail()
	}
}

func TestSearchDiagonal_Pass4(t *testing.T) {
	lines := []string{
		"SESS",
		"XAAS",
		"XMMS",
		"XXAX",
	}
	if !SearchDiagonal(lines, Position{Row: 3, Col: 0}) {
		t.Fail()
	}
}

func TestSearchDiagonal_Fail0(t *testing.T) {
	lines := []string{
		"SESS",
		"XXAS",
		"XMMS",
		"XXAS",
	}
	if SearchDiagonal(lines, Position{Row: 1, Col: 1}) {
		t.Fail()
	}
}
