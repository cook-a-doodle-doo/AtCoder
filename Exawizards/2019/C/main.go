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

func protagonist(reader basicIO.Reader, writer basicIO.Writer) {
	io := NewIO(reader, writer)
	var N int
	var Q int
	var str string
	N, _ = io.nextInt()
	Q, _ = io.nextInt()
	str = io.nextString()
	//	fmt.Scan(&N, &Q, &str)
	t := make([]byte, Q)
	d := make([]byte, Q)

	for i := 0; i < Q; i++ {
		t[i] = io.nextString()[0]
		d[i] = io.nextString()[0]
	}

	fl := FlowLeft(N, Q, str, t, d)
	fr := FlowRight(N, Q, str, t, d)
	io.PutInt(N - fr - fl)
}

func DriftToSection(n int, N int, Q int, str string, t []byte, d []byte) byte {
	for i := 0; i < Q; i++ {
		if str[n] == t[i] {
			if d[i] == 'L' {
				n--
			} else {
				n++
			}
			if n < 0 {
				return 'L'
			}
			if n >= N {
				return 'R'
			}
		}
	}
	return 'M'
}

func FlowLeft(N int, Q int, str string, t []byte, d []byte) int {
	l, r := 0, N-1
	tmp := DriftToSection(l, N, Q, str, t, d)
	if tmp == 'M' || tmp == 'R' {
		return 0
	} else if DriftToSection(r, N, Q, str, t, d) == 'L' {
		return N
	}

	for (r - l) > 1 {
		m := int((l + r) / 2)
		switch DriftToSection(m, N, Q, str, t, d) {
		case 'L':
			l = m
		case 'R':
			r = m
		case 'M':
			r = m
		}
	}
	return r
}
func FlowRight(N int, Q int, str string, t []byte, d []byte) int {
	l, r := 0, N-1
	tmp := DriftToSection(r, N, Q, str, t, d)
	if tmp == 'M' || tmp == 'L' {
		return 0
	} else if DriftToSection(l, N, Q, str, t, d) == 'R' {
		return N
	}

	for (r - l) > 1 {
		m := int((l + r) / 2)
		switch DriftToSection(m, N, Q, str, t, d) {
		case 'L':
			l = m
		case 'R':
			r = m
		case 'M':
			l = m
		}
	}
	return N - r
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
