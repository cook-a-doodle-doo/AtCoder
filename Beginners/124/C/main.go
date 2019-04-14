package main

import (
	"bufio"
	"fmt"
	basicIO "io"
	"os"
	"strconv"
)

func main() {
	protagonist(os.Stdin, os.Stdout)
}

func protagonist(r basicIO.Reader, w basicIO.Writer) {
	io := NewIO(r, w)
	//	S := io.nextString()
	var S string
	fmt.Scan(&S)

	turnF1 := 0
	//1010101010101
	for i, v := range S {
		if i%2 == 0 {
			if v != '1' {
				turnF1++
			}
		} else {
			if v != '0' {
				turnF1++
			}
		}
	}

	//0101010101010
	turnF0 := 0
	for i, v := range S {
		if i%2 == 0 {
			if v != '0' {
				turnF0++
			}
		} else {
			if v != '1' {
				turnF0++
			}
		}
	}

	min := turnF1
	if turnF1 > turnF0 {
		min = turnF0
	}
	io.PutInt(min)
}

type IO struct {
	Scanner *bufio.Scanner
	Writer  basicIO.Writer
}

func NewIO(r basicIO.Reader, w basicIO.Writer) *IO {
	s := bufio.NewScanner(r)
	s.Split(bufio.ScanWords)
	return &IO{
		Scanner: s,
		Writer:  w,
	}
}

func (io *IO) nextString() string {
	io.Scanner.Scan()
	return io.Scanner.Text()
}

func (io *IO) nextInt() (int, error) {
	io.Scanner.Scan()
	i, err := strconv.Atoi(io.Scanner.Text())
	return i, err
}

func (io *IO) PutInt(v int) {
	fmt.Fprintf(io.Writer, "%d\n", v)
}

func (io *IO) PutString(s string) {
	fmt.Fprintf(io.Writer, "%s\n", s)
}
