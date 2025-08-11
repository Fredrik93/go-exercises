package main

import (
	"testing"
)

func TestReverseWordsPreserveStructure_He_Wo(t *testing.T) {
	input := "He wo"
	expected := "eH ow"
	actual := ReverseWordsPreserveStructure(input)
	if actual != expected {
		t.Errorf("input:\n%q\nexpected:\n%q\ngot:\n%q", input, expected, actual)
	}
}

func TestReverseWordsPreserveStructure(t *testing.T) {
	input := "He, who must."
	expected := "eH, ohw tsum."
	actual := ReverseWordsPreserveStructure(input)
	if actual != expected {
		t.Errorf("\ninput:%q,\nexpected: %q,\ngot: %q", input, expected, actual)
	}
}

func TestReverseWordsPreserveStructure_HelloWorld(t *testing.T) {
	input := "Hello world!"
	expected := "olleH dlrow!"
	actual := ReverseWordsPreserveStructure(input)
	if actual != expected {
		t.Errorf("input:\n%q\nexpected:\n%q\ngot:\n%q", input, expected, actual)
	}
}

func TestReverseWordsPreserveStructure_GoIsFun(t *testing.T) {
	input := "Go is fun."
	expected := "oG si nuf."
	actual := ReverseWordsPreserveStructure(input)
	if actual != expected {
		t.Errorf("input:\n%q\nexpected:\n%q\ngot:\n%q", input, expected, actual)
	}
}

func TestReverseWordsPreserveStructure_TestCase(t *testing.T) {
	input := "A test-case."
	expected := "A esac-tset."
	actual := ReverseWordsPreserveStructure(input)
	if actual != expected {
		t.Errorf("input:\n%q\nexpected:\n%q\ngot:\n%q", input, expected, actual)
	}
}

func TestReverseWordsPreserveStructure_Empty(t *testing.T) {
	input := ""
	expected := ""
	actual := ReverseWordsPreserveStructure(input)
	if actual != expected {
		t.Errorf("input:\n%q\nexpected:\n%q\ngot:\n%q", input, expected, actual)
	}
}
