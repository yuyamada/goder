package io

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type IO struct {
	reader    *bufio.Reader
	writer    *bufio.Writer
	tokens    []string
	nextToken int
}

func New(r io.Reader, w io.Writer) *IO {
	return &IO{
		reader: bufio.NewReader(r),
		writer: bufio.NewWriter(w),
	}
}

func (io *IO) Flush() {
	err := io.writer.Flush()
	if err != nil {
		panic(err)
	}
}

func (io *IO) NextLine() string {
	var buffer []byte
	for {
		line, isPrefix, err := io.reader.ReadLine()
		if err != nil {
			panic(err)
		}
		buffer = append(buffer, line...)
		if !isPrefix {
			break
		}
	}
	return string(buffer)
}

func (io *IO) Next() string {
	for io.nextToken >= len(io.tokens) {
		line := io.NextLine()
		io.tokens = strings.Fields(line)
		io.nextToken = 0
	}
	r := io.tokens[io.nextToken]
	io.nextToken++
	return r
}

func (io *IO) NextInt() int {
	i, err := strconv.Atoi(io.Next())
	if err != nil {
		panic(err)
	}
	return i
}

func (io *IO) NextInts(n int) []int {
	ret := make([]int, n)
	for i := 0; i < n; i++ {
		ret[i] = io.NextInt()
	}
	return ret
}

func (io *IO) NextFloat() float64 {
	f, err := strconv.ParseFloat(io.Next(), 64)
	if err != nil {
		panic(err)
	}
	return f
}

func (io *IO) NextFloats(n int) []float64 {
	ret := make([]float64, n)
	for i := 0; i < n; i++ {
		ret[i] = io.NextFloat()
	}
	return ret
}

func (io *IO) Println(a ...interface{}) {
	fmt.Fprintln(io.writer, a...)
}

func (io *IO) Printf(format string, a ...interface{}) {
	fmt.Fprintf(io.writer, format, a...)
}

func (io *IO) Log(a ...interface{}) {
	fmt.Fprintln(os.Stderr, a...)
}

func (io *IO) Logf(format string, a ...interface{}) {
	fmt.Fprintf(os.Stderr, format, a...)
}
