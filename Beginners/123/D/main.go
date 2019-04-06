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

func protagonist(r basicIO.Reader, w basicIO.Writer) {
	io := NewIO(r, w)
	X := io.nextInt()
	Y := io.nextInt()
	Z := io.nextInt()
	K := io.nextInt()

	fmt.Println(X, Y, Z, K)
	var arrx, arry, arrz []int
	for i := 0; i < X; i++ {
		arrx = append(arrx, io.nextInt())
		arry = append(arry, io.nextInt())
		arrz = append(arrz, io.nextInt())
	}

	sort.Sort(sort.Reverse(sort.IntSlice(arrx)))
	sort.Sort(sort.Reverse(sort.IntSlice(arry)))
	sort.Sort(sort.Reverse(sort.IntSlice(arrz)))

	fmt.Println(arrx, arry)

	var x, y, z int
	for i := 0; i < K; i++ {
		x++
		y++
		z++
		io.PutInt(4)
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
