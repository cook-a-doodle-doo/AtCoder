package main

import (
	"bytes"
	basicIO "io"
	"strings"
	"testing"
)

func TestMain1(t *testing.T) {
	inputString := "5\n3\n2\n4\n3\n5"
	expect := "7\n"

	output := checker(inputString, protagonist)
	if output != expect {
		t.Errorf("expcet %v ,but %v", expect, output)
	}
}

func TestMain2(t *testing.T) {
	inputString := "10\n123\n123\n123\n123\n123"
	expect := "5\n"

	output := checker(inputString, protagonist)
	if output != expect {
		t.Errorf("expcet %v ,but %v", expect, output)
	}
}

func TestMain3(t *testing.T) {
	inputString := "10000000007\n2\n3\n5\n7\n11"
	expect := "5000000008\n"

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
