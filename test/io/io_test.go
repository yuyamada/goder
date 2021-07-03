package io_test

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/yuyamada/goder/pkg/io"
)

func TestNextLine(t *testing.T) {
	stdin := bytes.NewBufferString("1 2 3\n")
	io := io.New(stdin, nil, nil)
	expected := "1 2 3"
	actual := io.NextLine()
	assert.Equal(t, expected, actual)
}

func TestNext(t *testing.T) {
	stdin := bytes.NewBufferString("1 2 3\n")
	io := io.New(stdin, nil, nil)
	expected := "1"
	actual := io.Next()
	assert.Equal(t, expected, actual)
}

func TestNextInt(t *testing.T) {
	stdin := bytes.NewBufferString("1 2 3\n")
	io := io.New(stdin, nil, nil)
	for _, expected := range []int{1, 2, 3} {
		actual := io.NextInt()
		assert.Equal(t, expected, actual)
	}
}

func TestNextInts(t *testing.T) {
	stdin := bytes.NewBufferString("1 2 3\n")
	io := io.New(stdin, nil, nil)
	expected := []int{1, 2, 3}
	actual := io.NextInts(3)
	assert.Equal(t, expected, actual)
}

func TestNextFloat(t *testing.T) {
	stdin := bytes.NewBufferString("1.0 2.0 3.0\n")
	io := io.New(stdin, nil, nil)
	for _, expected := range []float64{1.0, 2.0, 3.0} {
		actual := io.NextFloat()
		assert.Equal(t, expected, actual)
	}
}

func TestNextFloats(t *testing.T) {
	stdin := bytes.NewBufferString("1.0 2.0 3.0\n")
	io := io.New(stdin, nil, nil)
	expected := []float64{1.0, 2.0, 3.0}
	actual := io.NextFloats(3)
	assert.Equal(t, expected, actual)
}

func TestPrintln(t *testing.T) {
	stdout := new(bytes.Buffer)
	io := io.New(nil, stdout, nil)
	expected := "hello\n"
	io.Println("hello")
	io.Flush()
	actual := stdout.String()
	assert.Equal(t, expected, actual)
}

func TestPrintf(t *testing.T) {
	stdout := new(bytes.Buffer)
	io := io.New(nil, stdout, nil)
	expected := "hello, world\n"
	io.Printf("hello, %s\n", "world")
	io.Flush()
	actual := stdout.String()
	assert.Equal(t, expected, actual)
}

func TestLog(t *testing.T) {
	stderr := new(bytes.Buffer)
	io := io.New(nil, nil, stderr)
	expected := "hello\n"
	io.Log("hello")
	actual := stderr.String()
	assert.Equal(t, expected, actual)
}

func TestLogf(t *testing.T) {
	stderr := new(bytes.Buffer)
	io := io.New(nil, nil, stderr)
	expected := "hello, world\n"
	io.Logf("hello, %s\n", "world")
	actual := stderr.String()
	assert.Equal(t, expected, actual)
}
