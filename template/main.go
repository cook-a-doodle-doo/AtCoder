package main

import (
	"bufio"
	"os"
	"strconv"
)

func main() {
	input := NewInPut()
	for i := 0; i < 3; i++ {
		println(input.nextString())
	}

}

type InPut struct {
	Scanner *bufio.Scanner
}

func NewInPut() *InPut {
	s := bufio.NewScanner(os.Stdin)
	s.Split(bufio.ScanWords)
	return &InPut{Scanner: s}
}

func (s *InPut) nextString() string {
	s.Scanner.Scan()
	return s.Scanner.Text()
}

func (s *InPut) nextInt() (int, error) {
	s.Scanner.Scan()
	i, err := strconv.Atoi(s.Scanner.Text())
	return i, err
}
