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

type Restaurant struct {
	num   int
	city  string
	point int
}

type RS []Restaurant

func (rs RS) Swap(i, j int) { rs[i], rs[j] = rs[j], rs[i] }
func (rs RS) Len() int      { return len(rs) }
func (rs RS) Less(i, j int) bool {
	if rs[i].city < rs[j].city {
		return true
	}
	if rs[i].city == rs[j].city {
		if rs[i].point < rs[j].point {
			return false
		} else {
			return true
		}
	}
	return false
}

func protagonist(r basicIO.Reader, w basicIO.Writer) {
	io := NewIO(r, w)
	N, _ := io.nextInt()
	rs := make(RS, N)

	for i := 0; i < N; i++ {
		num := i + 1
		city := io.nextString()
		point, _ := io.nextInt()

		rs[i] = Restaurant{num, city, point}
	}
	sort.Sort(rs)
	for _, v := range rs {
		io.PutInt(v.num)
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
