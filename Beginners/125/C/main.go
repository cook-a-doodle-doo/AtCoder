package main

import (
	"bufio"
	"fmt"
	basicIO "io"
	"math"
	"os"
	"strconv"
)

func main() {
	protagonist(os.Stdin, os.Stdout)
}

func protagonist(r basicIO.Reader, w basicIO.Writer) {
	io := NewIO(r, w)
	N, _ := io.nextInt()
	A := make([]int, N)

	min := math.MaxInt64
	sec := math.MaxInt64
	for i := 0; i < N; i++ {
		A[i], _ = io.nextInt()
		if A[i] < sec {
			if A[i] < min {
				sec = min
				min = A[i]
			} else {
				sec = A[i]
			}
		}
	}

	solution := 1

	for num := sec; num > 0; num-- {
		isLimit := false
		for j := 0; j < N; j++ {
			if A[j]%num == 0 {
				continue
			}
			if isLimit {
				goto NotFound
			}
			isLimit = true
		}
		solution = num
		goto Done
	NotFound:
	}

Done:
	io.PutInt(solution)
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
