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
	N := io.nextInt()
	A := make([]int, N)

	nega := 0
	absMin := math.MaxInt64
	maxIfAllPosi := 0
	for i := 0; i < N; i++ {
		A[i] = io.nextInt()
		absA := int(math.Abs(float64(A[i])))
		if absA < absMin {
			absMin = absA
		}
		maxIfAllPosi += absA

		if A[i] == 0 {
			continue
		}
		if A[i] < 0 {
			nega++
		}
	}

	var sol int
	if nega%2 == 0 {
		sol = maxIfAllPosi
	} else {
		sol = maxIfAllPosi - 2*absMin
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
