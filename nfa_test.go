package main

import "testing"

func TestStarMatched(t *testing.T) {
	regex := "a*"
	input := "a"
	stateFrag := compile(regex)
	matched := simulate(input, stateFrag.StartState)
	if !matched {
		t.Errorf("Expect %s match %s", input, regex)
	}
}

func TestStarUnmatched(t *testing.T) {
	regex := "a*"
	input := "b"
	stateFrag := compile(regex)
	matched := simulate(input, stateFrag.StartState)
	if matched {
		t.Errorf("Expect %s not match %s", input, regex)
	}
}

func TestBarMatched(t *testing.T) {
	regex := "a|b"
	input := "b"
	stateFrag := compile(regex)
	matched := simulate(input, stateFrag.StartState)
	if !matched {
		t.Errorf("Expect %s not match %s", input, regex)
	}
}

func TestBarUnmatched(t *testing.T) {
	regex := "a|b"
	input := ""
	stateFrag := compile(regex)
	matched := simulate(input, stateFrag.StartState)
	if matched {
		t.Errorf("Expect %s not match %s", input, regex)
	}
}

func TestDotMatched(t *testing.T) {
	regex := "a.b"
	input := "ab"
	stateFrag := compile(regex)
	matched := simulate(input, stateFrag.StartState)
	if !matched {
		t.Errorf("Expect %s not match %s", input, regex)
	}
}

func TestDotUnmatched(t *testing.T) {
	regex := "a.b"
	input := "a"
	stateFrag := compile(regex)
	matched := simulate(input, stateFrag.StartState)
	if matched {
		t.Errorf("Expect %s not match %s", input, regex)
	}
}

func TestDotStarMatched1(t *testing.T) {
	regex := "(a.b)*"
	input := "abab"
	stateFrag := compile(regex)
	matched := simulate(input, stateFrag.StartState)
	if !matched {
		t.Errorf("Expect %s not match %s", input, regex)
	}
}

func TestDotStarUnmatched1(t *testing.T) {
	regex := "(a.b)*"
	input := "abb"
	stateFrag := compile(regex)
	matched := simulate(input, stateFrag.StartState)
	if matched {
		t.Errorf("Expect %s not match %s", input, regex)
	}
}

func TestDotStarMatched2(t *testing.T) {
	regex := "a.b*"
	input := "abbb"
	stateFrag := compile(regex)
	matched := simulate(input, stateFrag.StartState)
	if !matched {
		t.Errorf("Expect %s not match %s", input, regex)
	}
}

func TestDotStarUnmatched2(t *testing.T) {
	regex := "a.b*"
	input := "aba"
	stateFrag := compile(regex)
	matched := simulate(input, stateFrag.StartState)
	if matched {
		t.Errorf("Expect %s not match %s", input, regex)
	}
}

func TestDotStarMatched3(t *testing.T) {
	regex := "a*.b"
	input := "aaab"
	stateFrag := compile(regex)
	matched := simulate(input, stateFrag.StartState)
	if !matched {
		t.Errorf("Expect %s not match %s", input, regex)
	}
}

func TestDotStarUnmatched3(t *testing.T) {
	regex := "a*.b"
	input := "baab"
	stateFrag := compile(regex)
	matched := simulate(input, stateFrag.StartState)
	if matched {
		t.Errorf("Expect %s not match %s", input, regex)
	}
}

func TestDotStarMatched4(t *testing.T) {
	regex := "a*.b*"
	input := "aaabbbb"
	stateFrag := compile(regex)
	matched := simulate(input, stateFrag.StartState)
	if !matched {
		t.Errorf("Expect %s not match %s", input, regex)
	}
}

func TestDotStarUnmatched(t *testing.T) {
	regex := "a*.b*"
	input := "bbbaaa"
	stateFrag := compile(regex)
	matched := simulate(input, stateFrag.StartState)
	if matched {
		t.Errorf("Expect %s not match %s", input, regex)
	}
}
