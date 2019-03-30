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
	N, _ := io.nextInt()
	Q, _ := io.nextInt()
	s := io.nextString()

	a := make([]int, N+2)
	for i, _ := range a {
		a[i] = 1
	}
	a[0] = 0
	a[len(a)-1] = 0

	b := make([]int, N+2)
	for i := 0; i < Q; i++ {
		t := io.nextString()
		d := io.nextString()

		for i, v := range a {
			b[i] = v
		}

		fmt.Println(t, d)
		fmt.Println(b)
		for j := 1; j < N+1; j++ {
			if s[j-1] == t[0] {
				if d == "L" {
					a[j-1] += b[j]
					a[j] -= b[j]
				} else {
					a[j+1] += b[j]
					a[j] -= b[j]
				}
			}
		}
		fmt.Println(a)
	}

	io.PutInt(N - (a[0] + a[len(a)-1]))
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
