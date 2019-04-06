package main

import (
	"bytes"
	basicIO "io"
	"strings"
	"testing"
)

func TestMain(t *testing.T) {
	inputString := "1\n2\n4\n8\n9\n15"
	expect := "Yay!\n"

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
