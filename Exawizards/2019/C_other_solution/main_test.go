package main

import (
	"bytes"
	basicIO "io"
	"strings"
	"testing"
)

func TestMain1(t *testing.T) {
	inputString := "3 4\nABC\nA L\nB L\nB R\nA R"
	expect := "2\n"

	output := checker(inputString, protagonist)
	if output != expect {
		t.Errorf("expcet %v ,but %v", expect, output)
	}
}

func TestMain2(t *testing.T) {
	inputString := "8 3\nAABCBDBA\nA L\nB R\nA R"
	expect := "5\n"

	output := checker(inputString, protagonist)
	if output != expect {
		t.Errorf("expcet %v ,but %v", expect, output)
	}
}

func TestMain3(t *testing.T) {
	inputString := "10 15\nSNCZWRCEWB\nB R\nR R\nE R\nW R\nZ L\nS R\nQ L\nW L\nB R\nC L\nA L\nN L\nE R\nZ L\nS L"
	expect := "3\n"

	output := checker(inputString, protagonist)
	if output != expect {
		t.Errorf("expcet %v ,but %v", expect, output)
	}
}

func checker(input string, f func(basicIO.Reader, basicIO.Writer)) string {
	reader := strings.NewReader(input)
	writer := bytes.NewBufferString("")
	f(reader, writer)
	return writer.String()
}
