package main

import (
	"bufio"
	"fmt"
	basicIO "io"
	"os"
	"sort"
	"strconv"
)

type Element struct {
	sum int
	x   int
	y   int
	z   int
}

func main() {
	protagonist(os.Stdin, os.Stdout)
}

func protagonist(r basicIO.Reader, w basicIO.Writer) {
	io := NewIO(r, w)
	X := io.nextInt()
	Y := io.nextInt()
	Z := io.nextInt()
	K := io.nextInt()

	var arrx, arry, arrz []int
	for i := 0; i < X; i++ {
		arrx = append(arrx, io.nextInt())
	}
	for i := 0; i < Y; i++ {
		arry = append(arry, io.nextInt())
	}
	for i := 0; i < Z; i++ {
		arrz = append(arrz, io.nextInt())
	}

	sort.Sort(sort.Reverse(sort.IntSlice(arrx)))
	sort.Sort(sort.Reverse(sort.IntSlice(arry)))
	sort.Sort(sort.Reverse(sort.IntSlice(arrz)))

	var pool []Element
	pool = append(pool, Element{arrx[0] + arry[0] + arrz[0], 0, 0, 0})

	for i := 0; i < K; i++ {
		maxNum := 0
		for j, _ := range pool {
			if pool[maxNum].sum < pool[j].sum {
				maxNum = j
			}
		}
		io.PutInt(pool[maxNum].sum)

		x := pool[maxNum].x
		y := pool[maxNum].y
		z := pool[maxNum].z

		hasx := false
		hasy := false
		hasz := false
		for _, e := range pool {
			if e.x == x+1 && e.y == y && e.z == z {
				hasx = true
			} else if e.x == x && e.y == y+1 && e.z == z {
				hasy = true
			} else if e.x == x && e.y == y && e.z == z+1 {
				hasz = true
			}
		}
		if !hasx && x < X-1 {
			pool = append(pool, Element{arrx[x+1] + arry[y] + arrz[z], x + 1, y, z})
		}
		if !hasy && y < Y-1 {
			pool = append(pool, Element{arrx[x] + arry[y+1] + arrz[z], x, y + 1, z})
		}
		if !hasz && z < Z-1 {
			pool = append(pool, Element{arrx[x] + arry[y] + arrz[z+1], x, y, z + 1})
		}

		pool = unset(pool, maxNum)
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

func unset(s []Element, i int) []Element {
	if i >= len(s) {
		return s
	}
	return append(s[:i], s[i+1:]...)
}
