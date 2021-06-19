package io_test

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/yuyamada/goder/pkg/io"
)

func TestNextLine(t *testing.T) {
	stdin := bytes.NewBufferString("1 2 3\n")
	stdout := new(bytes.Buffer)
	io := io.New(stdin, stdout)
	expected := "1 2 3"
	actual := io.NextLine()
	assert.Equal(t, expected, actual)
}

func TestNext(t *testing.T) {
	stdin := bytes.NewBufferString("1 2 3\n")
	stdout := new(bytes.Buffer)
	io := io.New(stdin, stdout)
	expected := "1"
	actual := io.Next()
	assert.Equal(t, expected, actual)
}

func TestNextInt(t *testing.T) {
	stdin := bytes.NewBufferString("1 2 3\n")
	stdout := new(bytes.Buffer)
	io := io.New(stdin, stdout)
	for _, expected := range []int{1, 2, 3} {
		actual := io.NextInt()
		assert.Equal(t, expected, actual)
	}
}

func TestNextInts(t *testing.T) {
	stdin := bytes.NewBufferString("1 2 3\n")
	stdout := new(bytes.Buffer)
	io := io.New(stdin, stdout)
	expected := []int{1, 2, 3}
	actual := io.NextInts(3)
	assert.Equal(t, expected, actual)
}

func TestNextFloat(t *testing.T) {
	stdin := bytes.NewBufferString("1.0 2.0 3.0\n")
	stdout := new(bytes.Buffer)
	io := io.New(stdin, stdout)
	for _, expected := range []float64{1.0, 2.0, 3.0} {
		actual := io.NextFloat()
		assert.Equal(t, expected, actual)
	}
}

func TestNextFloats(t *testing.T) {
	stdin := bytes.NewBufferString("1.0 2.0 3.0\n")
	stdout := new(bytes.Buffer)
	io := io.New(stdin, stdout)
	expected := []float64{1.0, 2.0, 3.0}
	actual := io.NextFloats(3)
	assert.Equal(t, expected, actual)
}
