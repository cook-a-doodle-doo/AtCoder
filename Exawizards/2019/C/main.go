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
	N, _ := io.nextInt()
	Q, _ := io.nextInt()
	s := io.nextString()

	t := make([]byte, Q)
	d := make([]byte, Q)

	for i := 0; i < Q; i++ {
		t[i] = io.nextString()[0]
		d[i] = io.nextString()[0]
	}

	l := FlowLeft(N, Q, s, t, d)
	if l >= N {
		io.PutInt(0)
		return
	}
	r := FlowRight(N, Q, s, t, d)
	if r >= N {
		io.PutInt(0)
		return
	}
	io.PutInt(N - (l + r))
}

func FlowLeft(N int, Q int, str string, t []byte, d []byte) int {
	l := 0
	r := N - 1
	if !isReachTheLeft(l, N, Q, str, t, d) {
		return 0
	}
	for {
		m := (l + r) / 2
		if isReachTheLeft(m, N, Q, str, t, d) {
			l = m
		} else {
			r = m
		}
		if (r - l) <= 1 {
			return r
		}
	}
	return 1
}
func FlowRight(N int, Q int, str string, t []byte, d []byte) int {
	l := 0
	r := N - 1
	if !isReachTheRight(r, N, Q, str, t, d) {
		return 0
	}
	for {
		m := int(float64(l+r)/2 + 0.5)
		if isReachTheRight(m, N, Q, str, t, d) {
			r = m
		} else {
			l = m
		}
		if (r - l) <= 1 {
			return N - l - 1
		}
	}
}

func isReachTheLeft(n int, N int, Q int, str string, t []byte, d []byte) bool {
	for i := 0; i < Q; i++ {
		if str[n] == t[i] {
			if d[i] == 'L' {
				n--
			} else {
				n++
			}
			if n < 0 {
				return true
			}
			if N <= n {
				return false
			}
		}
	}
	return false
}

func isReachTheRight(n int, N int, Q int, str string, t []byte, d []byte) bool {
	for i := 0; i < Q; i++ {
		if str[n] == t[i] {
			if d[i] == 'L' {
				n--
			} else {
				n++
			}
			if N <= n {
				return true
			}
			if n < 0 {
				return false
			}
		}
	}
	return false
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
