package main

import (
	"bytes"
	basicIO "io"
	"strings"
	"testing"
)

func TestMain1(t *testing.T) {
	inputString := "2 2\n2 1 2\n1 2\n0 1"
	expect := "1\n"

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
