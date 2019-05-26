package main

import (
	"bytes"
	basicIO "io"
	"strings"
	"testing"
)

func TestMain1(t *testing.T) {
	inputString := "6\nkhabarovsk 20\nmoscow 10\nkazan 50\nkazan 35\nmoscow 60\nkhabarovsk 40"
	expect := "3\n4\n6\n1\n5\n2\n"
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
