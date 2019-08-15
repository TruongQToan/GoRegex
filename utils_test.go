package main

import "testing"

func Test1(t *testing.T) {
	infix := "a.b"
	postfix := in2post(infix)
	expect := "ab."
	if postfix != expect {
		t.Errorf("Expect %s match %s", postfix, expect)
	}
}

func Test2(t *testing.T) {
	infix := "a|b"
	postfix := in2post(infix)
	expect := "ab|"
	if postfix != expect {
		t.Errorf("Expect %s match %s", postfix, expect)
	}
}

func Test3(t *testing.T) {
	infix := "a*"
	postfix := in2post(infix)
	expect := "a*"
	if postfix != expect {
		t.Errorf("Expect %s match %s", postfix, expect)
	}
}

func Test4(t *testing.T) {
	infix := "a|b*"
	postfix := in2post(infix)
	expect := "ab*|"
	if postfix != expect {
		t.Errorf("Expect %s match %s", postfix, expect)
	}
}

func Test5(t *testing.T) {
	infix := "a*|b"
	postfix := in2post(infix)
	expect := "a*b|"
	if postfix != expect {
		t.Errorf("Expect %s match %s", postfix, expect)
	}
}

func Test6(t *testing.T) {
	infix := "a*.b"
	postfix := in2post(infix)
	expect := "a*b."
	if postfix != expect {
		t.Errorf("Expect %s match %s", postfix, expect)
	}
}

func Test7(t *testing.T) {
	infix := "a.b*"
	postfix := in2post(infix)
	expect := "ab*."
	if postfix != expect {
		t.Errorf("Expect %s match %s", postfix, expect)
	}
}

func Test8(t *testing.T) {
	infix := "a.b*|c"
	postfix := in2post(infix)
	expect := "ab*.c|"
	if postfix != expect {
		t.Errorf("Expect %s match %s", postfix, expect)
	}
}

func Test9(t *testing.T) {
	infix := "a*.b*|c"
	postfix := in2post(infix)
	expect := "a*b*.c|"
	if postfix != expect {
		t.Errorf("Expect %s match %s", postfix, expect)
	}
}

func Test10(t *testing.T) {
	infix := "a*.b*.c*"
	postfix := in2post(infix)
	expect := "a*b*c*.."
	if postfix != expect {
		t.Errorf("Expect %s match %s", postfix, expect)
	}
}

func Test11(t *testing.T) {
	infix := "a|b.c"
	postfix := in2post(infix)
	expect := "abc.|"
	if postfix != expect {
		t.Errorf("Expect %s match %s", postfix, expect)
	}
}

func Test12(t *testing.T) {
	infix := "a.b|c"
	postfix := in2post(infix)
	expect := "ab.c|"
	if postfix != expect {
		t.Errorf("Expect %s match %s", postfix, expect)
	}
}

func Test13(t *testing.T) {
	infix := "(a|b).c"
	postfix := in2post(infix)
	expect := "ab|c."
	if postfix != expect {
		t.Errorf("Expect %s match %s", postfix, expect)
	}
}
