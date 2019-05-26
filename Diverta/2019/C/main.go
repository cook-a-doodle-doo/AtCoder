package main

import (
	"bufio"
	"fmt"
	basicIO "io"
	"os"
	"strconv"
	"strings"
)

func main() {
	protagonist(os.Stdin, os.Stdout)
}

func protagonist(r basicIO.Reader, w basicIO.Writer) {
	io := NewIO(r, w)
	N, _ := io.nextInt()

	rA := 0
	lB := 0
	rAlB := 0
	sol := 0
	for i := 0; i < N; i++ {
		str := io.nextString()
		sol += strings.Count(str, "AB")
		if 'B' == str[0] {
			lB++
		}
		if 'A' == str[len(str)-1] {
			rA++
			if 'B' == str[0] {
				rAlB++
			}
		}
	}

	if rA < lB {
		sol += rA
	} else {
		sol += lB
	}
	if rAlB == rA || rA == lB {
		sol--
	}

	io.PutInt(sol)
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
