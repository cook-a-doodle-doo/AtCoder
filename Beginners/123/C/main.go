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
	var a [5]int
	var b [5]int
	for i := 0; i < 5; i++ {
		a[i], _ = io.nextInt()
		b[i] = N / a[i]
		if N%a[i] != 0 {
			b[i]++
		}
	}

	max := b[0]
	for i := 1; i < 5; i++ {
		if b[i] > max {
			max = b[i]
		}
	}

	io.PutInt(max + 4)
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
