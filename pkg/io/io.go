package io

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

var Stdio = New(os.Stdin, os.Stdout, os.Stdout)

type IO struct {
	stdin     *bufio.Reader
	stdout    *bufio.Writer
	stderr    io.Writer
	tokens    []string
	nextToken int
}

func New(stdin io.Reader, stdout io.Writer, stderr io.Writer) *IO {
	return &IO{
		stdin:  bufio.NewReader(stdin),
		stdout: bufio.NewWriter(stdout),
		stderr: stderr,
	}
}

func (io *IO) Flush() {
	err := io.stdout.Flush()
	if err != nil {
		panic(err)
	}
}

func (io *IO) NextLine() string {
	var buffer []byte
	for {
		line, isPrefix, err := io.stdin.ReadLine()
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
	fmt.Fprintln(io.stdout, a...)
}

func (io *IO) Printf(format string, a ...interface{}) {
	fmt.Fprintf(io.stdout, format, a...)
}

func (io *IO) Log(a ...interface{}) {
	fmt.Fprintln(io.stderr, a...)
}

func (io *IO) Logf(format string, a ...interface{}) {
	fmt.Fprintf(io.stderr, format, a...)
}
