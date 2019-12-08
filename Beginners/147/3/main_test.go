package main

import (
	"bytes"
	basicIO "io"
	"strings"
	"testing"
)

func TestMain(t *testing.T) {
	inputString := "3\n1\n2 1\n1\n1 1\n1\n2 0"
	expect := "2\n"

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
