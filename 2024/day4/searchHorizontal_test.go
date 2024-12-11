package main

import (
	"strings"
	"testing"
)

func TestSearchHorizontalForward_Fail(t *testing.T) {
	var line = "123XMA"
	var col = strings.Index(line, "X")
	if SearchHorizontal([]string{line}, Position{Row: 0, Col: col}) {
		t.Fail()
	}
}

func TestSearchHorizontalForward_Pass0(t *testing.T) {
	var line = "123XMAS"
	var col = strings.Index(line, "X")
	if !SearchHorizontal([]string{line}, Position{Row: 0, Col: col}) {
		t.Fail()
	}
}

func TestSearchHorizontalForward_Pass1(t *testing.T) {
	var line = "123XMAS23"
	var col = strings.Index(line, "X")
	if !SearchHorizontal([]string{line}, Position{Row: 0, Col: col}) {
		t.Fail()
	}
}

func TestSearchHorizontalForward_Pass2(t *testing.T) {
	var line = "XMAS23"
	var col = strings.Index(line, "X")
	if !SearchHorizontal([]string{line}, Position{Row: 0, Col: col}) {
		t.Fail()
	}
}

func TestSearchHorizontalBackwards_Fail(t *testing.T) {
	var line = "AMX23"
	var col = strings.Index(line, "X")
	if SearchHorizontal([]string{line}, Position{Row: 0, Col: col}) {
		t.Fail()
	}
}

func TestSearchHorizontalBackwards_Pass0(t *testing.T) {
	var line = "SAMX23"
	var col = strings.Index(line, "X")
	if !SearchHorizontal([]string{line}, Position{Row: 0, Col: col}) {
		t.Fail()
	}
}

func TestSearchHorizontalBackwards_Pass1(t *testing.T) {
	var line = "123SAMX"
	var col = strings.Index(line, "X")
	if !SearchHorizontal([]string{line}, Position{Row: 0, Col: col}) {
		t.Fail()
	}
}

func TestSearchHorizontalBackwards_Pass2(t *testing.T) {
	var line = "123SAMX3432"
	var col = strings.Index(line, "X")
	if !SearchHorizontal([]string{line}, Position{Row: 0, Col: col}) {
		t.Fail()
	}
}
