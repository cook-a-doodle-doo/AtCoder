package main

import (
	"bytes"
	basicIO "io"
	"strings"
	"testing"
)

func TestMain1(t *testing.T) {
	inputString := "2 2 2 8\n4 6\n1 5\n3 8"
	expect := "19\n17\n15\n14\n13\n12\n10\n8\n"
	output := checker(inputString, protagonist)
	if output != expect {
		t.Errorf("expcet %v ,but %v", expect, output)
	}
}

func TestMain2(t *testing.T) {
	inputString := "3 3 3 5\n1 10 100\n2 20 200\n1 10 100"
	expect := "400\n310\n310\n301\n301\n"
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
