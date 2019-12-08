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
	N := io.nextInt()
	arr := make([][][]int, N)
	for i := range arr {
		arr[i] = make([][]int, io.nextInt())
		for j := range arr[i] {
			arr[i][j] = make([]int, 2)
			arr[i][j][0] = io.nextInt()
			arr[i][j][1] = io.nextInt()
		}
	}
	for i := 2 ^ N; i >= 0; i-- {
	}
}

func isCont(arr [][][]int, s int, now int) {
	for i := range arr {
		for j := range arr[i] {
			p := arr[i][j][0]
			if 2^p
		}
	}
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

func (io *IO) nextInt() int {
	io.Scanner.Scan()
	i, _ := strconv.Atoi(io.Scanner.Text())
	return i
}

func (io *IO) PutInt(v int) {
	fmt.Fprintf(io.Writer, "%d\n", v)
}

func (io *IO) PutString(s string) {
	fmt.Fprintf(io.Writer, "%s\n", s)
}
