package main

import (
	"bytes"
	basicIO "io"
	"strings"
	"testing"
)

func TestMain(t *testing.T) {
	inputString := "5\n7\n9"
	expect := "win\n"

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
