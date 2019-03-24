package main

import (
	"bytes"
	basicIO "io"
	"strings"
	"testing"
)

func TestMain(t *testing.T) {
	inputString := "9\n1 1 1 2 2 1 2 3 2"
	expect := "1\n2\n2\n3\n1\n2\n2\n1\n1\n"

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
