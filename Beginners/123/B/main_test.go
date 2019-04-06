package main

import (
	"bytes"
	basicIO "io"
	"strings"
	"testing"
)

func TestMain1(t *testing.T) {
	inputString := "29\n20\n7\n35\n120"
	expect := "215\n"

	output := checker(inputString, protagonist)
	if output != expect {
		t.Errorf("expcet %v ,but %v", expect, output)
	}
}

func TestMain2(t *testing.T) {
	inputString := "101\n86\n119\n108\n57"
	expect := "481\n"
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
