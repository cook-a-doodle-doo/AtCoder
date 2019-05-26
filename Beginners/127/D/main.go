package main

import (
	"bufio"
	"fmt"
	basicIO "io"
	"os"
	"sort"
	"strconv"
)

func main() {
	protagonist(os.Stdin, os.Stdout)
}

type BC []struct {
	B int
	C int
}

func (bc BC) Len() int { return len(bc) }
func (bc BC) Swap(i, j int) {
	bc[i], bc[j] = bc[j], bc[i]
}
func (bc BC) Less(i, j int) bool { return bc[i].C > bc[j].C }

func protagonist(r basicIO.Reader, w basicIO.Writer) {
	io := NewIO(r, w)
	N := io.nextInt()
	M := io.nextInt()
	A := make([]int, N)
	bc := make(BC, M)

	for i := 0; i < N; i++ {
		A[i] = io.nextInt()
	}
	sort.Ints(A)

	for i := 0; i < M; i++ {
		bc[i].B = io.nextInt()
		bc[i].C = io.nextInt()
	}
	sort.Sort(bc)

	ap := 0

	for n := 0; n < M; n++ {
		for i := 0; i < bc[n].B; i++ {
			if ap >= N || A[ap] > bc[n].C {
				goto end
			}
			A[ap] = bc[n].C
			ap++
		}
	}
end:

	sol := 0
	for i := 0; i < N; i++ {
		sol += A[i]
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
