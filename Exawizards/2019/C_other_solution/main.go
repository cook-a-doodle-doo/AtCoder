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

	l := -1
	r := N
	for i := Q - 1; i >= 0; i-- {
		if d[i] == 'L' {
			if l < N-1 && t[i] == str[l+1] {
				l++
			}
			if r < N && t[i] == str[r] {
				r++
			}
		} else {
			if r > 0 && t[i] == str[r-1] {
				r--
			}
			if l > -1 && t[i] == str[l] {
				l--
			}
		}
		fmt.Println(i, t[i], d[i], l, r)
	}
	if r-l < 0 {
		io.PutInt(0)
		return
	}
	io.PutInt(r - l - 1)
	return
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
