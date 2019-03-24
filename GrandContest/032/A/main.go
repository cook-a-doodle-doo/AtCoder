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

	a := make([]int, 0, 100)
	b := make([]int, 0, 100)
	length, _ := io.NextInt()
	for i := 0; i < length; i++ {
		v, _ := io.NextInt()
		a = append(a, v)
	}

	count := 0
START:
	for i := length - count; i > 0; i-- {
		if a[i-1] == i {
			a = unset(a, i-1)
			b = append(b, i)
			count++
			if len(a) == 0 {
				goto DONE
			}
			goto START
		}
	}
	io.PutInt(-1)
	return

DONE:
	for i, _ := range b {
		io.PutInt(b[len(b)-1-i])
	}
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

func (io *IO) NextString() string {
	io.Scanner.Scan()
	return io.Scanner.Text()
}

func (io *IO) NextInt() (int, error) {
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

func unset(s []int, i int) []int {
	if i >= len(s) {
		return s
	}
	return append(s[:i], s[i+1:]...)
}
